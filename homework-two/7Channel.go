package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
编写一个程序，使用通道实现两个协程之间的通信。
一个协程生成从1到10的整数，并将这些整数发送到通道中，
另一个协程从通道中接收这些整数并打印出来
*/
func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			ch <- i
		}
		close(ch) // 关闭通道
	}()
	go func() {
		defer wg.Done()
		for num := range ch {
			fmt.Println("收到消息：", num)
		}
	}()
	wg.Wait()
}
