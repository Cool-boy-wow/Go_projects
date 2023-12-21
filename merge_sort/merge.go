package main

import "fmt"

const sizeA = 4
const sizeB = 5

var c [9]int

func merge(a [sizeA]int, b [sizeB]int) [9]int {
	i := 0
	j := 0
	k := 0
	for i < 4 && j < 5 {
		if a[i] < b[j] {
			c[k] = a[i]
			i += 1
		} else {
			c[k] = b[j]
			j += 1
		}
		k += 1
	}
	if i == 4 {
		for p := j; p < 9; p++ {
			c[p] = b[j]
		}
	} else {
		for p := k; p < 9; p++ {
			c[p] = a[i]
			i += 1
		}
	}

	return c
}

func main() {
	a := [sizeA]int{0, 1, 7, 8}
	b := [sizeB]int{2, 3, 4, 5, 6}
	c = merge(a, b)
	fmt.Println(c)
}
