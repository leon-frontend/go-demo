package main

func main() {
	// // 1. 定义字符，字符默认属于 int 类型
	// var a byte = 'a'
	// fmt.Printf("a value: %v, type: %T\n", a, a) // a value: 97, type: int32，%v 默认输出 ASCII 码值
	// fmt.Printf("a value: %c, type: %T\n", a, a) // a value: a, type: int32，%c 输出字符
	// var china = '中'
	// fmt.Printf("china value: %v, type: %T\n", china, china) // china value: 20013, type: int32，中文字符对应的 Unicode 码值

	// // 2. 定义一个字符串，输出字符串中的某个字符
	// var str = "hello"
	// // default output e: 101, char output e: e, type: uint8
	// fmt.Printf("default output e: %v, char output e: %c, type: %T\n", str[1], str[1], str[1])

	// // 3. 一个汉字占用 3 个字节（utf-8 编码），一个英文字符占用 1 个字节
	// var str2 = "this"
	// var str3 = "中文go"
	// fmt.Printf("str2 占用空间大小: %d\n", len(str2)) // str2 占用空间大小: 4
	// fmt.Printf("str3 占用空间大小: %d\n", len(str3)) // str3 占用空间大小: 8

	// 4. 遍历字符串
	// str := "hello你好"

	// // for 循环相当于先将字符串转换为 byte 类型，然后以 byte 类型为单位遍历字符串。若字符串中含有中文字符，会出现乱码。
	// for i := 0; i < len(str); i++ {
	// 	fmt.Printf("%v(%c) ", str[i], str[i]) // 104(h) 101(e) 108(l) 108(l) 111(o) 228(ä) 189(½) 160( ) 229(å) 165(¥) 189(½)
	// }

	// 使用 for range 遍历包含中文的字符串，可以避免乱码问题。
	// for _, value := range str {
	// 	fmt.Printf("%v(%c) ", value, value) // 104(h) 101(e) 108(l) 108(l) 111(o) 20320(你) 22909(好)
	// }

	// 5. 修改字符串中的字符
	str := "你好，hello"
	// str[0] = 'H' // 编译错误：cannot assign to str[0]，字符串是不可变的
	// 正确做法：先将字符串转换为 byte 数组，然后修改数组中的元素，最后再将 byte 数组转换为字符串。
	byteStr := []rune(str)
	byteStr[0] = 'H'      // 修改 byte 数组中的元素
	str = string(byteStr) // 将 byte 数组转换为字符串
	println(str)          // 输出：H好，hello
}
