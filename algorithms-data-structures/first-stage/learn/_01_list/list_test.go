package list

import (
	"fmt"
	"testing"
)

func initList[T comparable]() List[T] {
	var list List[T]
	//  TestArrayList
	//	list = NewArrayList()
	//  TestLinkedList
	list = NewLinkedList[T]()
	return list
}

func TestList(t *testing.T) {
	list := initList[int]()
	for i := 0; i <= 10; i++ {
		list.Add(i)
	}
	fmt.Printf("%#v\n", list)
	//var obj int
	//list.GetObj(1, &obj)
	//fmt.Println(obj)
	fmt.Println(list.Get(2))
	fmt.Println(list.Contains(2))
	fmt.Println(list.Remove(2))
	fmt.Println(list.Contains(2))
	fmt.Println(list.Size())
	list.Clear()
	fmt.Printf("%#v\n", list)

}

//
//type Animal struct {
//	age int
//}
//
//func NewAnimal(age int) Animal {
//	return Animal{age: age}
//}
//func NewAnimal1(age int) *Animal {
//	return &Animal{age: age}
//}
//
//func TestListStruct(t *testing.T) {
//	// list := initList()
//	list := NewLinkedList[int]()
//	for i := 0; i <= 10; i++ {
//		list.Add(NewAnimal(i))
//	}
//	fmt.Printf("%#v\n", list)
//	fmt.Println(list.Get(0))
//	fmt.Println(list.Remove(0))
//	fmt.Println(list.Get(2))
//	fmt.Println(list.Contains(NewAnimal(2)))
//	fmt.Println(list.Remove(2))
//	fmt.Println(list.Contains(NewAnimal(2)))
//	fmt.Println(list.Size())
//	list.Clear()
//	fmt.Printf("%#v\n", list)
//}
//
//// 不能有指针
//func TestListStructPointer(t *testing.T) {
//	// list := initList()
//	list := NewLinkedList[int]()
//	for i := 0; i <= 10; i++ {
//		list.Add(NewAnimal1(i))
//	}
//	fmt.Printf("%#v\n", list)
//	fmt.Println(list.Get(2))
//	fmt.Println(list.Contains(NewAnimal1(2)))
//	fmt.Println(list.Remove(2))
//	fmt.Println(list.Contains(NewAnimal1(2)))
//	fmt.Println(list.Size())
//	list.Clear()
//	fmt.Printf("%#v\n", list)
//}
//
//type demo struct {
//	weight int
//}
//
//func f1() interface{} {
//	return &demo{weight: 1}
//}
//
//// 不能赋值
//func f2(obj interface{}) {
//	obj = f1()
//	return
//}
//func TestGetObj(t *testing.T) {
//	var obj demo
//	f2(&obj)
//	fmt.Printf("%v", obj)
//}
