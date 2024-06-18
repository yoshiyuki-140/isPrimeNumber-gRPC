package main

import (
	"fmt"
	"log"
	"net"

	"github.com/yoshiyuki-140/isprimenumber-gRPC-go/isPrime"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Go gRPC Prime judgement server!")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 判定サーバー
	s := isPrime.Server{}

	// サーバー作成
	grpcServer := grpc.NewServer()

	isPrime.RegisterIsPrimeFuncServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server: %s", err)
	}
}
