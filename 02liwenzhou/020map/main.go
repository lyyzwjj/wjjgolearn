package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// map

func main() {
	deleteInTravel()
	var m1 map[string]int
	fmt.Println(m1 == nil)        // 还没有初始化 (没有在内存中开辟空间)
	m1 = make(map[string]int, 10) // make 初始化 要估算好该map容量,避免在程序运行期间再动态扩容
	m1["理想"] = 18
	m1["jiwuming"] = 35
	fmt.Println(m1)
	fmt.Println(m1["理想"])
	fmt.Println(m1["娜扎"]) // 没有不存在这个key拿到对应值类型的零值	int即是0	打印的是0
	// 约定成俗用ok接收返回的布尔值
	value, ok := m1["娜扎"]
	if !ok {
		fmt.Println("查无此key")
	} else {
		fmt.Println(value)
	}

	// map的遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	// 只遍历key
	for k := range m1 { //只需要遍历key
		fmt.Println(k)
	}
	// 只遍历value
	for _, value := range m1 { //	只需要遍历key
		fmt.Println(value)
	}

	// 删除
	delete(m1, "jiwuming")
	fmt.Println(m1)
	delete(m1, "沙河") // 删除不存在的key
	// www.studygolang.com/pkgdoc
	// go doc builtin.delete	看go语言的文档
	/*
		package builtin // import "builtin"

		func delete(m map[Type]Type1, key Type)
		    The delete built-in function deletes the element with the specified key
		    (m[key]) from the map. If m is nil or there is no such element, delete is a
		    no-op.
		内建函数delete按照指定的键将元素从映射中删除。若m为nil或无此元素，delete不进行操作。
	*/
	getMapValueBySortedKey()
}
func getMapValueBySortedKey() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子
	var scoreMap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}
	// 取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	// 对切片进行排序
	sort.Strings(keys)
	// 按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}

func deleteInTravel() {
	m := make(map[int]int)
	for i := 0; i < 10; i++ {
		m[i] = i
	}
	fmt.Printf("%#v\n", m)
	for k := range m {
		if k == 3 {
			delete(m, k)
		}
	}
	fmt.Printf("%#v\n", m)
}
