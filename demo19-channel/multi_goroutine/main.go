package main

import (
	"fmt"
	"sync"
	"time"
)

/*
需求：定义两个方法，一个方法给管道里面写数据，一个给管道里面读取数据。要求同步进行。
1、开启一个 write 的协程给向管道 ch 中写入 100 条数据
2、开启一个 read 的协程读取 ch 中写入的数据
3、注意：write 和 read 同时操作一个管道
4、主线程必须等待操作完成后才可以退出
*/

// 声明一个全局变量用于控制主线程的结束时机
var wg sync.WaitGroup

func write(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Printf("写入数据：%d\n", i)
		time.Sleep(time.Millisecond * 500)
	}
	// 使用 for range 读取管道之前必须关闭管道。否则报错。
	close(ch) // 由于 ch 是引用数据类型，因此修改形参的值，也会影响原变量的值。
	wg.Done() // 协程计数器减 1
}

func read(ch chan int) {
	// 使用 for range 读取管道之前必须关闭管道。否则报错。
	for value := range ch {
		fmt.Printf("读取数据：%d\n", value) // 当读取不到数据时，会等待其他协程写入数据后再读取
		time.Sleep(time.Millisecond * 10)
	}
	wg.Done() // 协程计数器减 1
}

func main() {
	// 1. 创建一个 channel 管道
	ch := make(chan int, 10)

	// 2. 开启两个 goroutine 协程
	wg.Add(2) // 协程计数器加 2
	go write(ch)
	go read(ch)

	// 3. 等待两个 goroutine 执行完毕之后，再关闭主线程
	wg.Wait()
	fmt.Println("主线程完毕，退出...")
}

/*
输出：
写入数据：0
读取数据：0
读取数据：1
写入数据：1
... ...
写入数据：9
读取数据：9
*/
