package main

import "fmt"

func main() {
	// 1. ---------------- 声明切片 ----------------
	var slice1 []int
	fmt.Printf("value: %v, type: %T, length: %d\n", slice1, slice1, len(slice1)) // value: [], type: []int, length: 0
	fmt.Println(slice1 == nil)                                                   // true

	var slice2 = []int{1, 2, 3, 4, 5}
	fmt.Printf("value: %v, type: %T, length: %d\n", slice2, slice2, len(slice2)) // value: [1 2 3 4 5], type: []int, length: 5

	var slice3 = []int{1: 2, 3: 6, 2: 10}
	fmt.Printf("value: %v, type: %T, length: %d\n", slice3, slice3, len(slice3)) // value: [0 2 10 6], type: []int, length: 4

	// 2. ---------------- 切片的循环遍历 ----------------
	strSlice := []string{"php", "go", "python", "java", "c++"}
	for index, value := range strSlice {
		fmt.Printf("index: %v, value: %v\n", index, value)
	}

	// 3. ---------------- 基于数组创建切片 ----------------
	arr := [5]int{1, 2, 3, 4, 5}
	slice4 := arr[:]                                                   // arr[:] 表示获取整个数组 arr 的切片
	fmt.Printf("slice4 value: %v, silice4 type: %T\n", slice4, slice4) // slice4 value: [1 2 3 4 5], silice4 type: []int
	slice5 := arr[1:4]                                                 // arr[1:4] 表示获取数组 arr 的下标 1 到下标 3 元素的切片
	fmt.Printf("slice5 value: %v, silice5 type: %T\n", slice5, slice5) // slice5 value: [2 3 4], silice5 type: []int
	slice6 := arr[:3]                                                  // arr[:3] 表示获取数组 arr 的前 3 个元素的切片
	fmt.Printf("slice6 value: %v, silice6 type: %T\n", slice6, slice6) // slice6 value: [1 2 3], silice6 type: []int
	slice7 := arr[1:]                                                  // arr[1:] 表示获取数组 arr 的下标 1 到最后一个元素的切片
	fmt.Printf("slice7 value: %v, silice7 type: %T\n", slice7, slice7) // slice7 value: [2 3 4 5], silice7 type: []int

	// 4. ---------------- 基于切片创建切片 ----------------
	slice8 := []int{1, 2, 3, 4, 5}
	slice9 := slice8[1:]                                               // slice8[1:] 表示获取切片 slice8 的下标 1 到最后一个元素的切片
	fmt.Printf("slice9 value: %v, silice9 type: %T\n", slice9, slice9) // slice9 value: [2 3 4 5], silice9 type: []int

	// 5. ---------------- 切片的长度和容量 ----------------
	slice10 := []int{1, 2, 3, 4, 5, 23, 55, 699}
	fmt.Printf("slice10 length: %d, capacity: %d\n", len(slice10), cap(slice10)) // slice10 length: 8, capacity: 8

	slice11 := slice10[2:5]                                                      // slice10[2:5] 表示获取切片 slice10 的下标 2 到下标 4 元素的切片
	fmt.Printf("slice11 length: %d, capacity: %d\n", len(slice11), cap(slice11)) // slice11 length: (5-2=)3, capacity: (8-2=)6

	// 5.1 基于底层数组创建切片时，指定第三个数
	// 在切片时，若不指定第三个数，则切片的终止索引是底层数组的末尾元素。而如果指定了第三个数，那么切片终止索引是底层数组的该索引值减 1。
	slice12 := slice10[2:5:7]                                                    // slice10[2:5:6] 表示获取切片 slice10 的下标 2 到下标 4 元素的切片，容量为 6
	fmt.Printf("slice12 length: %d, capacity: %d\n", len(slice12), cap(slice12)) // slice12 length: (5-2=)3, capacity: (7-2=)5

	// 6. --------------- 使用 make() 函数创建切片 ------------------
	slice13 := make([]int, 5, 10)                                                                    // make([]int, 5, 10) 表示创建一个长度为 5，容量为 10 的切片
	fmt.Printf("slice13 value: %v, length: %d, capacity: %d\n", slice13, len(slice13), cap(slice13)) // slice13 value: [0 0 0 0 0], length: 5, capacity: 10

	// 6.1 修改切片中的元素
	slice13[0] = 100
	slice13[1] = 200
	fmt.Printf("new slice13 value: %v, length: %d, capacity: %d\n", slice13, len(slice13), cap(slice13)) // new slice13 value: [100 200 0 0 0], length: 5, capacity: 10

	// 7. ---------------- 切片的扩容方式 ----------------
	slice14 := []int{1}
	// slice14[3] = 100 // 错误的扩容方法，切片无法通过下标的方法进行扩容
	fmt.Printf("slice14 value: %v, length: %d, capacity: %d\n", slice14, len(slice14), cap(slice14)) // slice14 value: [1], length: 1, capacity: 1

	// 7.1 切片扩容时，将新切片的容量扩展到大于原始容量的两倍
	slice14 = append(slice14, 100, 222)                                                                  // 使用 append() 函数进行扩容
	fmt.Printf("new slice14 value: %v, length: %d, capacity: %d\n", slice14, len(slice14), cap(slice14)) // new slice14 value: [1 100 222], length: 3, capacity: 3

	// 7.2 切片扩容时，原始容量 < 扩容后的新切片容量 < 2 * 原始容量
	slice15 := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice15 value: %v, length: %d, capacity: %d\n", slice15, len(slice15), cap(slice15))     // slice15 value: [1 2 3 4 5], length: 5, capacity: 5
	slice15 = append(slice15, 100, 222)                                                                  // 使用 append() 函数进行扩容
	fmt.Printf("new slice15 value: %v, length: %d, capacity: %d\n", slice15, len(slice15), cap(slice15)) // new slice15 value: [1 2 3 4 5 100 222], length: 7, capacity: 10

	// 8. ---------------- 切片的合并 ----------------
	slice16 := []int{1, 2, 3}
	slice17 := []int{4, 5}
	slice18 := append(slice16, slice17...)                                                           // 使用 ... 将切片 slice17 展开，合并到切片 slice16 中
	fmt.Printf("slice18 value: %v, length: %d, capacity: %d\n", slice18, len(slice18), cap(slice18)) // slice18 value: [1 2 3 4 5], length: 5, capacity: 6

	// 9. ---------------- 切片的复制 ----------------
	// 9.1 切片是引用数据类型，改变副本的值会影响原切片的值
	slice19 := []int{1, 2, 3}
	slice20 := slice19                                           // slice20 和 slice19 指向同一个底层数组（引用地址）
	slice20[0] = 100                                             // 修改 slice20 中的值会影响 slice19 中的值
	fmt.Printf("slice19 = %v, slice20 = %v\n", slice19, slice20) // slice19 = [100 2 3], slice20 = [100 2 3]

	// 9.2 使用 copy() 函数复制切片
	slice21 := []int{1, 2, 3}
	slice22 := make([]int, 3, 3)                                 // 创建一个长度为 3 ，容量为 3 的切片
	copy(slice22, slice21)                                       // 使用 copy() 函数复制切片
	slice22[0] = 100                                             // 修改 slice22 中的值不会影响 slice21 中的值
	fmt.Printf("slice21 = %v, slice22 = %v\n", slice21, slice22) // slice21 = [1 2 3], slice22 = [100 2 3]

	// 10. ---------------- 删除切片中的某个元素 ----------------
	slice23 := []int{1, 2, 3, 4, 5}
	// 删除下标为 2 的元素
	slice23 = append(slice23[:2], slice23[3:]...)
	fmt.Printf("slice23 = %v\n", slice23) // slice23 = [1 2 4 5]

}
