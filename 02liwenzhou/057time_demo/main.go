package main

import (
	"fmt"
	"time"
)

// 时间
// 时间戳是自1970年1月1日 (08:00:00GMT)至当前时间的总毫秒数,它也被称为Unix时间戳(UnixTimestamp)

// 时区
func f2() {
	now := time.Now() // 本地的时间
	fmt.Println(now)
	// 明天的这个时间
	parseDirect, err := time.Parse("2006-01-02 15:04:05", "2021-04-07 12:12:50")
	if err != nil {
		fmt.Printf("parse time failed, err:%v\n", err)
		return
	}
	fmt.Println(parseDirect) // 2021-04-07 12:12:50 +0000 UTC
	// 按照东八区的时区和格式去解析一个字符串格式的时间
	// 根据字符串加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("load loc failed, err:%v\n", err)
		return
	}
	// 按照指定时区解析时间
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2021-04-07 01:21:50", loc)
	if err != nil {
		fmt.Printf("parse time failed, err:%v\n", err)
		return
	}
	fmt.Println(timeObj)   // 2021-04-07 12:12:50 +0800 CST
	td := timeObj.Sub(now) // 24h0m24.923548s
	fmt.Println(td)
}

func main() {
	f2()
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Date())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	// 时间戳
	// 秒	1617639859
	fmt.Println(now.Unix())
	// 纳秒	1617639859724311000
	fmt.Println(now.UnixNano())

	// time.Unix()	Unix创建一个本地时间，对应sec和nsec表示的Unix时间（从January 1, 1970 UTC至该时间的秒数和纳秒数） nsec的值在[0, 999999999]范围外是合法的。
	ret := time.Unix(1617639859, 0)
	fmt.Println(ret)
	fmt.Println(ret.Year())
	fmt.Println(ret.Day())

	// 时间间隔	类型 type Duration int64
	fmt.Println(time.Second)
	// now + 1小时
	fmt.Println(now.Add(24 * time.Hour))
	nextYear, err := time.Parse("2006-01-02", "2021-04-07") // 取的是2021-04-06
	if err != nil {
		fmt.Printf("parse time failed,err:%v\n", err)
		return
	}
	// Sub 两个时间相减
	fmt.Println("-------- sub --------")
	d := now.Sub(nextYear) // -31h2m30.814596s
	fmt.Println(d)

	// 格式化时间对象转换成字符串类型的时间
	// 2006年1月2日15点04分05秒 000毫秒 Mon Jan
	// Y-m-d H:M:S
	// 2016-01-02 15:04:05.000
	// 2021-04-06
	fmt.Println(now.Format("2006-01-02"))
	// 2021-04-06 00:45:04
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	// 2021-04-06 12:45:44 AM
	fmt.Println(now.Format("2006-01-02 03:04:05 PM"))
	// 2021-04-06 00:46:26.504
	fmt.Println(now.Format("2006-01-02 15:04:05.000"))

	// 按照对应的格式解析字符串类型的时间
	// 1617641476
	// 1617667200		// 取的是2021-04-06 08:00:00 AM
	timeObj, err := time.Parse("2006-01-02", "2021-04-06") // 取的是2021-04-06
	if err != nil {
		fmt.Printf("parse time failed,err:%v\n", err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())

	// duration := time.Duration(1 << 31)
	duration := time.Second * 3
	fmt.Print(duration) // 4.294967296s
	time.Sleep(duration)
	// 没有实现的方法  func Sleep(d Duration)
	time.Sleep(10 * 31)
	n := 100
	// time.Sleep(n) 这样传类型不匹配了
	time.Sleep(time.Duration(n))

	// 定时器
	// Timer()
}

func Timer() {
	// 定时器
	timer := time.Tick(time.Second)
	for t := range timer {
		fmt.Println(t) // 1秒钟执行一次
	}
}
