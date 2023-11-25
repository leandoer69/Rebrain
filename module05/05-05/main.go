package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	k1   = "key1"
	step = 7
)

func main() {
	cache := Cache{storage: make(map[string]int)}
	semaphore := make(chan int, 4)
	wg := &sync.WaitGroup{}
	wg.Add(20)

	for i := 0; i < 10; i++ {
		semaphore <- step
		go func() {
			defer wg.Done()
			cache.Increase(k1, <-semaphore)
			time.Sleep(time.Millisecond * 100)
		}()
	}

	for i := 0; i < 10; i++ {
		semaphore <- i * step
		go func() {
			defer wg.Done()
			value := <-semaphore
			fmt.Println(value)
			cache.Set(k1, value)
			time.Sleep(time.Millisecond * 100)
		}()
	}
	wg.Wait()

	fmt.Println(cache.Get(k1))
}
