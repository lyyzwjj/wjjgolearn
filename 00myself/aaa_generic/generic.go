package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// go run -gcflags=-G=3 generic.go

//	泛型基本语法
func printNormal[T any](s []T) {
	for _, v := range s {
		fmt.Printf(" type: %T,value: %v", v, v)
	}
	fmt.Print("\n")
}
func normal() {
	printNormal[int]([]int{66, 77, 88, 99, 100})
	printNormal[float64]([]float64{1.1, 2.2, 3.3, 4.4, 5.5})
	printNormal[string]([]string{"红烧肉", "清蒸鱼", "大闸蟹", "九转大肠", "重烧海参"})
	printNormal([]int64{55, 44, 33, 22, 11})
}

/*
	[T any]参数的类型，意思是该函数支持任何T类型；
	在调用这个泛型函数的时候，可以显示指定类型参数，
	如：printSlice[int] ([]int{66, 77, 88, 99, 100})
	也可以省略显示类型
	比如 printSlice([]int64{55, 44, 33, 22, 11})
*/
func printSlice[T any](s []T) {
	for _, v := range s {
		fmt.Printf(" type: %T,value: %v", v, v)
	}
	fmt.Print("\n")
}

func slice() {
	printSlice[int]([]int{66, 77, 88, 99, 100})
	printSlice[float64]([]float64{1.1, 2.2, 3.3, 4.4, 5.5})
	printSlice[string]([]string{"红烧肉", "清蒸鱼", "大闸蟹", "九转大肠", "重烧海参"})
	printSlice([]int64{55, 44, 33, 22, 11})
}

/*
	泛型函数
	Addable，新增了类型列表表达式，它是对类型参数进行约束。
*/
type Addable interface {
	type int, int8, int16, int32, int64, uint, uint8, uint16, uint32,
		uint64, uintptr, float32, float64, complex64, complex128, string
}

func printMethodAdd[T Addable] (a, b T) T { return a + b }

func method() {
	fmt.Println(printMethodAdd(3, 4))
	fmt.Println(printMethodAdd("Go", "lang"))
}

/*
	泛型 使用interface约束泛型
	不仅在interface中显示声明类型约束外，还可以在函数的入参中使用接口对泛型类型的约束。
	下面声明了ShowPrice接口，并且使用Price类型实现了此接口。
*/
type Price int

type ShowPrice interface {
	String() string
}

func (i Price) String() string {
	return strconv.Itoa(int(i))
}

func showPriceList[T ShowPrice](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return ret
}

func priceList() {
	ret := showPriceList([]Price{48, 88, 152, 219, 328})
	fmt.Println(reflect.TypeOf(ret))
	fmt.Println(ret)
}

/*
	类型列表和方法列表双重约束
*/
type ShowPrice2 interface {
	type int, int8, int16, int32, int64
	String() string
}

func showPriceList2[T ShowPrice](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return ret
}
func priceList2() {
	// ret := showPriceList([]Price{48.1, 88.1, 152.1, 219.1, 328.1})
	//ret := showPriceList([]Price{48.0, 88.0, 152.0, 219.0, 328.0})
	ret := showPriceList([]Price{48, 88, 152, 219, 328})
	fmt.Println(reflect.TypeOf(ret))
	fmt.Println(ret)
}

/*
	使用comparable内置类型比较
	Go语言对原生类型比较的支持，整型，字符串，以及我们自定义的结构体都可以直接作为实参传给由comparable约束的类型参数。
	comparable可以理解为由Go编译器特殊梳理的、包含由所有内置可比较类型组成的类型列表(type list)的interface类型.
	因此也可以被嵌套在其他作为约束的接口类型定义中。
	type XXX interface{
	  comparable
	  ...
	}
*/

func index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

type Food struct {
	Name  string
	Price int
}

func innerMethodComparable() {
	fmt.Println(index([]int{11, 22, 33, 44, 55}, 55))
	fmt.Println(index([]string{"红烧肉", "清蒸鱼", "大闸蟹", "九转大肠", "烤全羊"}, "九转大肠"))
	fmt.Println(
		index([]Food{
			{"红烧肉", 1},
			{"清蒸鱼", 2},
			{"大闸蟹", 3},
			{"九转大肠", 4},
			{"烤全羊", 5},
		}, Food{"清蒸鱼", 2}))
}

/*
	泛型Set类型（泛型方法）
*/

type addable interface {
	comparable
}

type set[T addable] map[T]struct{}

func (s set[T]) add(v T) {
	s[v] = struct{}{}
}

func (s set[T]) contains(v T) bool {
	_, ok := s[v]
	return ok
}
func (s set[T]) len() int   { return len(s) }
func (s set[T]) delete(v T) { delete(s, v) }
func (s set[T]) iterate(f func(T)) {
	for v := range s {
		f(v)
	}
}

func print[T addable](s T) {
	fmt.Printf("%v ", s)
}

func innerMethodSet() {
	s := make(set[string])
	s.add("红烧肉")
	s.add("清蒸鱼")
	s.add("九转大肠")
	s.add("大闸蟹")
	s.add("烤羊排")
	fmt.Printf("%v\n", s)
	if s.contains("大闸蟹") {
		println("包含大闸蟹")
	} else {
		println("不包含大闸蟹")
	}
	fmt.Printf("the len of set: %d\n", s.len())
	s.delete("大闸蟹")
	fmt.Println("\nafter delete 大闸蟹:")
	if s.contains("大闸蟹") {
		println("包含大闸蟹")
	} else {
		println("不包含大闸蟹")
	}
	fmt.Printf("the len of set: %d\n", s.len())
	s.iterate(func(x string) { fmt.Println("您点的菜: " + x) })
}

/*
	泛型Set类型（泛型方法）
*/
type queue[T any] []T

func (q *queue[T]) enqueue(v T) {
	*q = append(*q, v)
}
func (q *queue[T]) dequeue() (T, bool) {
	if len(*q) == 0 {
		var zero T
		return zero, false
	}
	r := (*q)[0]
	*q = (*q)[1:]
	return r, true
}

func innerMethodQueue() {
	q := new(queue[string])
	q.enqueue("红烧肉")
	q.enqueue("清蒸鱼")
	fmt.Println(q)
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
}

/*
	泛型HashMap
*/
type hashmap[K comparable, V any] map[K]V

func (h hashmap[K, V]) put(k K, v V) {
	h[k] = v
}
func (h hashmap[K, V]) get(k K) (v V) {
	return h[k]
}

type Person struct {
	Id   int
	Name string
}

func (p *Person) showName() {
	fmt.Println(p.Name)
}
func myHashMap() {
	h1 := make(hashmap[string, int], 10)
	key1 := "haha"
	h1.put(key1, 100)
	value1 := h1.get(key1)
	fmt.Printf("K type: %T, K value: %v V type: %T, V value: %v \n", key1, key1, value1, value1)
	h2 := make(hashmap[int, string], 10)
	key2 := 100
	h2.put(key2, "haha")
	value2 := h2.get(key2)
	fmt.Printf("K type: %T, K value: %v V type: %T, V value: %v \n", key2, key2, value2, value2)
	h3 := make(hashmap[int, Person], 10)
	key3 := 100
	h3.put(key3, Person{
		Id:   1,
		Name: "张三",
	})
	value3 := h3.get(key3)
	value3.showName()
	fmt.Printf("K type: %T, K value: %v V type: %T, V value: %v \n", key3, key3, value3, value3)
}
func main() {
	//normal()
	//slice()
	//method()
	//priceList()
	//priceList2()
	//innerMethodComparable()
	//innerMethodSet()
	//innerMethodQueue()
	myHashMap()
}
