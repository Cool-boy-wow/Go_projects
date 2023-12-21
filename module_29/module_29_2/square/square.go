package square

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func Square(wg *sync.WaitGroup, flag chan os.Signal) {
	go func() {
		defer wg.Done()
		var i = 1
		for {

			select {
			case <-flag:
				fmt.Println("Выхожу из программы")
				return
			default:
				fmt.Println("Квадрат числа", i, "равен:", i*i)
				time.Sleep(time.Second * 2)
			}
			i += 1
		}
	}()
}
