package main

import (
	"os"

	"github.com/alejandro31118/slb-lang/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
