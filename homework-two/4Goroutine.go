package main

import (
	"fmt"
	"sync"
	"time"
)

func runTasks(tasks []func()) {
	var wg sync.WaitGroup
	wg.Add(len(tasks))
	for i, task := range tasks {
		go func(id int, t func()) {
			defer wg.Done()
			start := time.Now()
			t()
			elapsed := time.Since(start)
			fmt.Printf("任务 %d 执行完成，用时 %v\n", id+1, elapsed)
		}(i, task)
	}
	wg.Wait()
	fmt.Println("所有任务执行完毕")
}

/*
设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间
*/
func main() {
	tasks := []func(){
		func() {
			time.Sleep(2 * time.Second)
			fmt.Println("任务1完成")
		},
		func() {
			time.Sleep(1 * time.Second)
			fmt.Println("任务2完成")
		},
		func() {
			time.Sleep(3 * time.Second)
			fmt.Println("任务3完成")
		},
	}
	// 启动调度器
	runTasks(tasks)
}
