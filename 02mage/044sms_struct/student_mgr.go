package main

import "fmt"

// 有一个物件:
// 1. 它保存了一些数据  ---> 结构体的字段
// 2. 它有4个功能		---> 结构体的方法

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

// 造一个学生的管理者
type StudentMgr struct {
	allStudent map[int64]Student
}

// 查看学生
func (s StudentMgr) ShowStudents() {
	// 从s.allStudent这个map中把所有的学生逐个拎出来
	for _, stu := range s.allStudent { // stu是具体每一个学生
		fmt.Printf("学号:%d 姓名:%s\n", stu.id, stu.name)
	}
}

// 增加学生
func (s StudentMgr) AddStudent() {
	var (
		id   int64
		name string
	)
	// 获取用户输入
	fmt.Print("请输入学号: ")
	fmt.Scanln(&id)
	fmt.Print("请输入姓名: ")
	fmt.Scanln(&name)
	// 根据用户输入创建结构体对象
	newStu := newStudent(id, name)
	s.allStudent[newStu.id] = newStu
	// 2. 把新的学生放到s.allStudent这个map中

}

// 修改学生
func (s StudentMgr) EditStudent() {
	// 1. 获取用户输入的学号
	var stuId int64
	fmt.Println("请输入学号: ")
	fmt.Scanln(&stuId)
	// 2. 展示该学号对应的学生信息,如果没有提示查无此人
	stuObj, ok := s.allStudent[stuId]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	fmt.Printf("你要修改的学生信息如下:学号:%d 姓名:%s\n", stuObj.id, stuObj.name)
	// 3. 请输入修改后的学生名
	fmt.Println("请输入学生的新名字:")
	var newName string
	fmt.Scanln(&newName)
	// 4. 更新学生的姓名
	stuObj.name = newName
	s.allStudent[stuId] = stuObj
}

// 删除学生
func (s StudentMgr) DeleteStudent() {
	// 1. 请用户输入要删除的学生id
	var stuId int64
	fmt.Println("请用户输入要删除的学生学号: ")
	fmt.Scanln(&stuId)
	// 2. 展示该学号对应的学生信息,如果没有提示查无此人
	_, ok := s.allStudent[stuId]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	// 3. 有的话就删除,如何从map中删除键值对
	delete(s.allStudent, stuId)
	fmt.Println("删除成功")
}
