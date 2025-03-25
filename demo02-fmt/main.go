package main

import "fmt"

func main() {
	/*
	 * Println：会自动换行，会在输出内容的末尾加上一个换行符
	 * Print：不会自动换行，不会在输出内容的末尾加上一个换行符
	 * Printf：不会自动换行，可以格式化输出，可以使用占位符(比如%d)
	 */
	// Println vs Print：一次输出多个值时，Println 中间会有空格，Print 中间没有空格
	// fmt.Println("A", "B", "C") // A B C\n, 三个字符之间会有空格，但是会在一行显示
	// fmt.Print("A", "B", "C")   // ABC，三个字符之间没有空格
	// fmt.Printf("Printf输出结果: Hello, World!")

	// Printf 格式化输出
	// 定义变量的格式为 var <name> <type>，其中 var 是关键字（固定不变），name 是变量名，type 是类型。
	var a int = 10
	var b int = 3
	var c int = 5
	// 使用 Println 方法实现格式化输出
	fmt.Println("a=", a, "b=", b, "c=", c) // a= 10 b= 3 c= 5，注意字符之间默认存在一个空格
	// 使用 Printf 方法实现格式化输出
	fmt.Printf("a=%d, b=%d, c的类型是 %T\n", a, b, c) // a=10, b=3, c的类型是 int\n，会按照指定的格式输出，并且换行
}
