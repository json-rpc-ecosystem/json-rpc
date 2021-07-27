package spec

import (
	"io/ioutil"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

type Definition struct {
	Version  string    `hcl:"version"`
	Endpoint string    `hcl:"endpoint"`
	Services []Service `hcl:"service,block"`
}

type Service struct {
	Name        string   `hcl:"name,label"`
	Description string   `hcl:"description"`
	Methods     []Method `hcl:"method,block"`
}

type Method struct {
	Name        string            `hcl:"name,label"`
	Description string            `hcl:"description"`
	Params      map[string]string `hcl:"params"`
	Result      map[string]string `hcl:"result"`
}

func DecodeFile(filename string, def *Definition) error {
	src, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	file, diags := hclsyntax.ParseConfig(src, filename, hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		return diags
	}

	diags = gohcl.DecodeBody(file.Body, nil, def)
	if diags.HasErrors() {
		return diags
	}

	return nil
}
