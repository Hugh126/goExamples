直接使用了grpc官网demo，参考 [QuickStart](https://grpc.io/docs/languages/go/quickstart/)  

> 注意：如果是windows，以下步骤请使用GitBash执行

1. 按照protoc和grpc
```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

2. 设置环境变量（如果commond not found，可以设置中自行变更，然后验证）
```
export PATH="$PATH:$(go env GOPATH)/bin"
```

3. 启动Server(会弹窗要求同意端口访问)
```
go run grpc-go/helloworld/greeter_server/main.go
```

4. 启动Client(这里我变更了请求内容，你可以试试改改客户端或者服务端内容)
```
go run grpc-go/helloworld/greeter_client/main.go
```