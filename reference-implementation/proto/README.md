# Protobuf generation

This repo ships a small **stub** (`ordersv1/stub.go`) for readability. For real gRPC codegen, generate Go types from `orders.proto`.

## Install tooling

You need:
- `protoc` (the Protocol Buffers compiler)
- `protoc-gen-go`
- `protoc-gen-go-grpc`

For Go plugins:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Make sure `$GOPATH/bin` (or `$HOME/go/bin`) is on your `PATH`.

## Generate code

From this folder:

```bash
protoc -I . orders.proto \
  --go_out=./ordersv1 --go_opt=paths=source_relative \
  --go-grpc_out=./ordersv1 --go-grpc_opt=paths=source_relative
```

If you replace the stub with generated code, re-run `go test` and ensure imports still reference `reference-implementation/proto/ordersv1`.
