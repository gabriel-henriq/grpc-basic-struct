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

# Running with Evans or Postman

```bash
    yay -S evans
```

```bash
    evans -r repl
```
```bash
    call SayHello
```


```https
    https://blog.postman.com/postman-now-supports-grpc/
```