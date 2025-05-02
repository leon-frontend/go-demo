package main

import "fmt"

// 定义结构体
type Student struct {
	Name  string `json:"name"` // 结构体标签
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

// 给 Student 结构体绑定办法
func (s Student) GetInfo() string {
	return fmt.Sprintf("name: %v, age: %v, score: %v", s.Name, s.Age, s.Score)
}

func (s *Student) SetInfo(name string, age int, score int) {
	s.Name = name
	s.Age = age
	s.Score = score
}

func (s Student) Print() {
	fmt.Printf("This is Print Method ...")
}

// 获取结构体中的字段
func GetStructField() {
	// 1. 通过反射对象中的 Field() 方法获取结构体中的字段

	// 2. 通过反射对象中的 FieldByName() 方法获取结构体中的字段

	// 3. 通过反射对象中的 NumField() 方法获取结构体中的字段数量
}

func main() {
	// 实例化 Student 结构体
	s1 := Student{
		Name:  "张三",
		Age:   18,
		Score: 100,
	}

}
