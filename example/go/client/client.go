package main

import (
	"fmt"
	"net/http"

	"github.com/json-rpc-ecosystem/json-rpc/example/rpc"
)

func main() {
	c := rpc.NewClient(http.DefaultClient, "http://localhost:8080")

	fmt.Println(c.Greeter.SayHello(&rpc.GreeterSayHelloParams{From: "Blain", To: "Austin"}))
}
