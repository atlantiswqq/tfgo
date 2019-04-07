# Remote Tensorflow Call - Go client
wip

From parent folder, generate the proto files:

```
protoc -I proto/ --go_out=plugins=grpc:rtf/ proto/rtf.proto 
```
