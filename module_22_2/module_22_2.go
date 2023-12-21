package main

import (
	"fmt"
	"math/rand"
	"time"
)

func find(a [n]int, value int) (index int) {
	index = -1
	for i := 0; i < n; i++ {
		if a[i] == value {
			index = i
			break
		}
	}
	return
}

const n = 12

func main() {
	rand.Seed(time.Now().UnixNano())
	var a [n]int
	l := 1
	for i := 0; i < n; i++ {
		if l%3 == 0 {
			l += 1
		}
		a[i] = l*i + rand.Intn(n)
	}
	fmt.Println(a)
	var value int
	fmt.Scan(&value)
	fmt.Println(find(a, value))
}
