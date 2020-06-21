# Step #1: Generate proto (*.pb.go) file // HL
# Syntax // HL
protoc -Ipath/to/protodir -Ipath/to/boilerplate --go_out=plugins=grpc:path/to/protodir \
  path/to/gateway.proto

# Example // HL
protoc -I. \
  -I$GOPATH/src/github.com/johanbrandhorst/grpc-gateway-boilerplate/third_party/googleapis \
  --go_out=plugins=grpc:. gateway.proto

# Step #2: Generate gateway (*.pb.gw.go) file // HL
# Syntax  // HL
protoc -I/usr/local/include -Ipath/to/protodir -Ipath/to/boilerplate \
  --grpc-gateway_out=logtostderr=true,allow_delete_body=true:path/to/protodir \
  path/to/gateway.proto

# Example // HL
protoc -I/usr/local/include \
  -I. \
  -I$GOPATH/src/github.com/johanbrandhorst/grpc-gateway-boilerplate/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true,allow_delete_body=true:. gateway.proto