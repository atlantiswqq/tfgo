"""Remote Tensorflow Execution, gRCP server."""

from concurrent import futures
import sys
import grpc
import rtf_pb2_grpc
from rtf import RTFServicer
import time


def serve():
    """Serving function."""
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    rtf_pb2_grpc.add_RTFServicer_to_server(RTFServicer(), server)
    server.add_insecure_port("[::]:50051")
    server.start()
    while True:
        time.sleep(1)
    return 0


if __name__ == "__main__":
    sys.exit(serve())
