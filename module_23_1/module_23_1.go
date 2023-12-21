package main

import "fmt"

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		fmt.Scan(&a[i])
	}
	for i := 1; i < 10; i++ {
		x := a[i]
		j := i
		for j > 0 && a[j-1] > x {
			a[j] = a[j-1]
			j -= 1
		}
		a[j] = x
	}
	fmt.Println(a)

}
