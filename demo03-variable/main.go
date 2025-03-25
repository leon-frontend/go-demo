package main

import "fmt"

// getUserInfo 函数返回两个值，用于测试匿名变量
func getUserInfo() (string, int) {
	return "zhangsan", 18
}

func main() {
	// var username string // 只声明变量，未赋值
	// println(username)   // 空字符串

	// // 一次性声明多个变量，统一定义变量类型
	// var name, age int // 统一定义 name 和 age 的类型为 int
	// // name = "hello" // 编译报错，因为 name 的类型是 int
	// age = 18
	// println(name, age)

	// 一次性声明多个变量，分开定义变量类型
	// var (
	// 	username string
	// 	age      int
	// 	isMale   bool
	// )

	// username = "zhangsan"
	// age = 18
	// isMale = true
	// fmt.Println(username, age, isMale) // zhangsan 18 true

	// // 等价于
	// var (
	// 	username string = "lisi"
	// 	age      int    = 20
	// )
	// fmt.Println(username, age) // lisi 20

	// // 短类型声明
	// username := "wangwu"
	// fmt.Printf("username的类型为%T", username) // wangwu

	// getUserInfo 函数返回两个值，用于测试匿名变量
	name, _ := getUserInfo()
	fmt.Println(name) // zhangsan
}
