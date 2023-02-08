package singleton

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func ClientOnce() {
	wg.Add(30)
	for i := 0; i < 30; i++ {
		go func() {
			getInstance()
			wg.Done()
		}()
	}
	wg.Wait()
}

func ClientLock() {

	for i := 0; i < 30; i++ {
		go getInstanceLock()
	}
	fmt.Scanln()
}
