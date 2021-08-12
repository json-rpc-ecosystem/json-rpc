package spec

import (
	"io"
	"io/ioutil"
	"text/template"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

type Definition struct {
	Version  string    `hcl:"version"`
	Services []Service `hcl:"service,block"`
}

type Service struct {
	Name        string   `hcl:"name,label"`
	Description string   `hcl:"description"`
	Endpoint    string   `hcl:"endpoint"`
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

func GenerateGo(w io.Writer, dir string, def *Definition) error {
	_typeMap := map[string]string{
		"String":  "string",
		"Number":  "float64",
		"Boolean": "bool",

		"[]String":  "[]string",
		"[]Number":  "[]float64",
		"[]Boolean": "[]bool",
	}

	tmpl, err := template.ParseFiles(dir + "/rpc.go.template")
	if err != nil {
		return err
	}

	for sidx := range def.Services {
		for midx := range def.Services[sidx].Methods {
			for field, _type := range def.Services[sidx].Methods[midx].Params {
				def.Services[sidx].Methods[midx].Params[field] = _typeMap[_type]
			}
			for field, _type := range def.Services[sidx].Methods[midx].Result {
				def.Services[sidx].Methods[midx].Result[field] = _typeMap[_type]
			}
		}
	}

	err = tmpl.ExecuteTemplate(w, "rpc.go.template", def)
	if err != nil {
		return err
	}

	return nil
}

func GenerateNode(w io.Writer, dir string, def *Definition) error {
	_typeMap := map[string]string{
		"String":  "String",
		"Number":  "Number",
		"Boolean": "Boolean",

		"[]String":  "Array<String>",
		"[]Number":  "Array<Number>",
		"[]Boolean": "Array<Boolean>",
	}

	tmpl, err := template.ParseFiles(dir + "/rpc.node.template")
	if err != nil {
		return err
	}

	for sidx := range def.Services {
		for midx := range def.Services[sidx].Methods {
			for field, _type := range def.Services[sidx].Methods[midx].Params {
				def.Services[sidx].Methods[midx].Params[field] = _typeMap[_type]
			}
			for field, _type := range def.Services[sidx].Methods[midx].Result {
				def.Services[sidx].Methods[midx].Result[field] = _typeMap[_type]
			}
		}
	}

	err = tmpl.ExecuteTemplate(w, "rpc.node.template", def)
	if err != nil {
		return err
	}

	return nil
}
