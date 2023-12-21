package main

import "fmt"

func function(a []int) (b []int, c []int) {
	for _, v := range a {
		if v%2 == 0 {
			b = append(b, v)
		} else {
			c = append(c, v)
		}
	}
	return
}

func main() {
	var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b, c := function(a)
	fmt.Println("Четные числа", b)
	fmt.Println("Нечетные числа", c)
}
