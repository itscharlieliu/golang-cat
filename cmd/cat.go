package main

import (
	"flag"

	"github.com/itscharlieliu/golang-cat/pkg"
)

func main() {

	showLineNumPtr := flag.Bool("n", false, "show line numbers")
	hideEmptyLinesPtr := flag.Bool("s", false, "hide empty lines")

	flag.Parse()

	args := flag.Args()

	for i := 0; i < len(args); i++ {
		pkg.ReadFile(args[i], *showLineNumPtr, *hideEmptyLinesPtr)
	}

}
