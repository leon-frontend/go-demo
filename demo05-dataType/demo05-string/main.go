package main

import (
	"fmt"
	"strings"
)

func main() {
	// 	// 1. 字符串的转义字符
	// 	str1 := "hello\nworld"
	// 	fmt.Println(str1)
	// 	str2 := "c:\\code\\go"
	// 	fmt.Println(str2) // c:\code\go

	// 	// 2. 多行字符串
	// 	str3 := `第一行
	// 第二行
	// 第三行`
	// 	fmt.Println(str3)

	// 3. 求字符串的长度
	str4 := "hello沙河"
	fmt.Println(len(str4)) // 11

	// 4. 拼接字符串：+ / fmt.Sprintf
	str5 := "hello"
	str6 := "world"
	fmt.Println(str5 + str6) // helloworld
	str7 := fmt.Sprintf("Sprintf output: %s %s", str5, str6)
	fmt.Println(str7) // Sprintf output: hello world

	// 5. 分割字符串：strings.Split
	str8 := "hello,world,ok"
	arr := strings.Split(str8, ",")
	fmt.Println(arr) // [hello world ok], 返回一个切片

	// 6. 使用 join 将切片中的字符串连接起来
	str9 := strings.Join(arr, "-")
	fmt.Println(str9) // hello-world-ok

	// 7. 判断字符串中是否包含某个字符串：strings.Contains
	fmt.Println(strings.Contains("this world ss", "w")) // true

	// 8. strings.HasPrefix 判断字符串是否以指定字符串开头
	fmt.Println(strings.HasPrefix("this world ss", "hi")) // false

	// 9. strings.HasSuffix 判断字符串是否以指定字符串结尾
	fmt.Println(strings.HasSuffix("this world ssi", "si")) // true

	// 10. strings.Index 查找字符串中是否包含另一个字符串, 返回第一次出现的位置，如果没有找到返回-1
	fmt.Println(strings.Index("this world ssi", "s")) // 3

	// 11. strings.LastIndex 查找字符串中是否包含另一个字符串, 返回最后一次出现的位置，如果没有找到返回-1
	fmt.Println(strings.LastIndex("this world ssi", "s")) // 12
}
