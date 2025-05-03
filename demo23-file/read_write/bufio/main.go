package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 1. 打开和关闭文件
	file, err := os.Open("../test.txt")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	// 2. 使用 bufio 读取文件
	var fileStr string // 定义一个变量用来存储读取到的全部内容
	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n') // 表示一次读取一行

		if err == io.EOF { // 表示文件读取完毕
			fileStr += str // 将最后一行读取到的数据添加到 fileStr 中
			break
		}

		if err != nil {
			panic(err)
		}

		fileStr += str // 将读取到的数据添加到 fileStr 中
	}

	fmt.Println(fileStr)

	// 3. 写文件
	file1, err := os.OpenFile("../write.txt", os.O_RDWR|os.O_APPEND, 0666)
	defer file1.Close()

	if err != nil {
		panic(err)
	}

	// 3.1 使用 bufio 写文件
	writer := bufio.NewWriter(file1)       // 创建一个 bufio.Writer 对象
	writer.WriteString("\nbuffer writer1") // 将数据写入 buffer 缓冲区中
	writer.Flush()                         // 将 buffer 缓冲区中的数据写入到文件中
}
