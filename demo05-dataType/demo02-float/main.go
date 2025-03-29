package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 1. float32 单精度浮点数
	var a float32 = 3.12
	fmt.Printf("a value: %f - a type: %T - sizeof: %d Bytes", a, a, unsafe.Sizeof(a)) // 3.120000 - float32 - 4 Bytes

	// 2. float64 双精度浮点数, 使用 .2f 表示保留两位小数
	var b float64 = 3.12
	fmt.Printf("\nb value: %.2f - b type: %T - sizeof: %d Bytes", b, b, unsafe.Sizeof(b)) // 3.12 - float64 - 8 Bytes

	// 3. 64 位系统中的默认浮点类型是 float64
	c := 3.12
	fmt.Printf("\nc value: %.3f - c type: %T - sizeof: %d Bytes", c, c, unsafe.Sizeof(c)) // 3.120 - float64 - 8 Bytes

	// 4. 使用科学计数法表示浮点数
	d := 3.1415e2
	fmt.Printf("\nd value: %f - d type: %T - sizeof: %d Bytes", d, d, unsafe.Sizeof(d)) // 314.150000 - float64 - 8 Bytes

	// 5. 精度丢失问题
	m1 := 85.223123
	m2 := 32.814237
	fmt.Printf("\nm1 - m2 value: %f", (m1 - m2))
}
