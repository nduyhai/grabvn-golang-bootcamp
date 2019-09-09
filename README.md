# Feedback service with gateway

# Setup

## generate proto

```
$GOPATH/pkg/mod/github.com/golang/protobuf@v1.3.2/proto/protoc -I/usr/local/include -I.  -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis   --go_out=plugins=grpc:internal/bootcamp/feedback
```
