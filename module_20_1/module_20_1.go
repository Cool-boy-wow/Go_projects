package main

import "fmt"

const rows_1 = 2
const cols_1 = 2

const rows_2 = 3
const cols_2 = 3

func determinant_1(a [rows_1][cols_1]int) int {
	m := 1
	n := 1
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if i == j {
				m *= a[i][j]
			} else {
				n *= a[i][j]
			}
		}
	}
	return m - n
}

func determinant_2(a [rows_2][cols_2]int) int {
	l := 0

	t := 0
	for i := 0; i < rows_2; i++ {
		var b = [rows_1][cols_1]int{
			{0, 0},
			{0, 0},
		}
		j := 0
		k := 0
		x := 0
		if i == 0 {
			for j != rows_1 {
				for k != rows_1 {
					b[j][k] = a[j+1][k+1]
					k += 1
				}
				k = 0
				j += 1
			}
		} else if i == 1 {
			for j != rows_1 {
				for k != rows_1 {
					if x != 1 {
						b[j][k] = a[j+1][k]
						k += 1
						x = 1
					} else {
						b[j][k] = a[j+1][cols_1]
						x = 0
						k += 1
					}
				}
				k = 0
				j += 1

			}
		} else {
			for j != rows_1 {
				for k != rows_1 {
					b[j][k] = a[j+1][k]
					k += 1
				}
				k = 0
				j += 1
			}
		}

		if i%2 == 0 {
			l += a[t][i] * determinant_1(b)

		} else {
			l -= a[t][i] * determinant_1(b)
		}
	}
	return l
}

func main() {
	var b = [rows_2][cols_2]int{
		{2, 3, 4},
		{4, 3, 2},
		{1, 1, 1},
	}

	fmt.Println(determinant_2(b))
}
