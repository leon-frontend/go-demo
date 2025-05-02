package main

import (
	"fmt"
	"reflect"
)

// 实现任意类型和 int 类型相加
func add(x any) {
	// 问题：由于 x 是任意类型，因此将任意类型与 int 类型相加时会报错，如何解决？
	// var num = x + 10 // invalid operation: x + 10 (mismatched types any and int)

	// 解决 1：使用类型断言
	// var num = x.(int) + 10 // 类型断言，可以正常输出

	// 解决 2：先使用 valueOf 获取反射类型的值，再进行类型转换，最后相加
	var xValue = reflect.ValueOf(x)                            // 获取反射类型的值
	fmt.Printf("xValue: %v, xValueType: %T\n", xValue, xValue) // xValue: 10, xValueType: reflect.Value
	var xInt = xValue.Int()                                    // 类型转换
	fmt.Printf("xInt: %v, xIntType: %T\n", xInt, xInt)         // xInt: 10, xIntType: int64
	var num = xInt + 10
	fmt.Printf("num value: %v\n", num) // 20
}

// 将任意类型转换为变量的原始类型
func reflectValueFn(x any) {
	// 获取反射类型的值
	var xValue = reflect.ValueOf(x)

	// 获取原始类型的种类
	var xKind = xValue.Kind()

	// 根据不同的原始类型输出不同的数据
	switch xKind {
	case reflect.Int:
		fmt.Printf("x 是 Int 类型，其原始值为：%d\n", xValue.Int())
	case reflect.Float32:
		fmt.Printf("x 是 Float32 类型，其原始值为：%f\n", xValue.Float())
	case reflect.String:
		fmt.Printf("x 是 String 类，其原始值为：%s\n", xValue.String())
	default:
		fmt.Printf("x 是 %T 类型，其原始值为：%v\n", xValue.Interface(), xValue.Interface())
	}
}

func main() {
	// 1. 测试任意类型和 int 类型相加
	add(10)

	// 2. 类型转换
	var (
		a int     = 10
		b float32 = 20.5
		c string  = "hello"
	)
	reflectValueFn(a)
	reflectValueFn(b)
	reflectValueFn(c)
	reflectValueFn(true)
	fmt.Println("---------------------------------------")

	// 3. 让反射对象具备可写性
	var x = "Hello"                                    // 声明一个字符串
	xValue1 := reflect.ValueOf(&x)                     // 1. 使用指针类型的 x 获取反射对象
	fmt.Printf("xValue1 的可写性为：%v\n", xValue1.CanSet()) // false
	fmt.Printf("xValue1 的底层类型为 %v\n", xValue1.Kind())  // ptr
	xValue2 := xValue1.Elem()                          // 2. 使用 Elem 函数指向获取 x 的值
	fmt.Printf("xValue2 的可写性为：%v\n", xValue2.CanSet()) // true
	fmt.Printf("xValue2 的底层类型为 %v\n", xValue2.Kind())  // string
	xValue2.SetString("World")                         // 修改 xValue2 的值
	fmt.Printf("xValue2 修改后的值为：%v\n", x)               // World
}
