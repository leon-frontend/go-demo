package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 1. 打开和关闭文件
	file, err := os.Open("../test.txt") // 以只读的形式打开文件，返回一个指针类型的变量
	defer file.Close()                  // 操作完文件之后，关闭文件

	if err != nil {
		fmt.Println("打开文件失败：", err)
		return
	}

	// 2. 使用 file.Read() 方法循环读取文件的内容，保证把文件中的内容全部读完
	var resSlice []byte            // resSlice 用来保存读取到的全部内容
	fileSlice := make([]byte, 128) // 创建一个长度为 128 的切片，每次读取文件的大小为 128 字节

	for {
		n, err := file.Read(fileSlice) // 读取文件，每次读取 128 字节。返回的是读取的字节数和错误信息。并将读取的内容放在 fileSlice 切片中。

		// 当 err == io.EOF 时，表示文件已经读取完毕，则退出循环
		if err == io.EOF {
			break
		}

		// 处理文件读取失败的情况
		if err != nil {
			fmt.Println("读取文件失败：", err)
			return
		}

		resSlice = append(resSlice, fileSlice[:n]...) // 将读取到的内容追加到 resSlice 切片中
	}

	fmt.Printf("文件中的全部内容读取完毕，内容是：%s\n", string(resSlice))

	// 3. 写文件，当 perm = 0666 时，表示所有用户都有读写权限
	file1, err := os.OpenFile("../write.txt", os.O_RDWR|os.O_TRUNC, 0666)
	defer file1.Close()

	if err != nil {
		fmt.Println("打开文件失败：", err)
		return
	}

	// 3.1 向 file1 写入数据
	file1.WriteString("hello world1")
	file1.Write([]byte("hello world2"))
}
