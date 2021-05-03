package main

import (
	"flag"
	"fmt"
	"time"
)

// flag 获取命令行参数

func main() {
	// ./main -name wangzhe -age 26
	// ./main -md=10000h
	// ./main -ct=10000h
	// 创建一个标志位参数  标志  默认值 提示
	name := flag.String("name", "王冶", "请输入名字")
	age := flag.Int("age", 9000, "请输入真实年龄")
	married := flag.Bool("married", false, "结婚了吗")
	cTime := flag.Duration("ct", time.Second, "结婚多久了")
	// 使用flag
	// ./main -name=chenchen -age=25 -married=false -ct=100000h a b c
	flag.Parse()
	fmt.Println(*name) // & 取地址 * 取值
	fmt.Println(*age)
	fmt.Println(*married)
	fmt.Println(*cTime)
	fmt.Printf("%T\n", *cTime)
	fmt.Println(flag.Args())  // 返回命令行参数后的其他参数,以[]string类型
	fmt.Println(flag.NArg())  // 返回命令行参数后的其他参数个数
	fmt.Println(flag.NFlag()) // 返回使用的命令行参数个数
	//var name string
	//var age int
	//var married bool
	//var delay time.Duration
	//flag.StringVar(&name, "name", "张三", "姓名")
	//flag.IntVar(&age, "age", 18, "年龄")
	//flag.BoolVar(&married, "married", false, "婚否")
	//flag.DurationVar(&delay, "d", 0, "时间间隔")
	// 使用flag
	// flag.Parse()
	// fmt.Println(name) // & 取地址 * 取值

}
