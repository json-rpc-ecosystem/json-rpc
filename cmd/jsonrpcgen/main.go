package main

import (
	"log"
	"os"
	"text/template"

	"github.com/json-rpc-ecosystem/json-rpc/spec"
)

func main() {
	var def spec.Definition

	err := spec.DecodeFile("./example.rpc", &def)
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.New("rpc.go.template").ParseFiles("./templates/rpc.go.template")
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(os.Stdout, def)
	if err != nil {
		log.Fatal(err)
	}
}
