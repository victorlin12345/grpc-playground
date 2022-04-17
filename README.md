# GRPC Playground

## Prerequisites
### protocol buffers :
  - Define : A message format mechanism light and easy to interchange.
  - Manual :  
    - Define a `.proto` file for message format.
    - Use the proto buffer compiler to generate message accessing structure for specific language ( eg. golang, C# ).
      - proto buffer compiler : `brew install protobuf`
      - specific language plugin for compiler to generate : `go get google.golang.org/protobuf/cmd/protoc-gen-go@v1.26`
    - Use proto buffer specific language library to read, write message. ( Not need for grpc )
      - package in golang `google.golang.org/protobuf/proto`
      - transform proto to binary wire format. `proto.Marshal(m Message)`
### grpc :
  - Define : In gRPC, a client application can directly call a method on a server application on a different machine as if it were a local object on client side. Its usage looks fine in micro service. Also it will can generate its rpc code by package.
  - gRPC vs REST API design : https://cloud.google.com/blog/products/api-management/understanding-grpc-openapi-and-rest-and-when-to-use-them  

## quick start
1. proto buffer compiler : `brew install protobuf`
2. go module init : `go mod init <project-name>`
3. go plugin for pb compiler : `go get google.golang.org/protobuf/cmd/protoc-gen-go@v1.26`
4. for generate grpc code : `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2`
5. generate order proto buffer and grpc go file :
   ```
   // --go_out: gen protobuffer go code init location (.)
   // --go_opt=paths=source_relative: the output file is placed in the same 
   // relative directory as the input file. ( here is order )
   // --go-grpc_out: gen grpc go code init location (.)
   //  --go-grpc_opt=paths=source_relative: the output file is placed in the same 
   // relative directory as the input file. ( here is order )
   protoc --go_out=. --go_opt=paths=source_relative \
   --go-grpc_out=. --go-grpc_opt=paths=source_relative \
   order/order.proto
   ```
6. run up server : `go run cmd/server/server.go`
7. run up client : `go run cmd/client/client.go`

expected clint message 
```
2021/06/15 20:58:04 Request Status: 200
2021/06/15 20:58:04 Response Message: Purchase Success!, Total price 20210.
```

expected server message
```
2021/06/15 20:58:00 Server listening on 127.0.0.1:9999
2021/06/15 20:58:04 Received Account victorlin54321 buy YOU ARE JUST A TOY item amount: 10
```