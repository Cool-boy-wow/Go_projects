package Multi

import (
	"fmt"
	"sync"
)

func Multi(wg *sync.WaitGroup, fc int) chan int {
	secondChan := make(chan int)
	fmt.Println("Квадрат числа равен :", fc)
	go func() {
		secondChan <- 2 * fc
	}()
	defer wg.Done()
	return secondChan
}
