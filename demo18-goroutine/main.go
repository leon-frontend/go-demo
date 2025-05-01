package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 定义开启 goroutine 协程的函数
func goroutineTest1() {
	for i := 0; i < 5; i++ {
		fmt.Printf("goroutineTest1 => %d\n", i)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Done() // 协程计数器减 1
}

// 定义开启 goroutine 协程的函数
func goroutineTest2() {
	for i := 0; i < 5; i++ {
		fmt.Printf("goroutineTest2 => %d\n", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() // 协程计数器减 1
}

/*
1. main 函数就是主线程
2. 在主线程（可以理解成进程）中，开启一个goroutine，该协程每隔50毫秒秒输出“你好golang"
3. 在主线程中也每隔一秒输出"你好golang"，输出10次后，退出程序
4. 要求主线程和goroutine同时执行
*/
func main() {
	wg.Add(1)           // 协程计数器加 1
	go goroutineTest1() // 开启一个协程，让协程和主线程并发执行
	wg.Add(1)           // 协程计数器加 1
	go goroutineTest2() // 开启一个协程，让协程和主线程并发执行

	// for i := 0; i < 10; i++ {
	// 	fmt.Println("main => 你好golang")
	// 	time.Sleep(time.Millisecond * 50)
	// }

	wg.Wait()                  // 等待协程执行完毕后，再之后后面的代码
	fmt.Println("主线程完毕，退出...") // 若不使用 sync.WaitGroup 等待协程执行完毕，则主线程退出时，协程也会退出
	/*
		goroutineTest1 => 0
		goroutineTest2 => 0
		goroutineTest1 => 1
		... ...
		goroutineTest2 => 4
		主线程完毕，退出...
	*/
}
