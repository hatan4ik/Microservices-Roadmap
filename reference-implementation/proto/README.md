# Protobuf generation

This repo ships a small **stub** for readability. For real gRPC codegen:

1) Install:
- protoc
- protoc-gen-go
- protoc-gen-go-grpc

2) Generate:
```bash
protoc -I . orders.proto   --go_out=./ordersv1 --go_opt=paths=source_relative   --go-grpc_out=./ordersv1 --go-grpc_opt=paths=source_relative
```

Then update the Orders service import paths accordingly.
