package main

import "fmt"

func main() {
	// 1. ------------------- if 语句 -------------------
	// age1 := 20 // age1 是当前作用于的全局变量
	// if age1 > 18 {
	// 	fmt.Println("已经成年了")
	// }
	// fmt.Println("年龄是：", age1) // 20

	// 等价于
	// if age2 := 20; age2 > 18 {
	// 	fmt.Println("已经成年了")
	// }
	// fmt.Println("年龄是：", age2) // undefined，age2 是 if 语句中的局部变量，作用域只在 if 语句中

	// 2. -------------- 使用 for range 遍历包含中文的字符串 ---------------
	// 中文字符在 UTF-8 编码中占用 3 个字节，所以使用 for range 遍历字符串时，i 的值是字符的起始索引。
	// for i, c := range "你好hello, 世界" {
	// 	fmt.Printf("%d(%c) ", i, c) // 0(你) 3(好) 6(h) 7(e) 8(l) 9(l) 10(o) 11(,) 12( ) 13(世) 15(界)
	// }

	// 3. ------------------ switch case 语句 ------------------
	// 3.1.1 练习：判断文件类型,如果后缀名是.html 输入 text/html, 如果后缀名.css 输出text/css ,如果后缀名是.js 输出 text/javascript
	// extname := ".html"
	// switch extname {
	// case ".html":
	// 	fmt.Println("text/html")
	// case ".css":
	// 	fmt.Println("text/css")
	// case ".js":
	// 	fmt.Println("text/javascript")
	// default:
	// 	fmt.Println("unknown")
	// }

	// // 3.1.2 另一种写法
	// switch extname := ".css"; extname {
	// case ".html":
	// 	fmt.Println("text/html")
	// case ".css":
	// 	fmt.Println("text/css")
	// case ".js":
	// 	fmt.Println("text/javascript")
	// default:
	// 	fmt.Println("unknown")
	// }

	// 3.2 一个分支可以有多个值，用逗号分隔
	switch n := 5; n {
	case 1, 2, 3:
		fmt.Println("1, 2, 3")
	case 4, 5:
		fmt.Println("4, 5")
	default:
		fmt.Println("unknown")
	}

	// 3.3 case 分支可以是表达式
	age := 31
	// 如果 case 分支是表达式，则 switch 后面不用跟变量
	switch {
	case age < 18:
		fmt.Println("未成年")
	case age >= 18 && age < 30:
		fmt.Println("青年")
	case age >= 30 && age < 50:
		fmt.Println("中年")
	default:
		fmt.Println("老年")
	}

	// 3.4 fallthrough 语句
	// fallthrough 语句用于在 case 分支中强制执行下一个 case 分支的语句
	switch n := 5; n {
	case 1, 2, 3:
		fmt.Println("1, 2, 3")
	case 4, 5:
		fmt.Println("4, 5")
		fallthrough // 强制执行下一个 case 分支的语句
	case 6:
		fmt.Println("6")
	default:
		fmt.Println("unknown")
	}

	// 4. ------------------ break 多重循环 ------------------
	// 使用 break 和 label 跳出多重循环
	// 输出：i=0, j=0 \n i=0, j=1
label:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 2 {
				break label // 跳出 label 标签的循环,即跳出双重循环
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}

	// 5. ------------------ goto 无条件跳转 ------------------
	// goto 语句用于无条件跳转到指定标签的位置
	var n int = 30

	if n > 24 {
		fmt.Println("成年人")
		goto label1
	}
	fmt.Println("aaa")
	fmt.Println("bbb")
label1:
	fmt.Println("ccc")
	fmt.Println("ddd")
	// 输出：成年人 \n ccc \n ddd
}
