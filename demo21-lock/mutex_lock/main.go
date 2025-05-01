package main

import (
	"fmt"
	"sync"
)

var x int64
var wg sync.WaitGroup
var lock sync.Mutex // 互斥锁

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock() // 加锁
		x = x + 2
		lock.Unlock() // 解锁
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x) // 应该输出 20000，但是如果不加互斥锁会输出小于 20000 的随机值
}
