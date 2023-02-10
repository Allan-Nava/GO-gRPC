# GO-gRPC

Regenerate gRPC code
Before you can use the new service method, you need to recompile the updated .proto file.

While still in the examples/helloworld directory, run the following command:


```

$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto

```