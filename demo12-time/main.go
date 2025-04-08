package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 打印当前日期
	timeObj := time.Now()
	fmt.Println("当前日期:", timeObj) // 当前日期: 2025-04-08 16:20:38.1691004 +0800 CST m=+0.000000001

	// 1.1 获取年月日
	fmt.Println("年:", timeObj.Year())   // 年: 2025
	fmt.Println("月:", timeObj.Month())  // 月: 4
	fmt.Println("日:", timeObj.Day())    // 日: 8
	fmt.Println("时:", timeObj.Hour())   // 时: 16
	fmt.Println("分:", timeObj.Minute()) // 分: 20
	fmt.Println("秒:", timeObj.Second()) // 秒: 38
	fmt.Println("-------------------------------------------")

	// 2. 将当前时间（默认格式）转换为时间戳
	timestamp := timeObj.Unix()      // 获取当前时间戳（秒）
	fmt.Println("当前时间戳:", timestamp) // 当前时间戳: 1712545238

	// 2.1 将时间戳转换为日期字符串
	// Unix 方法的第二个参数用于指定时区，如果不指定时区，则默认为 UTC 时区
	// 格式化模板解释：2006-01-02 15:04:05，其中 2006 是年份，01 是月份，02 是日期，15 是小时（24 小时制），04 是分钟，05 是秒钟
	timestampStr := time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
	fmt.Println("时间戳转换为日期字符串:", timestampStr) // 时间戳转换为日期字符串: 2025-04-08 16:20:38

	// 2.2 将日期字符串转换为时间戳：先将日期字符串转换为默认格式的时间对象，再将时间对象转换为时间戳
	dateStr := "2025-04-08 16:20:38"               // 日期字符串要与格式化模板一致，否则转换失败
	template := "2006-01-02 15:04:05"              // 时间格式模板
	timeObj1, err := time.Parse(template, dateStr) // 将日期字符串转换为时间对象
	if err != nil {
		fmt.Println("日期字符串转换为时间戳失败:", err)
	} else {
		timestamp1 := timeObj1.Unix()           // 将时间对象转换为时间戳
		fmt.Println("日期字符串转换为时间戳:", timestamp1) // 日期字符串转换为时间戳: 1712545238
	}
	fmt.Println("-------------------------------------------")

	// 3. 时间间隔表示一段时间的长度,  time 包中定义了一些时间间隔的常量
	fmt.Println("1秒:", time.Second)  // 1秒: 1s
	fmt.Println("1分钟:", time.Minute) // 1分钟: 1m0s
	fmt.Println("1小时:", time.Hour)   // 1小时: 1h0m0s
	fmt.Println("-------------------------------------------")

	// 4. 定时器
	// 4.1 使用 time.NewTicker 创建一个定时器，每隔 1 秒钟执行一次操作
	ticker := time.NewTicker(1 * time.Second)
	stopNum := 5 // 停止定时器的次数
	// 通过循环遍历 ticker.C 获取定时器中的时间值，并且每隔 1 秒执行一次循环体中的操作
	for t := range ticker.C {
		// fmt.Println("当前时间:", t) // 每 1 秒执行一次，t 的值是默认格式的当前时间
		fmt.Println("当前时间:", t.Format("2006-01-02 15:04:05")) // 每 1 秒执行一次
		stopNum--
		if stopNum == 0 {
			ticker.Stop() // 销毁定时器
			break
		}
	}

	// 4.2 使用 time.Sleep 暂停当前程序执行 n 秒
	fmt.Println("测试 time.Sleep")
	time.Sleep(3 * time.Second) // 暂停 3 秒
	fmt.Println("暂停 3 秒后继续执行")
}
