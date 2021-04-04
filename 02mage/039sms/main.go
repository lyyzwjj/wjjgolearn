package main

import (
	"fmt"
	"os"
)

/*
	函数版学生管理系统
	写一个系统能查看\新增\删除学生
*/

type Student struct {
	id   int64
	name string
}

func newStudent(id int64, name string) Student {
	return Student{
		id:   id,
		name: name,
	}
}

var (
	allStudent map[int64]*Student
)

func main() {
	allStudent = make(map[int64]*Student, 48) // 初始化(开辟内存空间)
	for {
		// 1. 打印菜单
		fmt.Println("欢迎光临学生管理系统!")
		fmt.Println(`
	1. 查看所有学生
	2. 新增学生
	3. 删除学生
	4. 退出
	`)
		fmt.Println("请输入你要干啥: ")
		// 2. 等待用户选择要做什么
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d选项!\n", choice)
		// 执行对应的函数
		switch choice {
		case 1:
			showAllStudents()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1) // 退出程序状态码
		default:
			fmt.Println("滚")
		}
	}
}

func deleteStudent() {
	// 1. 请用户输入要删除的学生的序号
	var deleteId int64
	fmt.Print("请输入要删除的学生号")
	fmt.Scanln(&deleteId)
	// 去allStudent这个map中根据学号删除对应的键值对
	delete(allStudent, deleteId)
}

func addStudent() {
	// 向allStudent中添加一个新的学生
	// 1. 创建一个新的学生
	// 1.1 获取用户输入
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学生学号: ")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名: ")
	fmt.Scanln(&name)
	// 1.2 造学生(调用student的构造函数)
	newStu := newStudent(id, name)
	// 2 追加到allStudent这个map中
	allStudent[id] = &newStu
}

func showAllStudents() {
	// 把所有学生打印出来
	for k, v := range allStudent {
		fmt.Printf("学号:%d 姓名:%s\n", k, v.name)
	}
}
