package main

import "fmt"

func change(x int16, y uint8, z float32) float32 {
	return float32(x)*2 + float32(y*y) - 3/z
}

func main() {
	var x int16
	var y uint8
	var z float32
	fmt.Scan(&x, &y, &z)
	S := change(x, y, z)
	fmt.Println(S)
}
