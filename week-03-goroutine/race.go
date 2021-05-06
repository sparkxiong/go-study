package main

import (
	"fmt"
	"sync"
)

//线程并发处理数据。 data race
func main() {
	a := 0

	var wait sync.WaitGroup

	var mtx sync.Mutex
	wait.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			mtx.Lock()
			defer mtx.Unlock()
			defer wait.Done()
			a++
		}()
	}
	wait.Wait()
	fmt.Println("a: ", a)
}
