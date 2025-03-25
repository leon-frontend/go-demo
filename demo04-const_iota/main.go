package main

import "fmt"

func main() {
	// // 一次性声明多个常量
	// const (
	// 	PI = 3.14
	// 	E  = 2.71
	// 	G  // 常量 G 的值和上一个常量 E 的值相同
	// )
	// fmt.Println(PI, E, G) // 3.14 2.71 2.71

	// iota 的用法
	const a = iota
	fmt.Println(a) // 0

	// 1. 使用 iota 一次性声明多个常量
	const (
		b = iota
		c
		d
	)
	fmt.Println(b, c, d) // 0 1 2

	// 2. 使用 iota 一次性声明多个常量时，可以使用下划线跳过某些值
	const (
		e = iota
		_
		f
	)
	fmt.Println(e, f) // 0 2

	// 3. 使用 iota 一次性声明多个常量时，使用 iota 声明中间插队
	const (
		g = iota
		h = 100
		i = iota
		j
	)
	fmt.Println(g, h, i, j) // 0 100 2 3

	// 4. 多个 iota 定义在一行
	const (
		k, l = iota + 1, iota + 2
		m, n
		o, p = iota + 3, iota + 4
	)
	fmt.Println(k, l, m, n, o, p) // 1 2 2 3 5 6
}
