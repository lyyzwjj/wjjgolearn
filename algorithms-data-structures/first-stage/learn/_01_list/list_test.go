package list

import (
	"fmt"
	"github.com/lyyzwjj/wjjgolearn/algorithms-data-structures/common"
	"testing"
)

func initObjList[T common.Comparable[T]]() ObjList[T] {
	var list ObjList[T]
	//  TestArrayList
	//	list = NewArrayList()
	//  TestLinkedList
	list = NewLinkedObjList[T]()
	return list
}

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

type Name struct {
	FirstName string
}

type Animal struct {
	age  int
	name Name
}

func (a Animal) Equals(o Animal) bool {
	return a.age == o.age
}

func (a Animal) CompareTo(o Animal) int {
	return a.age - o.age
}

func NewAnimal(age int) Animal {
	return Animal{age: age}
}

func TestObjList(t *testing.T) {
	list := initObjList[Animal]()
	// list := NewLinkedObjList[Animal]()
	// list := initList[Animal]()
	// list := NewLinkedList[Animal]()

	for i := 0; i <= 10; i++ {
		//var animal Comparable[Animal]
		//animal = NewAnimal(i)
		if i != 3 {
			list.Add(NewAnimal(i))
		} else {
			list.Add(Animal{i, Name{"haha"}})
		}
	}
	fmt.Printf("%#v\n", list)
	fmt.Println(list.Get(0))
	fmt.Println(list.Remove(0))
	fmt.Println(list.Get(2))
	fmt.Println(list.Contains(NewAnimal(3)))
	fmt.Println(list.Contains(Animal{3, Name{"haha"}}))
	fmt.Println(list.Remove(2))
	fmt.Println(list.Contains(NewAnimal(2)))
	fmt.Println(list.Size())
	list.Clear()
	fmt.Printf("%#v\n", list)
}

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
