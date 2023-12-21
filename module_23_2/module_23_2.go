package main

import "fmt"

func main() {
	f := func(a ...int) bool {
		if len(a) < 1 {
			return false
		}
		for i := 0; i < len(a)-1; i++ {
			for j := i + 1; j < len(a); j++ {
				if a[i] < a[j] {
					a[i], a[j] = a[j], a[i]
				}
			}
		}
		fmt.Println(a)
		return true
	}
	fmt.Println(f(1, 3, 3, 34, 5))
	fmt.Println(f(1, 3, 6))
	fmt.Println(f(1, 3))
	fmt.Println(f(1))
	fmt.Println(f())
}
