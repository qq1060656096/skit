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