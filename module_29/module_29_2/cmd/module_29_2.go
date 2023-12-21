package main

import (
	"os"
	"os/signal"
	"sync"

	"../square"
)

func main() {
	flag := make(chan os.Signal, 1)
	signal.Notify(flag, os.Interrupt)

	var wg sync.WaitGroup
	wg.Add(1)
	square.Square(&wg, flag)
	wg.Wait()
}
