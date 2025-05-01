package main

import (
	"fmt"
	"sync"
)

// wg 全局变量用于控制主线程的结束时机
var wg sync.WaitGroup

// putNum 函数用于向 intChan 管道中存入数据，向 intChan 放入 1-120000 个数
func putNum(intChan chan int) {
	for i := 2; i < 100; i++ {
		intChan <- i // 向管道中写入数据
	}
	close(intChan) // 关闭 intChan 管道
	wg.Done()      // 协程计数器减 1
}

/*
 1. PrimeNum 函数的作用：从 intChan 取出数据，并判断是否为素数，如果是，就把得到的素数放在 primeChan
 2. 在 main 函数中会开启多个 primeNum 协程，同时该协程存在向 primeChan 管道写入数据的操作，
    因此需要考虑何时关闭 primeChan 管道的问题？使用 exitChan 管道标识 16 个 primeNum 协程是否都执行完毕。
*/
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	// 使用 for range 遍历 intChan 管道中的数据，因此 intChan 必须是关闭的。
	for num := range intChan {
		flag := true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag && num != 1 {
			primeChan <- num // 将素数 num 写入primeChan管道
		}
	}

	// 每个 primeNum 协程向 primeChan 管道写入数据完成后，向 exitChan 管道写入一个 true 值
	exitChan <- true // 用于标识当前 primeNum 协程执行完毕

	wg.Done() // 协程计数器减 1
}

// printPrime 函数用于读取 primeChan 管道中的素数，并打印素数的方法
func printPrime(primeChan chan int) {
	// 使用 for range 遍历  primeChan 管道中的数据之前需要关闭该管道。否则报错。
	for primeNum := range primeChan {
		fmt.Printf("printPrime 输出的素数为：%d\n", primeNum)
	}
	wg.Done() // 协程计数器减 1
}

func main() {
	// 1. 定义 Channel 管道
	// 管道的存取数据操作会交替进行，因此 1000 是一个比较合适的容量
	intChan := make(chan int, 1000)   // 定义存放 1- 120000 个数字的管道
	primeChan := make(chan int, 1000) // 定义存放素数的管道
	exitChan := make(chan bool, 16)   // exitChan 管道用于标识 16 个 primeChan 线程是否都执行完毕

	// 2. 开启一个 goroutine 协程，用于向 intChan 管道中存放数字
	wg.Add(1)
	go putNum(intChan)

	// 3. 开启多个 goroutine 协程，用于判断是否为素数
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go primeNum(intChan, primeChan, exitChan)
	}

	// 4. 开启一个 goroutine 协程，用于打印素数
	wg.Add(1)
	go printPrime(primeChan)

	// 5. 开启一个 goroutine 协程，判断 16 个 primeNum 协程是否都执行完毕
	wg.Add(1)
	go func() {
		// 必须循环 16 次，取出 exitChan 中的所有数据才能判断 16 个 primeNum 协程都执行完毕
		for i := 0; i < 16; i++ {
			<-exitChan // 从 exitChan 管道中读取数据
		}
		close(primeChan) // 等待 exitChan 管道中的 16 个数据取出之后，关闭 primeChan 管道
		wg.Done()        // 协程计数器减 1
	}()

	wg.Wait() // 主线程等待所有 goroutine 运行完毕
	fmt.Println("主线程执行完毕，退出...")
}
