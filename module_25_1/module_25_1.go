package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var a string
	var b string
	flag.StringVar(&b, "b", " ", "set b")
	flag.StringVar(&a, "a", " ", "set a")

	flag.Parse()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(strings.Contains(a, b))
}
