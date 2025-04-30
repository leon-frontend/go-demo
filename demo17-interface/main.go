package main

import "fmt"

// 1. 定义接口
type Usber interface {
	Start()
	Stop()
}

// 2. 定义结构体
type Phone struct {
	Name string
}

// 2.1 Phone 结构体实现 Usber 接口中的所有方法
func (p Phone) Start() {
	fmt.Printf("%v 开始工作 start ...\n", p.Name)
}

func (p Phone) Stop() {
	fmt.Printf("%v 停止工作 stop ...\n", p.Name)
}

// 3. 定义 computer 结构体
type Computer struct {
}

// 3.1 给 Computer 结构体绑定方法，Work 函数的形参必须实现了 Usber 接口
func (c Computer) Work(u Usber) {
	u.Start()
	u.Stop()
}

// ----------------------------------------------------------------------------------------------------------------------
// ----------------------------------------------------------------------------------------------------------------------
func main() {
	// 1. 在 Phone 结构体的实例上实现 Usber 接口
	// 1.1 实例化 Phone 结构体
	var p1 = Phone{
		Name: "1+13t",
	}

	// 1.2 声明变量 u 为 Usber 接口类型
	var u Usber = p1 // 实现接口，实例 p 必须实现 Usber 接口中的所有方法
	u.Start()        // Phone 开始工作 start ...
	u.Stop()         // Phone 停止工作 stop ...

	// 2. 调用 Computer 结构体中的 Work 方法
	var c = Computer{}
	var p2 = Phone{
		Name: "huawei",
	}
	c.Work(p2) // 相当于 u = p2，若 p2 实例没有实现 Usber 接口中的所有方法，则编译报错
	fmt.Println("--------------------------------------------------------")

	// 3. 实现空接口

	var e interface{} // 字符串实现空接口，空接口可以存储任何类型的值
	e = "hello world"
	fmt.Printf("e value = %v, e type = %T\n", e, e) // e value = hello world, e type = string

	// 4. 类型断言
	var f interface{}       // 声明一个空接口类型的变量
	f = "s"                 // 字符串实现空接口
	value, ok := f.(string) // 类型断言，返回值为两个值，第一个值是断言的结果，第二个值是断言是否成功
	if ok {
		fmt.Printf("f value = %v, f type = %T\n", value, value) // f value = nihao, f type = string
	} else {
		fmt.Printf("f value = %v, f type = %T\n", value, value) // f value = "" , f type = string
	}
}
