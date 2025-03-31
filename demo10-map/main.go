package main

import "fmt"

func main() {
	var map1 map[string]string                   // 声明一个 map，但没有初始化
	fmt.Printf("map1 == nil: %t\n", map1 == nil) // true
	var map2 = make(map[string]string)
	fmt.Printf("map2 == nil: %t\n", map2 == nil) // false

	// 1. 声明并初始化 map
	// 1.1 使用 make() 函数
	userinfo := make(map[string]string)
	fmt.Println(userinfo)                                // map[]
	fmt.Printf("userinfo == nil: %t\n", userinfo == nil) // false
	userinfo["name"] = "zhangsan"
	userinfo["age"] = "18"
	userinfo["sex"] = "男"
	fmt.Println(userinfo) // map[age:18 name:zhangsan sex:男]

	// 1.2 使用字面量
	userinfo2 := map[string]string{
		"name": "lisi",
		"age":  "20",
		"sex":  "女",
	}
	fmt.Println(userinfo2) // map[age:20 name:lisi sex:女]

	// 2. 使用 for range 遍历 map
	userinfo3 := map[string]string{
		"name": "wangwu",
		"age":  "30",
		"sex":  "男",
	}
	for key, value := range userinfo3 {
		fmt.Printf("key: %v - value: %v\n", key, value)
	}
	// 输出：key: name - value: wangwu \n key: age - value: 30 \n key: sex - value: 男 \n

	// 3. map 数据的基本操作：CURD
	// 3.1 增加数据
	userinfo4 := make(map[string]string)
	userinfo4["name"] = "zhaoliu"
	userinfo4["age"] = "30"
	fmt.Println(userinfo4) // map[age:30 name:zhaoliu]

	// 3.2 修改数据
	userinfo4["name"] = "zhaoliu2"
	fmt.Println(userinfo4) // map[age:30 name:zhaoliu2]

	// 3.3 获取数据
	name := userinfo4["name"]
	fmt.Println(name) // zhaoliu2

	// 3.4 查找数据
	// 3.4.1 查找存在的 key
	value, ok := userinfo4["name"]
	fmt.Printf("value=%s, ok=%t\n", value, ok) // value=zhaoliu2, ok=true
	// 3.4.2 查找不存在的 key
	value, ok = userinfo4["xxx"]
	fmt.Printf("value=%s, ok=%t\n", value, ok) // value=空字符串（默认值）, ok=false

	// 3.5 删除数据
	delete(userinfo4, "name")
	fmt.Println(userinfo4) // map[age:30]

	// 4. 创建元素为 map 类型的切片
	userinfo5 := make([]map[string]string, 3, 3)                 // 初始化 userinfo5 切片
	fmt.Printf("userinfo5[0] == nil: %t\n", userinfo5[0] == nil) // userinfo5[0] == nil: true，因为 Map 元素没有初始化

	// 初始化切片中的 Map 元素
	if userinfo5[0] == nil {
		userinfo5[0] = make(map[string]string) // 初始化切片中的第一个 Map 元素
		userinfo5[0]["name"] = "zhaoliu"
	}
	if userinfo5[1] == nil {
		userinfo5[1] = make(map[string]string) // 初始化切片中的第二个 Map 元素
		userinfo5[1]["name"] = "wangwu"
	}

	fmt.Printf("userinfo5: %v\n", userinfo5) // userinfo5: [map[name:zhaoliu] map[name:wangwu] map[]]

	// 5. 定义值为 Slice 类型的 Map
	userinfo6 := make(map[string][]string) // 初始化值为切片类型的 Map
	userinfo6["name"] = []string{"zhangsan", "lisi", "wangwu"}
	userinfo6["work"] = []string{"北京", "上海", "广州"}
	fmt.Printf("userinfo6: %v\n", userinfo6) // userinfo6: map[name:[zhangsan lisi wangwu] work:[北京 上海 广州]]
}
