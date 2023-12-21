package main

import "fmt"

func some(a, b int, A func(a int, b int) int) int {
	return A(a, b) + 7
}

func main() {
	f1 := some(7, 6, func(a, b int) int {
		fmt.Print("7**2=")
		return a * b
	})
	fmt.Println(f1)
	f2 := some(2, 9, func(a, b int) int {
		fmt.Print("7-7=")
		return a - b
	})
	fmt.Println(f2)
	f3 := some(3, 4, func(a, b int) int {
		fmt.Print("7*2=")
		return a + b
	})
	fmt.Println(f3)
}
