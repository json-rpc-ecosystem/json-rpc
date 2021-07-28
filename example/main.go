package main

import (
	"log"
	"math"
	"net/http"

	"github.com/json-rpc-ecosystem/json-rpc/example/rpc"
)

type arithService struct{}

func (s *arithService) Add(params *rpc.ArithAddParams, result *rpc.ArithAddResult) error {
	return nil
}

func (s *arithService) Pow(params *rpc.ArithPowParams, result *rpc.ArithPowResult) error {
	result.Num = math.Pow(params.Base, params.Pow)
	return nil
}

func (s *arithService) IsNegative(params *rpc.ArithIsNegativeParams, result *rpc.ArithIsNegativeResult) error {
	return nil
}

type greeterService struct{}

func (s *greeterService) SayHello(params *rpc.GreeterSayHelloParams, result *rpc.GreeterSayHelloResult) error {
	result.Message = "HELLO"
	return nil
}

func main() {
	server := rpc.Server{
		ArithService:   &arithService{},
		GreeterService: &greeterService{},
	}

	err := http.ListenAndServe(":8080", &server)
	if err != nil {
		log.Fatal(err)
	}
}
