package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印
*/
func main() {
	ch := make(chan int, 100) // 定义一个带有缓冲的通道，容量为10
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 120; i++ {
			// 缓存100，这里生产者协程发送120个整数，超过100个后，会阻塞
			ch <- i
		}
		close(ch) // 关闭通道
		fmt.Println("生产者协程已关闭通道")
	}()
	go func() {
		defer wg.Done()
		for num := range ch {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			fmt.Println("收到消息：", num)
		}
	}()
	wg.Wait()
}
