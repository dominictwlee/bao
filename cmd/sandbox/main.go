package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dominictwlee/bao/internal/pkgjson"
)

func main() {
	pkg := pkgjson.PackageJSON{
		Name:    "dummy",
		Author:  "dom",
		Version: "0.1.0",
		License: "MIT",
		Main:    "index.js",
		Module:  "index.js.esm",
		Typings: "index.d.ts",
		Files:   []string{"something.js"},
		Engines: struct {
			Node string
		}{
			Node: ">=10",
		},
		Scripts: struct {
			Start string
			Build string
		}{
			Start: "bao start",
			Build: "bao build",
		},
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")

	err := enc.Encode(pkg)
	if err != nil {
		panic(err)
	}
	fmt.Println(buf.String())
}
