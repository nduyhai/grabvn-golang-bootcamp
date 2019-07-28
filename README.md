# Self Improvement Week 3

## Proto
* [Proto syntax](https://developers.google.com/protocol-buffers/docs/proto3)
* Generating your classes
```shell script
protoc -I api/ api/msg.proto --go_out=plugins=grpc:internal/bootcamp/msg
```
