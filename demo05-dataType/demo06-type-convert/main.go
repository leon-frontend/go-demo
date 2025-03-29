package main

import (
	"fmt"
	"strconv"
)

func main() {
	// // 1. int 类型之间的转换
	// var a int32 = 10
	// var b int64 = 20
	// fmt.Printf("a + b = %d", int64(a)+b)

	// // 2. float 类型之间的转换
	// var c float32 = 10.5
	// var d float64 = 20.5
	// fmt.Printf("c + d = %f", float64(c)+d)

	// // 3. int 类型和 float 类型之间的转换
	// var e int = 10
	// var f float32 = 20.5
	// fmt.Printf("e + f = %f", float32(e)+f)

	// 4. ----------------- 其他类型转换为 string 类型 -----------------
	// 4.1 使用 fmt.Sprintf() 函数
	// var g int = 10
	// var h float32 = 20.5576
	// var i bool = true
	// var j byte = 'a'

	// // Sprintf() 函数的参数：第一个值是要转换的值的原有类型，第二个值是要转换的值，返回一个 string 类型的值
	// str1 := fmt.Sprintf("%d", g)   // int 转 string
	// str2 := fmt.Sprintf("%.2f", h) // float 转 string
	// str3 := fmt.Sprintf("%t", i)   // bool 转 string
	// str4 := fmt.Sprintf("%c", j)   // byte 转 string，使用 %c 原样输出字符，否认默认输出 ascii 码

	// fmt.Printf("value = %s, type = %T \n", str1, str1) // value = 10, type = string
	// fmt.Printf("value = %s, type = %T \n", str2, str2) // value = 20.50, type = string
	// fmt.Printf("value = %s, type = %T \n", str3, str3) // value = true, type = string
	// fmt.Printf("value = %s, type = %T \n", str4, str4) // value = a, type = string

	// 4.2 使用 strconv 包中的函数
	// 使用 strconv.FormatInt() 函数将 int 转换为 string
	// strconv.FormatInt() 函数的参数：第一个值是要转换的值(必须是 int64 类型)，第二个值是进制，10 表示十进制
	// str5 := strconv.FormatInt(int64(g), 10)            // g = 10
	// fmt.Printf("value = %s, type = %T \n", str5, str5) // value = 10, type = string

	// 使用 strconv.FormatFloat() 函数将 float 转换为 string
	/*
		strconv.FormatFloat() 函数的参数：
		第一个值是要转换的值(必须是 float64 类型)，如果要转换的值是 float32 类型，必须先转换为 float64 类型
		第二个值是转换后想要显示的浮点数形式，'f' 表示十进制，'e' 表示科学计数法，'g' 表示自动选择格式
		第三个值是保留几位小数（精度），0 表示默认精度，-1 表示自动选择精度
		第四个值是 bit size，32 表示 float32，64 表示 float64
	*/
	// str6 := strconv.FormatFloat(float64(h), 'e', 2, 32) // h = 20.5576
	// fmt.Printf("value = %s, type = %T \n", str6, str6)  // value = 2.06e+01, type = string

	// 使用 strconv.FormatBool() 函数将 bool 转换为 string
	// strconv.FormatBool() 函数的参数：第一个值是要转换的值，返回一个 string 类型的值
	// str7 := strconv.FormatBool(i)                      // i = true
	// fmt.Printf("value = %s, type = %T \n", str7, str7) // value = true, type = string

	// 使用 strconv.FormatUint() 函数将 byte 转换为 string
	// strconv.FormatUint() 函数的参数：第一个值是要转换的值(必须是 uint64 类型)，第二个值是进制，10 表示十进制
	// str8 := strconv.FormatUint(uint64(j), 10)          // j = 'a'
	// fmt.Printf("value = %s, type = %T \n", str8, str8) // value = 97, type = string

	// 5. ----------------- string 类型转换为其他类型 -----------------
	// 5.1 string 类型转换为 int 类型：使用 strconv.ParseInt() 函数
	// str1 := "130"
	// str2 := "xxx"
	// // strconv.ParseInt() 函数的参数：第一个值是要转换的值；第二个值是转换后的值的进制，10 表示十进制；第三个值是 bitSize 用于指定解析后的整数的范围
	// // 返回值：第一个值是转换后的值，若转换失败则返回 0，并返回错误信息；第二个值是错误信息，如果没有错误，第二个值为 nil
	// // 注意：strconv.ParseInt 函数的返回值始终是 int64 类型，无论你设置的 bitSize 是 8、16 还是 32。第三个参数 bitSize 的作用是 限制解析后的数值范围，而不会改变返回值的类型。
	// num1, err := strconv.ParseInt(str1, 10, 8)                       // str1 = "130"
	// fmt.Printf("value = %d, type = %T, err: %v \n", num1, num1, err) // value = 127, type = int64, err: strconv.ParseInt: parsing "130": value out of range
	// num2, err := strconv.ParseInt(str2, 10, 32)                      // str2 = "xxx"
	// fmt.Printf("value = %d, type = %T, err: %v \n", num2, num2, err) // value = 0, type = int64, err: strconv.ParseInt: parsing "xxx": invalid syntax

	// 5.2 string 类型转换为 float 类型：使用 strconv.ParseFloat() 函数
	str3 := "20.5576"
	// str4 := "20.5576a"
	// strconv.ParseFloat() 函数的参数：第一个值是要转换的值；第二个值是 bitSize，用于指定解析后的浮点数的精度范围，32 表示 float32，64 表示 float64
	// 返回值：第一个值是转换后的值，若转换失败则返回 0，并返回错误信息；第二个值是错误信息，如果没有错误，第二个值为 nil
	num3, err := strconv.ParseFloat(str3, 32)                        // str3 = "20.5576"
	fmt.Printf("value = %f, type = %T, err: %v \n", num3, num3, err) // value = 20.557600, type = float64, err: <nil>
}
