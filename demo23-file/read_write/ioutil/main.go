package main

import (
	"os"
)

func main() {
	// 1. 使用 ioutil 读取文件
	// ioutil.ReadFile("../test.txt") // 此方法过时，推荐使用 os.ReadFile 方法代替
	byteSlice, err := os.ReadFile("../test.txt")

	if err != nil {
		panic(err)
	}

	println(string(byteSlice))

	// 2. 使用 ioutil 写文件
	err = os.WriteFile("../write.txt", []byte("\nioutil hahahah"), 0666)

	if err != nil {
		panic(err)
	}
}
