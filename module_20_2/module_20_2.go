package main

import "fmt"

const rows_a = 3
const cols_a = 5

const rows_b = 5
const cols_b = 4

func mult(a [rows_a][cols_a]int, b [rows_b][cols_b]int) [rows_a][cols_b]int {
	var c = [rows_a][cols_b]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	p, k := 0, 0
	s := 0
	for k != rows_a {
		for i := 0; i < len(a); i++ {
			for l := 0; l < cols_b; l++ {
				for j := 0; j < len(a[i]); j++ {
					s += a[i][j] * b[j][l]
				}
				c[k][p] = s
				s = 0
				p += 1
				if p == cols_b {
					k += 1
					p = 0
				}
			}
		}
	}

	return c
}

func main() {
	var a = [rows_a][cols_a]int{
		{1, 2, 3, 4, 5},
		{2, 3, 4, 5, 6},
		{3, 4, 5, 6, 7},
	}
	var b = [rows_b][cols_b]int{
		{4, 5, 6, 7},
		{1, 2, 3, 4},
		{2, 3, 4, 5},
		{3, 2, 1, 2},
		{2, 3, 2, 1},
	}
	fmt.Println(mult(a, b))
}
