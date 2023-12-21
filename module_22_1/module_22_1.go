package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 10

func find(a [n]int, value int) int {
	for i := 0; i < n; i++ {
		if a[i] == value {
			return n - 1 - i
		}
	}
	return 0
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var a [n]int
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(n * 2)
	}
	fmt.Println(a)
	var value int
	fmt.Print("Введите число : ")
	fmt.Scan(&value)
	fmt.Println(find(a, value))
}
