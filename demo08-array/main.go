package main

import "fmt"

// 练习 1：请求出一个数组里面元素的和以及这些元素的平均值。分别用for和for-range实现
func addAndAverage() {
	arr := [...]int{12, 34, 56, 78, 90}
	sum := 0 // 求和

	for _, value := range arr {
		sum += value
	}

	fmt.Printf("数组的和为：%d, 平均值为：%.2f\n", sum, float64(sum)/float64(len(arr)))
}

// 练习 2：请求出一个数组的最大值，并得到对应的下标
func maxValue() {
	arr := [...]int{12, 34, 56, 100, 78, 90}
	max := arr[0] // 假设第一个元素是最大值
	index := 0    // 记录最大值的下标

	for i, value := range arr {
		if value > max {
			max = value
			index = i
		}
	}

	fmt.Printf("数组的最大值为：%d，下标为：%d\n", max, index)
}

func main() {
	// 调用练习中的方法
	addAndAverage()
	maxValue()

	// 1. 声明和初始化数组
	var arr1 [3]int                                      // 声明一个长度为 3 的整型数组，默认值为 [0 0 0]
	var arr2 = [3]int{1, 2, 3}                           // 声明一个长度为 3 的整型数组，初始化为 [1 2 3]
	var arr3 = [...]int{1, 2, 3, 4, 5}                   // 声明一个长度为 5 的整型数组，初始化为 [1 2 3 4 5]
	var arr4 = [3]int{1: 2, 0: 1}                        // 声明一个长度为 3 的整型数组，初始化为 [1 2 0]
	var arr5 = [3]int{1, 2}                              // 声明一个长度为 3 的整型数组，初始化为 [1 2 0]
	fmt.Printf("arr1 value: %v, type: %T\n", arr1, arr1) // arr1 value: [0 0 0], type: [3]int
	fmt.Printf("arr2 value: %v, type: %T\n", arr2, arr2) // arr2 value: [1 2 3], type: [3]int
	fmt.Printf("arr3 value: %v, type: %T\n", arr3, arr3) // arr3 value: [1 2 3 4 5], type: [5]int
	fmt.Printf("arr4 value: %v, type: %T\n", arr4, arr4) // arr4 value: [1 2 0], type: [3]int
	fmt.Printf("arr5 value: %v, type: %T\n", arr5, arr5) // arr5 value: [1 2 0], type: [3]int

	// 2. 使用 for range 遍历数组
	arr6 := [...]string{"php", "java", "python", "golang"}
	for index, value := range arr6 {
		fmt.Printf("index: %v, value: %v\n", index, value)
	}
	// 输出：index: 0, value: php \n index: 1, value: java \n index: 2, value: python \n index: 3, value: golang \n

	// 3. 数组是值类型
	arr7 := [...]int{1, 2, 3}
	arr8 := arr7
	arr8[0] = 100                                  // 修改 arr8 中的某个元素
	fmt.Printf("arr7: %v, arr8: %v\n", arr7, arr8) // arr7: [1 2 3], arr8: [100 2 3]

	// 4. 声明二维数组
	arr9 := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Printf("arr9[0] = %d, arr9[0][0] = %d\n", arr9[0], arr9[0][0]) // arr9[0] = [1 2], arr9[0][0] = 1

	// 5. 遍历二维数组
	for _, rowVal := range arr9 {
		for _, colVal := range rowVal {
			fmt.Printf("%d ", colVal) // 1 2 3 4 5 6
		}
	}

	// 6. 使用 ... 推导二维数组的第一层的长度
	arr10 := [...][2]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Printf("\n%v", arr10)
}
