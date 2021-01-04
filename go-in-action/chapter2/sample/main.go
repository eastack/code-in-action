package main

import (
	"log"
	"os"

	_ "github.com/rwrr/goinaction/chapter2/sample/matchers"

	"github.com/rwrr/goinaction/chapter2/sample/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run(os.Args[1])
}
