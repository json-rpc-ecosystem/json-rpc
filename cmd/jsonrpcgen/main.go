package main

import (
	"flag"
	"log"
	"os"

	"github.com/json-rpc-ecosystem/json-rpc/spec"
)

func main() {
	specFile := flag.String("spec-file", "", "JSON-RPC spec file")
	browserOutFile := flag.String("browser-out-file", "", "file of the generate browser source")
	csharpOutFile := flag.String("csharp-out-file", "", "file of the generate C# source")
	goOutFile := flag.String("go-out-file", "", "file of the generate Go source")
	nodeOutFile := flag.String("node-out-file", "", "file of the generate Node source")
	pythonOutFile := flag.String("python-out-file", "", "file of the generate Python source")
	flag.Parse()

	var def spec.Definition

	err := spec.DecodeFile(*specFile, &def)
	if err != nil {
		log.Fatal(err)
	}

	if *browserOutFile != "" {
		f, err := os.Create(*browserOutFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		err = spec.GenerateBrowser(f, "./templates", &def)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Generated: ", *browserOutFile)
	}

	if *csharpOutFile != "" {
		f, err := os.Create(*csharpOutFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		err = spec.GenerateCSharp(f, "./templates", &def)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Generated: ", *csharpOutFile)
	}

	if *goOutFile != "" {
		f, err := os.Create(*goOutFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		err = spec.GenerateGo(f, "./templates", &def)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Generated: ", *goOutFile)
	}

	if *nodeOutFile != "" {
		f, err := os.Create(*nodeOutFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		err = spec.GenerateNode(f, "./templates", &def)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Generated: ", *nodeOutFile)
	}

	if *pythonOutFile != "" {
		f, err := os.Create(*pythonOutFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		err = spec.GeneratePython(f, "./templates", &def)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Generated: ", *pythonOutFile)
	}
}
