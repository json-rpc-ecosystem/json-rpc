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
	Version   string    `hcl:"version"`
	Namespace Namespace `hcl:"namespace,block"`
	Services  []Service `hcl:"service,block"`
}

type Namespace struct {
	Go     string `hcl:"go"`
	Java   string `hcl:"java"`
	Kotlin string `hcl:"kotlin"`
	Rust   string `hcl:"rust"`
	CPP    string `hcl:"cpp"`
	CSharp string `hcl:"csharp"`
	C      string `hcl:"c"`
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

func copyDefinition(dst *Definition, src *Definition) {
	dst.Version = src.Version
	dst.Namespace = src.Namespace
	dst.Services = make([]Service, len(src.Services))

	for sidx := range src.Services {
		dst.Services[sidx].Name = src.Services[sidx].Name
		dst.Services[sidx].Description = src.Services[sidx].Description
		dst.Services[sidx].Endpoint = src.Services[sidx].Endpoint
		dst.Services[sidx].Methods = make([]Method, len(src.Services[sidx].Methods))

		for midx := range src.Services[sidx].Methods {
			dst.Services[sidx].Methods[midx].Name = src.Services[sidx].Methods[midx].Name
			dst.Services[sidx].Methods[midx].Description = src.Services[sidx].Methods[midx].Description

			dst.Services[sidx].Methods[midx].Params = make(map[string]string)
			for field := range src.Services[sidx].Methods[midx].Params {
				dst.Services[sidx].Methods[midx].Params[field] = src.Services[sidx].Methods[midx].Params[field]
			}

			dst.Services[sidx].Methods[midx].Result = make(map[string]string)
			for field := range src.Services[sidx].Methods[midx].Result {
				dst.Services[sidx].Methods[midx].Result[field] = src.Services[sidx].Methods[midx].Result[field]
			}
		}
	}
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

func GenerateBrowser(w io.Writer, dir string, def *Definition) error {
	_typeMap := map[string]string{
		"String":  "String",
		"Number":  "Number",
		"Boolean": "Boolean",

		"[]String":  "Array<String>",
		"[]Number":  "Array<Number>",
		"[]Boolean": "Array<Boolean>",
	}

	tmpl, err := template.ParseFiles(dir + "/rpc.browser.template")
	if err != nil {
		return err
	}

	var localdef Definition
	copyDefinition(&localdef, def)

	for sidx := range localdef.Services {
		for midx := range localdef.Services[sidx].Methods {
			for field, _type := range localdef.Services[sidx].Methods[midx].Params {
				localdef.Services[sidx].Methods[midx].Params[field] = _typeMap[_type]
			}
			for field, _type := range localdef.Services[sidx].Methods[midx].Result {
				localdef.Services[sidx].Methods[midx].Result[field] = _typeMap[_type]
			}
		}
	}

	err = tmpl.ExecuteTemplate(w, "rpc.browser.template", localdef)
	if err != nil {
		return err
	}

	return nil
}

func GenerateCSharp(w io.Writer, dir string, def *Definition) error {
	_typeMap := map[string]string{
		"String":  "string",
		"Number":  "double",
		"Boolean": "bool",

		"[]String":  "string[]",
		"[]Number":  "double[]",
		"[]Boolean": "bool[]",
	}

	tmpl, err := template.ParseFiles(dir + "/rpc.csharp.template")
	if err != nil {
		return err
	}

	var localdef Definition
	copyDefinition(&localdef, def)

	for sidx := range localdef.Services {
		for midx := range localdef.Services[sidx].Methods {
			for field, _type := range localdef.Services[sidx].Methods[midx].Params {
				localdef.Services[sidx].Methods[midx].Params[field] = _typeMap[_type]
			}
			for field, _type := range localdef.Services[sidx].Methods[midx].Result {
				localdef.Services[sidx].Methods[midx].Result[field] = _typeMap[_type]
			}
		}
	}

	err = tmpl.ExecuteTemplate(w, "rpc.csharp.template", localdef)
	if err != nil {
		return err
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

	var localdef Definition
	copyDefinition(&localdef, def)

	for sidx := range localdef.Services {
		for midx := range localdef.Services[sidx].Methods {
			for field, _type := range localdef.Services[sidx].Methods[midx].Params {
				localdef.Services[sidx].Methods[midx].Params[field] = _typeMap[_type]
			}
			for field, _type := range localdef.Services[sidx].Methods[midx].Result {
				localdef.Services[sidx].Methods[midx].Result[field] = _typeMap[_type]
			}
		}
	}

	err = tmpl.ExecuteTemplate(w, "rpc.go.template", localdef)
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

	var localdef Definition
	copyDefinition(&localdef, def)

	for sidx := range localdef.Services {
		for midx := range localdef.Services[sidx].Methods {
			for field, _type := range localdef.Services[sidx].Methods[midx].Params {
				localdef.Services[sidx].Methods[midx].Params[field] = _typeMap[_type]
			}
			for field, _type := range localdef.Services[sidx].Methods[midx].Result {
				localdef.Services[sidx].Methods[midx].Result[field] = _typeMap[_type]
			}
		}
	}

	err = tmpl.ExecuteTemplate(w, "rpc.node.template", localdef)
	if err != nil {
		return err
	}

	return nil
}

func GeneratePython(w io.Writer, dir string, def *Definition) error {
	_typeMap := map[string]string{
		"String":  `""`,
		"Number":  "0",
		"Boolean": "False",

		"[]String":  "[]",
		"[]Number":  "[]",
		"[]Boolean": "[]",
	}

	tmpl, err := template.ParseFiles(dir + "/rpc.python.template")
	if err != nil {
		return err
	}

	var localdef Definition
	copyDefinition(&localdef, def)

	for sidx := range localdef.Services {
		for midx := range localdef.Services[sidx].Methods {
			for field, _type := range localdef.Services[sidx].Methods[midx].Params {
				localdef.Services[sidx].Methods[midx].Params[field] = _typeMap[_type]
			}
			for field, _type := range localdef.Services[sidx].Methods[midx].Result {
				localdef.Services[sidx].Methods[midx].Result[field] = _typeMap[_type]
			}
		}
	}

	err = tmpl.ExecuteTemplate(w, "rpc.python.template", localdef)
	if err != nil {
		return err
	}

	return nil
}
