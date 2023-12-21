package Square

import "sync"

func Square(wg *sync.WaitGroup, a int) int {

	firstChan := make(chan int)
	go func() {
		firstChan <- a * a
	}()
	defer wg.Done()
	return <-firstChan
}
