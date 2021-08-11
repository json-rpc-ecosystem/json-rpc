package main

import (
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/json-rpc-ecosystem/json-rpc/example/rpc"
)

type arithService struct{}

func (s *arithService) Add(params *rpc.ArithAddParams) (*rpc.ArithAddResult, error) {
	return nil, nil
}

func (s *arithService) Pow(params *rpc.ArithPowParams) (*rpc.ArithPowResult, error) {
	result := &rpc.ArithPowResult{}
	result.Num = math.Pow(params.Base, params.Pow)
	return result, nil
}

func (s *arithService) IsNegative(params *rpc.ArithIsNegativeParams) (*rpc.ArithIsNegativeResult, error) {
	return nil, nil
}

type greeterService struct{}

func (s *greeterService) SayHello(params *rpc.GreeterSayHelloParams) (*rpc.GreeterSayHelloResult, error) {
	result := &rpc.GreeterSayHelloResult{}
	result.Message = "HELLO"
	return result, nil
}

func main() {
	server := rpc.Server{
		ArithService:   &arithService{},
		GreeterService: &greeterService{},
	}

	go func() {
		err := http.ListenAndServe(":8080", &server)
		if err != nil {
			log.Fatal(err)
		}
	}()

	c := rpc.NewClient(http.DefaultClient, "http://localhost:8080")

	fmt.Println(c.Greeter.SayHello(&rpc.GreeterSayHelloParams{From: "Blain", To: "Austin"}))
}
