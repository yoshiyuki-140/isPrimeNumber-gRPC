package main

import (
	"context"
	"log"

	"github.com/yoshiyuki-140/isprimenumber-gRPC-go/server/isPrime"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	// IPアドレスを指定するときは以下のコードの":9000"となっているところを"127.0.0.1:9000"等に変える(この例ではlocalhost)
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	p := isPrime.NewIsPrimeFuncClient(conn)

	// チェックしたい数字
	var num int64 = 100
	response, err := p.CheckPrime(context.Background(), &isPrime.Value{Value: num})
	if err != nil {
		log.Fatalf("Error when calling CheckPrime: %s", err)
	}
	if response.IsPrime {
		log.Printf("T")
	} else {
		log.Printf("F")
	}
}
