package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 定义一个 intChan 通道，并向其写入数据
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	// 2. 定义一个 stringChan 通道，并向其写入数据
	stringChan := make(chan string, 10)
	for i := 0; i < 10; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	// 3. 使用 Select 多路复用同时从多个通道中获取数据（需配合 for 无限循环）
	for {
		select {
		case value := <-intChan:
			fmt.Println("从 intChan 中获取数据：", value)
			time.Sleep(time.Millisecond * 100)
		case value := <-stringChan:
			fmt.Println("从 stringChan 中获取数据：", value)
			time.Sleep(time.Millisecond * 100)
		default:
			fmt.Println("所有数据获取完毕！！！")
			return // 注意，跳出 for 循环
		}
	}

}
