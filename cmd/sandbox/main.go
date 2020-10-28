package main

import "github.com/dominictwlee/bao/internal/fs"

func main() {
	err := fs.CopyDir("../../templates", "templates")
	if err != nil {
		panic(err)
	}
}
