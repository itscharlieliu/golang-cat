package main

import (
	"os"

	"github.com/itscharlieliu/golang-cat/pkg"
)

func main() {

	args := os.Args[1:]

	for i := 0; i < len(args); i++ {
		pkg.ReadFile(args[i])
	}

}
