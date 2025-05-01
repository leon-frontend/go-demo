package main

import "fmt"

func main() {
	// 1. 创建一个管道(引用类型)
	ch1 := make(chan int, 3)
	// 2. 给 ch1 管道写入数据（发送操作）
	ch1 <- 10
	ch1 <- 20
	// 3. 从 ch1 管道读取数据（接收操作）
	num := <-ch1                       // 取出的值为 10，因为遵循先进先出原则
	fmt.Printf("从管道中取出的数据为 %v\n", num) // 从管道中取出的数据为 10
	// 4. 获取管道的值、容量和长度
	fmt.Printf("管道的值为 %v, 容量为 %v, 长度为 %v\n", ch1, cap(ch1), len(ch1)) // 管道的值为 0x1400013a000, 容量为 3, 长度为 1

	// 5. 管道阻塞
	// ch2 := make(chan int, 1)
	// ch2 <- 100
	// ch2 <- 200 // 管道阻塞，报错 fatal error: all goroutines are asleep - deadlock!
	ch3 := make(chan int, 1)
	a := <-ch3 // 管道阻塞，报错 fatal error: all goroutines are asleep - deadlock!
	fmt.Println(a)

	// 6. 循环遍历管道取值
	ch4 := make(chan int, 3)
	for i := 0; i < 3; i++ {
		ch4 <- i
	}

	// 6.1 使用 for range 遍历管道取值
	close(ch4) // 使用 for range 遍历管道取值之前，必须先关闭管道。否则报错。
	for i := range ch4 {
		fmt.Println(i) // 0 1 2
	}

	// 6.2 使用 for 循环遍历管道取值
	// for 循环不会报错的原因是不会产生对空管道取值的操作，因为管道只有 3 个数据，for 循环只循环 3 次。
	for i := 0; i < 3; i++ {
		fmt.Println(<-ch4) // 0 1 2
	}

	// 7. 单向管道
	// 7.1 只读管道
	// readChan := make(<-chan int, 3)
	// readChan <- 100 //  invalid operation: cannot send to receive-only channel readChan (variable of type <-chan int)

	// 7.2 只写管道
	writeChan := make(chan<- int, 3)
	writeChan <- 100
	// <-writeChan // invalid operation: cannot receive from send-only channel writeChan (variable of type chan<- int)
}
