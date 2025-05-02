package main

import (
	"fmt"
	"reflect"
)

// 自定义类型
type MyInt int

// 定义一个结构体
type Person struct {
	Name string
	Age  int
}

// 使用反射在程序运行时获取任意变量的类型
func reflectTypeFn(x any) {
	xType := reflect.TypeOf(x)
	fmt.Printf("TypeOf(x): %v, t.Name(): %v, t.Kind(): %v\n", xType, xType.Name(), xType.Kind())
}

func main() {
	// 1. 声明几个不同类型的变量，然后使用反射获取它们的值
	var (
		username string = "zhangsan"
		age      int    = 18
		isMale   bool   = true
	)
	reflectTypeFn(username) // TypeOf(x): string, t.Name(): string, t.Kind(): string
	reflectTypeFn(age)      // TypeOf(x): int, t.Name(): int, t.Kind(): int
	reflectTypeFn(isMale)   // TypeOf(x): bool, t.Name(): bool, t.Kind(): bool

	// 2. 声明自定义类型、结构体、指针变量，然后使用反射获取它们的值
	var (
		myInt  MyInt  = 10
		person Person = Person{Name: "lisi", Age: 20}
		p      int    = 20
		arr    [3]int = [3]int{1, 2, 3}
		slice  []int  = []int{1, 2, 3}
	)
	reflectTypeFn(myInt)  // TypeOf(x): main.MyInt, t.Name(): MyInt, t.Kind(): int
	reflectTypeFn(person) // TypeOf(x): main.Person, t.Name(): Person, t.Kind(): struct
	reflectTypeFn(&p)     // TypeOf(x): *int, t.Name(): , t.Kind(): ptr
	reflectTypeFn(arr)    // TypeOf(x): [3]int, t.Name(): , t.Kind(): array
	reflectTypeFn(slice)  // TypeOf(x): []int, t.Name(): , t.Kind(): slice
}
