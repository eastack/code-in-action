package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		c := tempconv.Celsius(t)

		fmt.Printf("%s = %s, \n%s = %s\n",
			c, tempconv.CToF(c),
			c, tempconv.CToK(c),
		)
	}
}
