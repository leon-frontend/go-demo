package main

import "fmt"

// 1. 自定义类型
type myInt1 int

// 2. 类型别名
type myInt2 = int

/*
3. 定义自定义结构体：结构体类型为大写时，表示公有变量；当结构体类型为小写时，表示私有变量。
*/
type Person struct {
	name string
	age  int
	sex  string
}

// 4. 给 Person 结构体添加方法
func (p Person) PrintInfo() {
	fmt.Printf("姓名：%s, 年龄：%d\n", p.name, p.age)
}

// 5. 值类型的接受者和指针类型的接受者
// 5.1 值类型的接受者修改结构体中的字段值
func (p Person) SetInfo(name string, age int) {
	p.name = name
	p.age = 100
}

// 5.2 指针类型的接受者修改结构体中的字段值
func (p *Person) SetInfo2(name string, age int) {
	p.name = name // 底层实现是 (*p).name = name
	p.age = age
}

// 6. 给自定义类型添加方法
type MyInt int

func (m MyInt) SayHello() {
	fmt.Println("Hello, 我是一个 MyInt 自定义类型中的 SayHello 方法。")
}

// 7. 结构体的匿名字段
type AnonymousPerson struct {
	string // 姓名，会默认将 string 作为字段名
	int    // 年龄，会默认将 int 作为字段名
}

// 8. 定义一个包含引用类型字段的结构体
/*
	结构体的字段类型可以是基本数据类型、切片、Map 以及结构体。
	当结构体的字段类型是指针、Slice 或 Map 时，其零值都是 nil，即还没有分配空间。
	如果需要使用这样的字段，需要先使用 make 方法手动分配空间，才能使用。
*/
type ReferencePerson struct {
	Name    string
	Hobby   []string
	Address map[string]string
}

// 9. 嵌套结构体
type Address struct {
	Province string
	City     string
}

type User struct {
	Name string
	// Address Address // Address 是一个结构体
	Address // 使用匿名字段，Address 既作为类型名也作为字段名
}

// 10. 使用嵌套结构体实现继承
// 10.1 定义父结构体 Animal
type Animal struct {
	Name string
}

// 10.2 给 Animal 结构体定义方法
func (a Animal) run() {
	fmt.Printf("%s run...\n", a.Name)
}

// 10.3 定义子结构体 Dog
type Dog struct {
	Age    int
	Animal // 使用嵌套结构体实现继承
}

// 10.4 给 Dog 结构体添加方法
func (d Dog) wang() {
	// Dog 实例 d 中不应该存在 Name 字段，但是由于使用嵌套结构体实现了继承，所以 Name 字段会从 Animal 中继承下来
	fmt.Printf("%s wang...\n", d.Name)
}

// -----------------------------------------------------------------------------------------
// -----------------------------------------------------------------------------------------
func main() {
	// 1. 自定义类型
	var a myInt1 = 10
	fmt.Printf("a value: %d, a type: %T\n", a, a) // a value: 10, a type: main.myInt1

	// 2. 类型别名
	var b myInt2 = 20
	var c int = 20
	fmt.Printf("b value: %d, b type: %T\n", b, b) // b value: 20, b type: int
	fmt.Printf("c value: %d, c type: %T\n", c, c) // c value: 20, c type: int
	fmt.Println("---------------------------------------")

	// 3. 使用 var 关键字实例化 Person 结构体
	var p1 Person
	p1.name = "张三"
	p1.age = 18
	p1.sex = "男"
	fmt.Printf("p1 value: %v, p1 type: %T\n", p1, p1)  // p1 value: {张三 18 男}, p1 type: main.Person")
	fmt.Printf("p1 value: %#v, p1 type: %T\n", p1, p1) // p1 value: main.Person{name:"张三", age:18, sex:"男"}, p1 type: main.Person

	// 4.1 使用 new 关键字实例化 Person 结构体
	var p2 *Person = new(Person)
	p2.name = "李四" // 底层实现是 (*p2).name = "李四"
	p2.age = 20
	p2.sex = "女"
	fmt.Printf("p2 value: %#v, p2 type: %T\n", p2, p2) // p2 value: &main.Person{name:"李四", age:20, sex:"女"}, p2 type: *main.Person

	// 4.2 使用 & 操作符实例化 Person 结构体
	var p3 *Person = &Person{}
	p3.name = "王五"
	p3.age = 30
	p3.sex = "男"
	fmt.Printf("p3 value: %#v, p3 type: %T\n", p3, p3) // p3 value: &main.Person{name:"王五", age:30, sex:"男"}, p3 type: *main.Person

	// 5.1 使用键值对实例化 Person 结构体
	p4 := Person{
		name: "赵六",
		age:  40,
		sex:  "女",
	}
	fmt.Printf("p4 value: %#v, p4 type: %T\n", p4, p4) // p4 value: main.Person{name:"赵六", age:40, sex:"女"}, p4 type: main.Person

	// 5.2 键值对配合 & 操作符对指针类型的结构体进行初始化
	p5 := &Person{
		name: "钱七", // 只初始化 name 字段，其他字段的值为默认值
	}
	fmt.Printf("p5 value: %#v, p5 type: %T\n", p5, p5) // p5 value: &main.Person{name:"钱七", age:0, sex:""}, p5 type: *main.Person

	// 6. 使用值列表初始化，即初始化结构体的时候不写键，直接写值
	p6 := Person{"孙八", 50, "男"}
	fmt.Printf("p6 value: %#v, p6 type: %T\n", p6, p6) // p6 value: main.Person{name:"孙八", age:50, sex:"男"}, p6 type: main.Person

	// 7. 结构体是值类型
	p7 := Person{
		name: "周九",
		age:  60,
		sex:  "男",
	}
	p8 := p7
	p8.name = "吴十"
	fmt.Printf("p7 value: %#v, p7 type: %T\n", p7, p7) // p7 value: main.Person{name:"周九", age:60, sex:"男"}, p7 type: main.Person
	fmt.Printf("p8 value: %#v, p8 type: %T\n", p8, p8) // p8 value: main.Person{name:"吴十", age:60, sex:"男"}, p8 type: main.Person
	fmt.Println("---------------------------------------")

	// 8. 给 Person 结构体添加方法
	p6.PrintInfo()

	// 9.1 值类型的接受者修改结构体中的字段值
	fmt.Printf("p6 value: %#v, p6 type: %T\n", p6, p6) // p6 value: main.Person{name:"孙八", age:50, sex:"男"}, p6 type: main.Person
	p6.SetInfo("郑十一", 70)
	fmt.Printf("p6 value: %#v, p6 type: %T\n", p6, p6) // p6 value: main.Person{name:"孙八", age:50, sex:"男"}, p6 type: main.Person

	// 9.2 指针类型的接受者修改结构体中的字段值
	fmt.Printf("p7 value: %#v, p7 type: %T\n", p7, p7) // p7 value: main.Person{name:"周九", age:60, sex:"男"}, p7 type: main.Person
	p7.SetInfo2("郑十一", 70)
	fmt.Printf("p7 value: %#v, p7 type: %T\n", p7, p7) // p7 value: main.Person{name:"郑十一", age:70, sex:"男"}, p7 type: main.Person
	fmt.Println("---------------------------------------")

	// 10. 给自定义类型添加方法
	var m1 MyInt = 10
	m1.SayHello() // Hello, 我是一个 MyInt 自定义类型中的 SayHello 方法。

	// 11. 结构体的匿名字段
	a1 := AnonymousPerson{
		"张三", // 匿名字段的值
		18,   // 匿名字段的值
	}
	fmt.Printf("a1 value: %#v, a1 type: %T\n", a1, a1) // a1 value: main.AnonymousPerson{string:"张三", int:18}, a1 type: main.AnonymousPerson

	// 12. 实例化包含引用类型字段的结构体
	var r1 ReferencePerson
	// 12.1 给基本数据类型赋值
	r1.Name = "张三"
	// 12.2 给切片赋值，要么使用字面量形式，要么使用 make 方法
	r1.Hobby = []string{"足球", "篮球"}
	// 12.3 给 Map 赋值，要么使用字面量形式，要么使用 make 方法
	r1.Address = map[string]string{
		"province": "北京",
		"city":     "北京",
	}
	fmt.Printf("r1 value: %#v, r1 type: %T\n", r1, r1) // r1 value: main.ReferencePerson{Name:"张三", Hobby:[]string{"足球", "篮球"}, Address:map[string]string{"city":"北京", "province":"北京"}}, r1 type: main.ReferencePerson

	// 13. 实例化嵌套结构体
	var u1 User
	u1.Name = "张三"
	// u1.Address.Province = "北京"
	// u1.Address.City = "北京"
	// 若在 Address 是匿名字段，则通过 u1 访问 Address 中的字段时，可以省略 .Address
	u1.Province = "北京" // 当访问结构体成员时会先在结构体中查找该字段，找不到再去匿名结构体中查找。
	u1.City = "北京"
	fmt.Printf("u1 value: %#v, u1 type: %T\n", u1, u1) // u1 value: main.User{Name:"张三", Address:main.Address{Province:"北京", City:"北京"}}, u1 type:main.User

	// 14. 使用嵌套结构体实现继承
	var d1 = Dog{
		Age: 10,
		Animal: Animal{
			Name: "旺财",
		},
	}
	// 14.1 Dog 实例访问 Animal 结构体中的 run 方法
	d1.run()  // 旺财 run...
	d1.wang() // 旺财 wang...
}
