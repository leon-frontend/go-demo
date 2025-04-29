package main

/*
	1. 标准库排最前面，第三方包次之、项目内的其它包和当前包的子包排最后，每种分类以一空行分隔。
	2. 尽量不要使用相对路径来导入包。
*/
import (
	"fmt"

	"demo16-package/calc" // demo16-package 要与 go.mod 中的 module 名字一致
	"demo16-package/tools"
)

func main() {
	// 调用 calc 包中的方法
	result1 := calc.Add(10, 20)
	fmt.Println("result1:", result1)

	// 调用 tools 包中的方法
	result2 := tools.Mul(10, 20)
	fmt.Println("result2:", result2)
}
