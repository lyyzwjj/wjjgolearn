package main

import (
	"fmt"
	"os"
)

// 学生管理系统
// 菜单函数

var smr StudentMgr

func showMenu() {
	fmt.Println("----------welcome sms----------")
	fmt.Println(`
	1. 查看所有学生
	2. 新增学生
	3. 修改学生
	4. 删除学生
	5. 退出
	`)
}

func main() {
	smr = StudentMgr{
		allStudent: make(map[int64]Student, 100),
	}
	for {
		showMenu()
		// 等待用户输入选项
		fmt.Println("请输入序号: ")
		// 2. 等待用户选择要做什么
		var choice int
		fmt.Scanln(&choice)
		fmt.Println("你输入的是: ", choice)

		// 执行对应的函数
		switch choice {
		case 1:
			smr.ShowStudents()
		case 2:
			smr.AddStudent()
		case 3:
			smr.EditStudent()
		case 4:
			smr.DeleteStudent()
		case 5:
			os.Exit(1) // 退出程序状态码
		default:
			fmt.Println("滚~")
		}
	}
}
