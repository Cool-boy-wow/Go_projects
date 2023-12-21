package main

import (
	"fmt"
	"strings"
)

func parseTest(a []string, b []rune) {
	var slice string
	index := -2
	p := -2
	for _, v := range a {
		fmt.Println(v)
		index = strings.Index(v, " ")
		slice = v[index+1:]
		for _, l := range b {
			slice = strings.ToUpper(slice)
			p = strings.LastIndex(slice, string(l))
			if p != -1 {
				fmt.Println(string(l), "Position:", p)
			}
		}
	}
}

func main() {
	var a = []string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
	var b = []rune{'H', 'E', 'L', 'П', 'М'}
	parseTest(a, b)
}
