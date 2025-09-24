package main

import (
	"fmt"
	"sync"
)

func printOdd(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 10 {
		if i%2 != 0 {
			fmt.Printf("奇数：%d\n", i)
		}
	}
}

func printEven(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 10 {
		if i%2 == 0 {
			fmt.Printf("偶数：%d\n", i)
		}
	}
}

/*
编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
*/
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go printOdd(&wg)
	go printEven(&wg)
	wg.Wait()
	fmt.Println("所有协程执行完毕")
}
