package main

import (
	"fmt"
	"sync"

	"../pkg/Multi"
	"../pkg/Square"
)

func main() {
	var a int
	var s string
	for {
		var wg sync.WaitGroup
		wg.Add(2)
		fmt.Println("Введите число: ")
		fmt.Scan(&a)
		if a != 0 {
			sc := Square.Square(&wg, int(a))
			mc := Multi.Multi(&wg, sc)
			wg.Wait()
			fmt.Println("Произведение равно: ", <-mc)
		}

		fmt.Println("Если желаете закончить, введите 'стоп', иначе введите 'продолжить' ")
		fmt.Scan(&s)
		if s == "стоп" {
			break
		}
	}

}
