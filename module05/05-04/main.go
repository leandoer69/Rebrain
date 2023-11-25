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
	wg := &sync.WaitGroup{}

	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			cache.Increase(k1, step)
			time.Sleep(time.Millisecond * 100)
		}()
	}

	for i := 0; i < 10; i++ {
		i := i
		go func() {
			defer wg.Done()
			cache.Set(k1, i*step)
			time.Sleep(time.Millisecond * 100)
		}()
	}

	wg.Wait()
	fmt.Println(cache.Get(k1))
}
