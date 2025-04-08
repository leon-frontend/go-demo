package main

import (
	"errors"
	"fmt"
)

// sumFn 函数用于计算两个整数的和
func sumFn1(x int, y int) int {
	return x + y
}

// 1. 形参类型简写
func sumFn2(x, y int) int {
	return x + y
}

// 2. 函数的可变参数，sumFn3 函数的第一个参数是 x，第二个参数是可变参数 y
func sumFn3(x int, y ...int) int {
	fmt.Printf("x value: %v, type: %T; y value: %v, type: %T\n", x, x, y, y)
	sum := x
	// 使用 for-range 遍历切片 y
	for _, v := range y {
		sum += v
	}
	return sum
}

// 3. 多返回值函数，sumFn4 函数返回两个整数的和和差
func sumFn4(x, y int) (int, int) {
	return x + y, x - y
}

// 4. 返回值命名函数，sumFn5 函数返回两个整数的和和差，返回值命名为 sum 和 minus
func sumFn5(x, y int) (sum, minus int) {
	sum = x + y
	minus = x - y
	return // 直接返回命名返回值
}

// 5. 自定义函数类型
type calc func(int, int) int

// 6. 自定义变量类型
type myInt int

// 7. 高阶函数
// 自定义函数类型
type calcType func(int, int) int

// 7.1 函数作为参数时的高阶函数
func calcFn(x, y int, cb calcType) int {
	return cb(x, y)
}

// 7.2 函数作为返回值时的高阶函数
type calcType2 func(int, int) int

func do(o string) calcType2 {
	switch o {
	case "+":
		return sumFn1
	case "-":
		return func(x, y int) int {
			return x - y
		}
	case "*":
		return func(x, y int) int {
			return x * y
		}
	default:
		return nil
	}
}

// 8. 闭包
func add1() func() int {
	var i = 10
	return func() int {
		return i + 1 // 并没有改变 i 的值
	}
}

func add2() func() int {
	var i = 10
	return func() int {
		i = i + 2 // 改变了 i 的值，闭包会保存每次执行后 i 的值
		return i
	}
}

// 9. defer 延迟语句
// 9.1 defer 在匿名返回值函数中的表现
func anonReturnFn() int {
	var a int // 0

	defer func() {
		a++
	}()

	return a
}

// 9.2 defer 在命名返回值函数中的表现
func namedReturnFn() (a int) {
	defer func() {
		a++
	}()
	return a
}

// 练习 1
func f4() (x int) { // 1. 进入函数f4，命名返回值x初始化为 0
	defer func(x int) { // 3. 将f4的x的值0复制(值传递)给 defer 函数的形参x(新的局部变量)
		x++ // 5. 函数即将返回，执行defer函数。x++操作修改的是匿名函数自己的形参x(新的局部变量)，不影响f4的返回值x
	}(x) // 2. 遇到defer语句，计算defer函数的实参 x，此时f4的x的值为0
	return 5 // 4. 执行return 5语句，将f4函数的命名返回值x设置为5
} // 6. 函数f4返回最终结果为5

// 练习 2
func calc3(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

// 10. panic 和 recover
// 规则：panic 函数可以在任何地方触发，但 recover 函数只能在 defer 调用的函数中触发。
func fn1() {
	fmt.Println("panic before")
}

func panicFn() {
	// 使用 recover 函数捕获 panic 异常
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover from panic:", err)
		}
	}()
	panic("抛出一个异常") // 触发 panic
}

// 11. 使用 defer、panic 和 recover 抛出异常并处理
func readFile(fileName string) error {
	if fileName == "main.go" {
		return nil // 文件存在，返回 nil，表示没有错误
	} else {
		return errors.New("file not found") // 文件不存在，返回错误
	}
}

// 11.1 模拟读取文件的场景
func myFn() {
	// 使用 defer 语句捕获异常
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("文件不存在，请检查文件名！！！！！")
		}
	}()

	// 读取文件，并触发panic
	err := readFile("xxx.go")
	if err != nil {
		panic(err) // 触发 panic
	}
}

// ---------------------------------------------------------------------------------------------------------------------------------
func main() {
	// 调用 sumFn 函数并打印结果
	result1 := sumFn1(3, 4)
	println("The sum1 is:", result1)

	result2 := sumFn2(2, 11)
	println("The sum2 is:", result2)

	// 2. 函数的可变参数
	result3 := sumFn3(1, 2, 3, 4, 5) // x value: 1, type: int; y value: [2 3 4 5], type: []int
	println("The sum3 is:", result3) // The sum3 is: 15

	// 3. 多返回值函数
	result4, result5 := sumFn4(10, 5)
	fmt.Printf("The sum is: %d, the minus is: %d\n", result4, result5) // The sum is: 15, the minus is: 5

	// 4. 返回值命名函数
	result6, result7 := sumFn5(20, 10)
	fmt.Printf("The sum is: %d, the minus is: %d\n", result6, result7) // The sum is: 30, the minus is: 10

	// 5. 定义函数类型
	var f1 calc                    // 声明一个类型为 calc 的变量 f
	f1 = sumFn1                    // sumFn1 函数必须满足 calc 函数类型
	fmt.Printf("f type: %T\n", f1) // f type: main.calc

	f2 := sumFn2                    // 使用类型推导方式声明变量 f2
	fmt.Printf("f2 type: %T\n", f2) // f2 type: func(int, int) int

	// 6. 自定义变量类型
	var a int = 10                               // 声明一个 int 类型的变量 a
	var b myInt = 20                             // 声明一个 myInt 类型的变量 b
	fmt.Printf("a type: %T, b type: %T\n", a, b) // a type: int, b type: main.myInt
	// 不能直接相加，因为 a 是 int 类型，b 是 myInt 类型，类型不匹配
	// fmt.Printf("a + b = %d\n", a+b)              // invalid operation: a + b (mismatched types int and myInt)
	fmt.Printf("a + int(b) = %d\n", a+int(b)) // a + int(b) = 30

	// 7. 高阶函数，传入匿名函数作为参数
	result8 := calcFn(1, 2, func(x, y int) int {
		return x + y
	})
	fmt.Printf("The calcFn result is: %d\n", result8) // The calcFn result is: 3

	// 7.2 函数作为返回值时的高阶函数
	result9 := do("*")
	fmt.Printf("The result9 is: %d\n", result9(11, 2)) // The result9 is: 22

	// 8. 匿名函数
	// 8.1 将匿名函数赋值给变量
	anonFn := func(x, y int) int {
		return x + y
	}
	fmt.Printf("The anonFn result is: %d\n", anonFn(1, 2)) // The anonFn result is: 3

	// 8.2 使用匿名函数实现自执行函数
	func() {
		fmt.Println("Hello, World!") // Hello, World!
	}()

	// 9. 闭包
	closure1 := add1()                              // 返回一个函数
	fmt.Printf("The closure1 is: %d\n", closure1()) // The closure1 is: 11
	fmt.Printf("The closure1 is: %d\n", closure1()) // The closure1 is: 11
	fmt.Printf("The closure1 is: %d\n", closure1()) // The closure1 is: 11

	closure2 := add2()                              // 返回一个函数
	fmt.Printf("The closure2 is: %d\n", closure2()) // The closure2 is: 12
	fmt.Printf("The closure2 is: %d\n", closure2()) // The closure2 is: 14
	fmt.Printf("The closure2 is: %d\n", closure2()) // The closure2 is: 16

	// 10. 多个 defer 语句
	defer fmt.Println("defer 1") // defer 语句会在函数返回时执行，先入后出
	defer fmt.Println("defer 2") // defer 语句会在函数返回时执行，先入后出
	defer fmt.Println("defer 3") // defer 语句会在函数返回时执行，先入后出
	defer fmt.Println("defer 4") // defer 语句会在函数返回时执行，先入后出
	// defer 4 \n defer 3 \n defer 2 \n defer 1

	// 10.1 defer 在匿名返回值函数中的表现
	result10 := anonReturnFn()
	fmt.Printf("The anonReturnFn is: %d\n", result10) // The anonReturnFn is: 0

	// 10.2 defer 在命名返回值函数中的表现
	result11 := namedReturnFn()
	fmt.Printf("The namedReturnFn is: %d\n", result11) // The namedReturnFn is: 1

	// 练习 1
	f4Result := f4()
	fmt.Printf("The f4 result is: %d\n", f4Result) // The f4 result is: 5

	// 练习 2
	x := 1
	y := 2
	defer calc3("AA", x, calc3("A", x, y))
	x = 10
	defer calc3("BB", x, calc3("B", x, y))
	y = 20
	fmt.Println(1)
	// A 1 2 3
	// B 10 2 12
	// 1
	// BB 10 12 22
	// AA 1 3 4

	// 10. panic 和 recover
	fn1()                      // panic before
	panicFn()                  // 触发 panic
	fmt.Println("panic after") // 不会执行到这里

	// 模拟读取文件场景
	myFn() // 文件不存在，请检查文件名！！！！！
	fmt.Println("继续执行！！！！")
}
