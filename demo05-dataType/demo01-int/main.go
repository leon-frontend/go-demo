package main

import (
	"fmt"
)

func main() {
	// 1. int8 有符号 8 位整型 (-128 到 127)
	// var a int8 = 99
	// fmt.Printf("num=%d 类型为%T\n", a, a) // num=99 类型为int8
	// fmt.Println(unsafe.Sizeof(a))      // 输出 1，表示占用 1 个字节

	// 2. int 长度类型转换
	var a int32 = 10
	var b int64 = 21
	// fmt.Println(a + b) // invalid operation: a + b (mismatched types int32 and int64)
	fmt.Println(int64(a) + b) // 31，将 int32 类型转换为 int64 类型，然后相加
}
