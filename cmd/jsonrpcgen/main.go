package main

import (
	"flag"
	"log"
	"os"
	"path"

	"github.com/json-rpc-ecosystem/json-rpc/spec"
)

func main() {
	specFile := flag.String("spec-file", "", "JSON-RPC spec file")
	browserOutDir := flag.String("browser-out-dir", "", "directory of the generate browser source")
	goOutDir := flag.String("go-out-dir", "", "directory of the generate Go source")
	nodeOutDir := flag.String("node-out-dir", "", "directory of the generate Node source")
	flag.Parse()

	var def spec.Definition

	err := spec.DecodeFile(*specFile, &def)
	if err != nil {
		log.Fatal(err)
	}

	baseSpecFile := "/" + path.Base(*specFile)

	if *browserOutDir != "" {
		f, err := os.Create(*browserOutDir + baseSpecFile + ".browser.js")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		err = spec.GenerateBrowser(f, "./templates", &def)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Generated: ", *browserOutDir+baseSpecFile+".browser.js")
	}

	if *goOutDir != "" {
		f, err := os.Create(*goOutDir + baseSpecFile + ".go")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		err = spec.GenerateGo(f, "./templates", &def)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Generated: ", *goOutDir+baseSpecFile+".go")
	}

	if *nodeOutDir != "" {
		f, err := os.Create(*nodeOutDir + baseSpecFile + ".node.js")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		err = spec.GenerateNode(f, "./templates", &def)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Generated: ", *nodeOutDir+baseSpecFile+".node.js")
	}
}
