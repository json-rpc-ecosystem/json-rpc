package main

import (
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/json-rpc-ecosystem/json-rpc/example/rpc"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		next.ServeHTTP(w, r)
	})
}

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

	err := http.ListenAndServe(":8080", corsMiddleware(server.Mux()))
	if err != nil {
		log.Fatal(err)
	}
}
