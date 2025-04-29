package main

import (
	"encoding/json"
	"fmt"
)

// 注意：私有属性无法被 json 包访问
type Student struct {
	ID     int
	Gender string
	Name   string
	sno    string
}

// 2. 结构体标签
type StudentTag struct {
	ID     int    `json:"id"`
	Gender string `json:"gender"`
	Name   string // 默认情况下，结构体属性名作为 JSON 字符串的 key
	sno    string `json:"-"` // "-" 表示该属性不参与序列化
}

func main() {
	// 1.1 实例化 Student 结构体
	s1 := Student{
		ID:     1,
		Gender: "男",
		Name:   "张三",
		sno:    "001",
	}
	fmt.Printf("s1 value: %#v, s1 type: %T\n", s1, s1) // s1 value: main.Student{ID:1, Gender:"男", Name:"张三", sno:"001"}, s1 type: main.Student

	// 1.2 使用 json.marshal 将结构体转换为 JSON 字符串
	jsonByteSlice, _ := json.Marshal(s1)                            // 返回 []byte 类型
	jsonStr := string(jsonByteSlice)                                // 将 []byte 转换为字符串
	fmt.Printf("jsonStr: %s, jsonStr type: %T\n", jsonStr, jsonStr) // jsonStr: {"ID":1,"Gender":"男","Name":"张三"}, jsonStr type: string

	// 1.3 使用 json.unmarshal 将 JSON 字符串转换为结构体
	str := `{"ID":1,"Gender":"男","Name":"张三"}` // 定义 json 字符串
	var s2 Student                             // 实例化  Student 结构体
	/*
		Unmarshal 方法的参数解释：
			1. 第一个参数表示将  json 字符串转换为 []byte 类型
			2. 第二个参数表示修改 s2 结构体（由于结构体是值类型，所以修改其值时要传递结构体的地址）
		Unmarshal 方法的返回值：返回一个错误
	*/
	err := json.Unmarshal([]byte(str), &s2)
	if err != nil {
		fmt.Printf("json.Unmarshal failed, err: %v\n", err)
		return
	}
	fmt.Printf("s2 value: %#v, s2 type: %T\n", s2, s2) // s2 value: main.Student{ID:1, Gender:"男", Name:"张三", sno:""}, s2 type: main.Student

	// 2. 结构体标签
	s3 := StudentTag{
		ID:     1,
		Gender: "男",
		Name:   "张三",
		sno:    "001",
	}
	fmt.Printf("s2 value: %#v, s2 type: %T\n", s3, s3)
	jsonByteSlice2, _ := json.Marshal(s3) // 返回 []byte 类型
	jsonStr2 := string(jsonByteSlice2)
	fmt.Printf("jsonStr2: %s, jsonStr2 type: %T\n", jsonStr2, jsonStr2) // jsonStr2: {"id":1,"gender":"男","Name":"张三"}, jsonStr2 type: string
}
