"""Remote TensorFlow (RTF) gRPC service provider."""
import io
import queue
from collections import deque
import threading
import sys
from typing import Iterator
import re
import contextlib
import tensorflow as tf
import rtf_pb2_grpc
import rtf_pb2


def rreplace(s, old, new, occurrence):
    li = s.rsplit(old, occurrence)
    return new.join(li)


class Builder:

    HEADER = (
        "import tensorflow as tf\n"
        "import sys\n\n"
        # "@tf.function\n"
        "def f():\n"
    )
    FOOTER = "\nf()\n"

    def __init__(self):
        self._statements = []
        self._indent_re = re.compile(r"^(\s*)")

    def _flush_stdout(self, stmt):
        # TODO: convert stmt to AST
        # Understand if print is a function call or not
        # and if it is a call, then add the flush()
        # The current solution is unsafe and terrible.a
        if "tf.print(" in stmt and not "output_stream" in stmt:
            stmt = rreplace(stmt, ")", ", output_stream=sys.stdout)", -1)
        if "print(" in stmt:
            indent = self._indent_re.search(stmt)[0]
            stmt += f"\n{indent}sys.stdout.flush()"

        return stmt

    def build(self, stmt):
        # TODO: check stmt correctness using its AST
        stmt = self._flush_stdout(stmt)
        self._statements.append(stmt)

    def __call__(self):
        source = Builder.HEADER
        for stmt in self._statements:
            source += f"    {stmt}\n"
        source += Builder.FOOTER
        f = compile(source, "useless", "exec")
        # print(source)
        return eval(f)


class DoubleIO(io.StringIO):
    def __init__(self, initial_value="", newline="\n"):
        super().__init__(initial_value, newline)
        self._lines = deque()
        self._buffer = []

    def flush(self):
        self._lines.append("".join(self._buffer))
        self._buffer.clear()

    def write(self, s):
        # self._lines.append(s)
        self._buffer.append(s)

    def close(self):
        super().close()

    def readline(self):
        return self._lines.popleft()


class RTFServicer(rtf_pb2_grpc.RTFServicer):
    """Remote TensorFlow (RTF) gRPC service provider."""

    def DefineAndCall(self, request_iterator, context) -> Iterator[rtf_pb2.RTFResponse]:
        builder = Builder()
        for statement in request_iterator:
            builder.build(statement.stmt)

        fp = DoubleIO()
        stop = False
        response_q = queue.Queue()

        def executor():
            with contextlib.redirect_stdout(fp):
                output_value = builder()
            fp.close()

            response = rtf_pb2.RTFResponse()
            if output_value:
                response.body = bytes(output_value)
            response.status = True
            response_q.put(response)
            stop = True

        def stdout_sender():
            while not fp.closed:
                try:
                    line = fp.readline()
                    response = rtf_pb2.RTFResponse()
                    response.stdout = line
                    response.status = True
                    response_q.put(response)
                except IndexError as e:
                    pass

        threads = [
            threading.Thread(target=stdout_sender),
            threading.Thread(target=executor),
        ]
        for thread in threads:
            thread.start()

        while True:
            if stop:
                while not response_q.empty():
                    yield response_q.get()
                break
            yield response_q.get()
