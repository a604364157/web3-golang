package main

import (
	"fmt"
	"sync"
)

func useLock() {
	var count int
	var mu sync.Mutex
	var wg sync.WaitGroup

	g := 10
	num := 1000
	wg.Add(10)
	for range g {
		go func() {
			defer wg.Done()
			for range num {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println("使用锁的count:", count)
}

func noLock() {
	var count int
	var wg sync.WaitGroup

	g := 10
	num := 1000
	wg.Add(10)
	for range g {
		go func() {
			defer wg.Done()
			for range num {
				count++
			}
		}()
	}
	wg.Wait()
	fmt.Println("不使用锁的count:", count)
}

/*
编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
*/
func main() {
	useLock()
	noLock()
}
