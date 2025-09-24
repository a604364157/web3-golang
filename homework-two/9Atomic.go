package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
*/
func main() {
	var count = atomic.Int64{}
	var wg sync.WaitGroup
	g := 10
	num := 1000
	wg.Add(g)

	for range g {
		go func() {
			defer wg.Done()
			for range num {
				count.Add(1)
			}
		}()
	}

	wg.Wait()
	fmt.Println(count.Load())
}
