# Basic Structure to Run gRPC API

```bash
    protoc --go-grpc_out=. --proto_path=app/proto app/proto/hello_message.proto
```

```bash
    protoc --go_out=. --proto_path=app/proto app/proto/hello_message.proto
```

```go
    go run .server/main
```