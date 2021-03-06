package main

import (
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/json-rpc-ecosystem/json-rpc/example/rpc"
)

type arith struct{}

func (s *arith) Add(params *rpc.ArithAddParams) (*rpc.ArithAddResult, error) {
	return nil, nil
}

func (s *arith) Pow(params *rpc.ArithPowParams) (*rpc.ArithPowResult, error) {
	result := &rpc.ArithPowResult{}
	result.Num = math.Pow(params.Base, params.Pow)
	return result, nil
}

func (s *arith) IsNegative(params *rpc.ArithIsNegativeParams) (*rpc.ArithIsNegativeResult, error) {
	return nil, nil
}

type greeter struct{}

func (s *greeter) SayHello(params *rpc.GreeterSayHelloParams) (*rpc.GreeterSayHelloResult, error) {
	result := &rpc.GreeterSayHelloResult{}
	result.Message = fmt.Sprintf("Dear %s, Someone named %s says 'hello!'", params.To, params.From)
	return result, nil
}

func main() {
	server := rpc.Server{
		Arith:   &arith{},
		Greeter: &greeter{},
	}

	err := http.ListenAndServe(":8080", server.Mux())
	if err != nil {
		log.Fatal(err)
	}
}
