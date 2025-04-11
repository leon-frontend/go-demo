package main

import "fmt"

func main() {
	var a int = 10
	fmt.Printf("a 的值：%v, a 的地址：%p\n", a, &a) // a 的值：10, a 的地址：0xc00000a0f8

	// 1. 定义一个指针变量 b，指向 a 的地址
	var b *int = &a                          // b 是指针变量，指向 a 的地址
	fmt.Printf("b 的值：%v, b 的地址：%p\n", b, &b) // b 的值：0xc00000a0f8(a 的地址), b 的地址：0xc000058058

	// 2. 使用 *pointer 语法根据某个地址（指针变量的值）取值
	var c int = *b                           // c 的值是 b 指向的地址的值，也就是 a 的值
	fmt.Printf("c 的值：%v, c 的地址：%p\n", c, &c) // c 的值：10, c 的地址：0xc00000a0f0

	// 3. 只声明不赋值时，必须手动给指针变量分配内存空间
	// var d *int // 声明一个指针变量，但是不给其分配内存空间
	// *d = 100   // *pointer 表示根据 d 的值（地址）取值，然后再赋值
	// fmt.Printf("d 的值：%v, d 的地址：%p\n", d, &d) // panic: runtime error: invalid memory address or nil pointer dereference

	// 3.1 声明一个指针变量，并使用 new 函数分配内存空间
	var d *int = new(int) // d 是指针变量，指向一个 int 类型的地址
	*d = 100
	fmt.Printf("d 的值：%v, *d 的值: %d, d 的地址：%p\n", d, *d, &d)

	// 4. 使用 new 函数创建的指针类型的默认值
	var e *int = new(int)
	fmt.Printf("e 的值：%v, *e 的默认值: %d, e 的地址：%p\n", e, *e, &e)
	// e 的值：0x14000112048, *e 的默认值: 0, e 的地址：0x1400011e040
	var f *bool = new(bool)
	fmt.Printf("f 的值：%v, *f 的默认值: %t, f 的地址：%p\n", f, *f, &f)
	// f 的值：0x14000112050, *f 的默认值: false, f 的地址：0x1400011e048
}
