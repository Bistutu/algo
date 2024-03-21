package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// 启动一个子线程
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println("子线程执行中...", i)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// 主线程执行其他操作
	fmt.Println("主线程开始执行")
	time.Sleep(2 * time.Second)
	fmt.Println("主线程执行完毕")

	// 等待子线程执行完毕
	wg.Wait()
	fmt.Println("程序退出")
}
