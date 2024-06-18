package isPrime

import (
	"context"
	"log"
)

type Server struct {
	UnimplementedIsPrimeFuncServer
}

func (s *Server) CheckPrime(ctx context.Context, in *Value) (*IsPrimeResponse, error) {
	log.Println(in.Value)
	result := isPrime(in.Value) // resultにはtrueかfalseが入る
	return &IsPrimeResponse{IsPrime: result}, nil
}

func isPrime(num int64) bool {
	// 素数判定プログラム
	if num <= 1 {
		return false
	}
	// 試し切りは√Nまでで十分であることを利用したアルゴリズム
	// 参考:https://algo-method.com/descriptions/93
	for i := int64(2); i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
