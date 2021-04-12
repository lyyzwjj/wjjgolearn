package main

import (
	"fmt"
	"reflect"
)

// reflect
// type	类型
// kind	种类

/*
type Kind uint

const (
	Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Ptr
	Slice
	String
	Struct
	UnsafePointer
)

*/

// 取变量值 v.Int()
// 设置变量值 Elem()

type Cat struct {
}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type name :%v type kind:%v\n", v.Name(), v.Kind())
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值,然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %f\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值,然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值,然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

// 通过反射设置变量的值
func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) // 修改的是副本,reflect包会引发panic
	}
}

// 想要在函数中通过反射修改变量的值,需要注意函数参数传递的是值拷贝,必须传递变量地址才能修改变量值.而反射中使用专有的Elem()方法来获取指针对应的值
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

// IsNil()
// IsNil()报告v持有的值是否为nil,v持有的值的分类必须是通道\函数\接口\映射\指针\切片之一(引用类型,非简单类型);否则IsNil函数会导致panic
// IsValid()
// IsValid()返回v是否持有一个值.如果v是Value零值会返回假,此时v除了IsValid,String,Kind之外的方法都会导致panic

func TestIsNil() {
	// *int类型空指针
	var a *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
}

func TestIsValid() {
	// nil值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	// 实例化一个匿名结构体
	b := struct{}{}
	// 尝试从结构体中查找"abc"字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())
	// 尝试从结构体中查找"abc"方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid())
	// map
	c := map[string]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map中不存在的键：", reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")).IsValid())
}

func main() {
	var a float32 = 3.14
	reflectType(a) // type:float32
	var b int64 = 100
	reflectType(b) // type:int64
	var c = Cat{}
	reflectType(c)
	reflectValue(a)
	// reflectSetValue1(b)
	// reflectSetValue1(&b) // 也没有用
	reflectSetValue2(&b) // 也没有用
	fmt.Println(b)
	TestIsNil()
	TestIsValid()
}
