package main

import "fmt"

func main() {
	// 算数运算符
	var a = 3
	var b = 2
	fmt.Println("a + b =", a+b)  // 加法，输出 5
	fmt.Println("a - b = ", a-b) // 减法，输出 1
	fmt.Println("a * b = ", a*b) // 乘法，输出 6
	fmt.Println("a / b = ", a/b) // 除法，两个数都是整数，则输出整数部分 1
	fmt.Println("a % b = ", a%b) // 取余，输出 1

	// 浮点数除法
	var a1 = 10.0
	var b1 = 3.0
	fmt.Println("a1 / b1 = ", a1/b1) // 浮点数除法，输出 3.3333333333333335

	// 取余运算符
	fmt.Println("-10 % 3 = ", -10%3) // 余数 = 被除数 - (被除数 / 除数) * 除数 = -10 - (-10 / 3) * 3 = -1
	fmt.Println("10 % -3 = ", 10%-3) // 余数 = 被除数 - (被除数 / 除数) * 除数 = 10 - (10 / -3) * -3 = 1

	// 自增自减
	// var i int = 8
	// var a int
	// a = i++ // 错误，i++ 只能单独使用
	// a = i-- // 错误，i-- 只能单独使用

	// 正确的自增自减
	i := 8
	i++                   // i = 9
	i--                   // i = 8
	fmt.Println("i =", i) // 输出 8

	// 位运算符
	var x = 3                      // 二进制 0011
	var y = 5                      // 二进制 0101
	fmt.Println("x & y = ", x&y)   // 按位与，输出 1，0011 & 0101 = 0001
	fmt.Println("x | y = ", x|y)   // 按位或，输出 7，0011 | 0101 = 0111
	fmt.Println("x ^ y = ", x^y)   // 按位异或，输出 6，0011 ^ 0101 = 0110
	fmt.Println("y << 2 = ", y<<2) // 左移, 输出 20 (5 * 2^2 = 20)，0101 << 2 = 10100
	fmt.Println("x >> 1 = ", x>>1) // 右移，输出 1 (3 / (2^1) = 1)，0011 >> 1 = 0001
}
