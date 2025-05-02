package main

import (
	"fmt"
	"reflect"
)

// 定义结构体
type Student struct {
	Age   int `json:"age_json" form:"age_form"`
	Score int `json:"score_json"`
}

// 给 Student 结构体绑定办法
func (s Student) GetInfo() string {
	return fmt.Sprintf("age: %v, score: %v", s.Age, s.Score)
}

func (s *Student) SetInfo(age int, score int) {
	s.Age = age
	s.Score = score
}

func (s Student) Print() {
	fmt.Println("This is Print Method ...")
}

// ------------------- 判断形参 s 是否是结构体类型 -------------------
func isStruct(s interface{}) bool {
	// 通过反射对象中的 Kind() 方法来判断形参 s 是否是结构体类型
	var t = reflect.TypeOf(s) // 获取形参 s 的反射类型对象
	// 当 s 为指针类型时，需要使用 t.Elem().Kind() 才能获取底层类型
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		return false
	}
	return true
}

// ------------------- 获取结构体中的字段 -------------------
func GetStructField(s interface{}) {
	// 通过底层类型来判断传递进来的形参 s 是否是结构体类型
	if !isStruct(s) {
		fmt.Println("形参 s 不是结构体类型")
		return
	}

	var t = reflect.TypeOf(s) // 获取形参 s 的反射类型对象

	// 1. 通过反射对象中的 Field() 方法根据索引获取结构体中的字段
	field0 := t.Field(0)
	fmt.Printf("Field() 方法获取的字段：%#v\n", field0)     // Field() 方法获取的字段：reflect.StructField{Name:"Age", PkgPath:"", Type:(*reflect.rtype)(0xe3560), Tag:"json:\"age\"", Offset:0x0, Index:[]int{0}, Anonymous:false}
	fmt.Printf("字段名称：%v\n", field0.Name)            // 字段名称：Age
	fmt.Printf("字段类型：%v\n", field0.Type)            // 字段类型：int
	fmt.Printf("字段标签：%v\n", field0.Tag.Get("json")) // 字段标签：age_json
	fmt.Println("-----------------------------------------------------")

	// 2. 通过反射对象中的 FieldByName() 方法获取结构体中的指定字段
	field1, ok := t.FieldByName("Score") // 获取 Age 字段
	if ok {
		fmt.Printf("FieldByName() 方法获取的字段：%#v\n", field1) // FieldByName() 方法获取的字段：reflect.StructField{Name:"Age", PkgPath:"", Type:(*reflect.rtype)(0xe3560), Tag:"json:\"age\"", Offset:0x0, Index:[]int{0}, Anonymous:false}
		fmt.Printf("字段名称：%v\n", field1.Name)              // 字段名称：Score
		fmt.Printf("字段类型：%v\n", field1.Type)              // 字段类型：int
		fmt.Printf("字段标签：%v\n", field1.Tag.Get("json"))   // 字段标签：score_json
	}

	// 3. 通过反射对象中的 NumField() 方法获取结构体中的字段数量
	fmt.Printf("NumField() 方法获取的字段数量：%v\n", t.NumField()) // NumField() 方法获取的字段数量：2
	fmt.Println("-----------------------------------------------------")
}

// ------------------- 获取结构体中的方法 -------------------
func GetStructMethod(s interface{}) {
	// 判断形参 s 是否是结构体类型
	if !isStruct(s) {
		fmt.Println("形参 s 不是结构体类型")
		return
	}

	var t = reflect.TypeOf(s) // 获取反射中的类型对象

	// 1. 通过反射类型对象中的 Method() 方法根据索引获取结构体中的方法
	method0 := t.Method(0)                        // 注意，索引顺序是按照方法名首字母的 ASCII 排序，与方法声明顺序无关
	fmt.Printf("Method() 方法获取的方法：%#v\n", method0) // Method() 方法获取的方法：reflect.Method{Name:"GetInfo", PkgPath:"", Type:(*reflect.rtype)(0xc000108050), Func:reflect.Value{typ_:(*abi.Type)(0xc000108050), ptr:(unsafe.Pointer)(0xc0000580c0), flag:0x13}, Index:0}
	fmt.Printf("方法名称：%v\n", method0.Name)         // 方法名称：GetInfo
	fmt.Printf("方法类型：%v\n", method0.Type)         // 方法类型：func(main.Student) string
	fmt.Println("-----------------------------------------------------")

	// 2. 通过反射类型对象中的 MethodByName() 方法获取结构体中的指定方法
	method1, ok := t.MethodByName("Print")
	if ok {
		fmt.Printf("MethodByName() 方法获取的方法：%#v\n", method1) // MethodByName() 方法获取的方法：reflect.Method{Name:"Print", PkgPath:"", Type:(*reflect.rtype)(0xc000030200), Func:reflect.Value{typ_:(*abi.Type)(0xc000030200), ptr:(unsafe.Pointer)(0xc000058120), flag:0x13}, Index:1}
		fmt.Printf("方法名称：%v\n", method1.Name)               // 方法名称：Print
		fmt.Printf("方法类型：%v\n", method1.Type)               // 方法类型：func(main.Student)
	}
	fmt.Println("-----------------------------------------------------")

	// 3. 通过反射中的值对象实现方法调用
	var v = reflect.ValueOf(s)                       // 获取反射中的值对象
	v.Method(1).Call(nil)                            // 调用 Print 方法，nil 表示不传递任何参数
	infoSlice := v.MethodByName("GetInfo").Call(nil) // 调用 GetInfo 方法，nil 表示不传递任何参数
	fmt.Println(infoSlice)                           // [age: 18, score: 100]，切片长度为 1

	// 4. 执行方法时，对其进行传参
	// 定义 Call 函数的形参
	var params = []reflect.Value{
		reflect.ValueOf(20),
		reflect.ValueOf(200),
	}
	v.MethodByName("SetInfo").Call(params) // 执行方法并传入参数
	infoSlice2 := v.MethodByName("GetInfo").Call(nil)
	fmt.Printf("调用 SetInfo 方法后，结构体中的字段值：%v\n", infoSlice2)

	// 5. 获取结构体中的方法数量
	fmt.Printf("NumMethod() 方法获取的方法数量：%v\n", t.NumMethod()) // NumMethod() 方法获取的方法数量：3
}

// ------------------- 修改结构体中的属性值 -------------------
func setStructField(s interface{}) {
	var t = reflect.TypeOf(s) // 获取反射中的类型对象

	// 判断形参 s 是否是结构体指针类型
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		fmt.Println("形参 s 不是结构体指针类型")
		return
	}

	// 修改结构体中指定属性的值
	var v = reflect.ValueOf(s).Elem()     // 获取反射中的值对象，并将可写性设置为 true
	v.FieldByName("Age").SetInt(888)      // 获取结构体中的 Age 属性，并修改其值
	v.FieldByName("Score").SetInt(999999) // 获取结构体中的 Score 属性，并修改其值

}

func main() {
	// 实例化 Student 结构体
	s1 := Student{
		Age:   18,
		Score: 100,
	}

	s2 := Student{
		Age:   1000,
		Score: 400,
	}

	GetStructField(s1)   // 获取结构体中的字段
	GetStructMethod(&s1) // 获取结构体中的方法，注意要传递指针变量
	fmt.Println("-----------------------------------------------------")

	fmt.Printf("修改前的 s2 属性值：%#v\n", s2) // 修改前的 s2 属性值：main.Student{Age:1000, Score:400}
	setStructField(&s2)                 // 修改结构体中的属性
	fmt.Printf("修改后的 s2 属性值：%#v\n", s2) // 修改后的 s2 属性值：main.Student{Age:888, Score:999999}
}
