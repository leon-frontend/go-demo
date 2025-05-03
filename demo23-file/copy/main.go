package main

import "os"

func main() {
	// 1. 采用 ioutil 工具包实现复制文件
	byteSlice, err := os.ReadFile("./script.txt") // 1.1. 读取文件

	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./copy1.txt", byteSlice, 0666) // 1.2 写入文件
}
