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