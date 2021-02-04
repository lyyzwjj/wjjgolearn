package baa_object_oriented

import (
	"fmt"
	"testing"
	"unsafe"
)

/**
 * @author  wjj
 * @date  2020/8/29 1:25 上午
 * @description 面向对象
 */

// 定义结构体
type Employee struct {
	Id   string
	Name string
	Age  int
}

// 此种定义在实例对应方法被调用时,实例的成员会进行值复制
//func (e Employee) String() string {
//	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))  // 两者内存地址是不一样的
//	return fmt.Sprintf("ID:%s-Name;%s-Age:%d", e.Id, e.Name, e.Age)
//}

// 为了避免内存拷贝我们使用如下定义方式  方法定义在指针上
func (e *Employee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))          // 两者内存地址是一样的
	return fmt.Sprintf("ID:%s/Name;%s/Age:%d", e.Id, e.Name, e.Age) // 此处的e.Id e.Name e.Age进行值复制
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	// e := &Employee{"0", "Bob", 20}
	e1 := Employee{Name: "Bob", Age: 20}
	e2 := new(Employee) // 返回指针  相当于 e:=&Employee{}

	e2.Id = "2" // 指针可以向对象一样直接操作  不需要->
	e2.Age = 22
	e2.Name = "Rose"
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	t.Logf("e is %T", e) // %T 表示类型
	// t.Logf("e is %T", &e) // 和下面的一致
	t.Logf("e2 is %T", e2)
}

func TestStructOperations(t *testing.T) {
	e := &Employee{"0", "Bob", 20}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name)) // 两者内存地址是一样的
	t.Log(e.String())
}
