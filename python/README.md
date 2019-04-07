# Remote Tensorflow Call (RTF)

wip

## 1. Proto

From parent directory:

```
python -m grpc_tools.protoc -I proto/ --python_out=python/ --grpc_python_out=python/ proto/rtf.proto
```

## 2. gRPC server

```
python main.py
```
