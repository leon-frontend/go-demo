package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func sayHello() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 50)
		fmt.Println("Hello, World!")
	}

	wg.Done()
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic from test():", err)
			wg.Done()
		}
	}()

	// 定义一个 Map
	var myMap map[int]string // 只声明了 map 类型的变量，未分配存储空间
	myMap[1] = "golang"      // 代码会报错

	wg.Done()
}

func main() {
	// 开启两个 goroutine 协程
	wg.Add(2)
	go sayHello()
	go test()

	wg.Wait()
	fmt.Println("main goroutine is done ....")
}
