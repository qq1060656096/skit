## grpc

```
# 生成go pb 文件
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
# 生成go grpc service
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1



# 生成 go pb文件
protoc --proto_path=./ --go_out=./ --go_opt=paths=source_relative  order.proto

# 生成 php pb 文件
protoc --php_out=./ order.proto
protoc --php-grpc_out=./ ./order.proto

# 生成 pb 文件和 grpc 服务端
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative account/account.proto
protoc --go_out=. --go-grpc_out=. ./demo/account/account.proto
# 生成 go pb 代码
protoc --go_out=./ ./account/account.proto
# 生成 grpc 服务端代码
protoc --go-grpc_out=. ./account/account.proto

```

# 创建 grpc 服务端
```go
package main

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"demo/account/v1"
	"log"
	"net"
)

type AccountService struct {
	v1.UnimplementedAccountServiceServer
}

func (a *AccountService) RegisterAccount(ctx context.Context, req * v1.RegisterAccountRequest) (*v1.Account, error) {
	if req.Pass != "1" {
		msg := "register password must (1), request password is " + req.Pass
		return nil , errors.New(msg)
	}

	account := &v1.Account{
		Name: req.Name,
		Nickname: req.Nickname,
	}
	return account, nil
}

func main()  {
	grpcServer := grpc.NewServer()
	v1.RegisterAccountServiceServer(grpcServer, new(AccountService))

	lis, err := net.Listen("tcp", ":5900")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
```

## grpc客户端
```
package main

import (
	"context"
	v1 "demo/account/v1"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"time"
)

func main() {
	grpcConn, err := grpc.Dial(":5900", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := v1.NewAccountServiceClient(grpcConn)
	rand.Seed(time.Now().Unix())
	res := &v1.RegisterAccountRequest{
		Name: fmt.Sprintf("demo%s", time.Now().Format("20060102150405")),
		Nickname: fmt.Sprintf("demo %s", time.Now().Format("2006-01-02 15:04:05")),
		Pass: fmt.Sprintf("%d", rand.Intn(2)),
	}
	resp , err := client.RegisterAccount(context.Background(), res)
	if err != nil {
		log.Fatal("grpc.rsponse: ", err)
	}
	fmt.Println("request: ", res)
	fmt.Println("response: ", resp)
}
```

```cmd
# 启动服务端
go run cmd/demo-grpc-server/main.go
# 启动客户端
go run cmd/demo-grpc-client/main.go
```