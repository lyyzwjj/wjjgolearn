package main

import "fmt"

// append() 为切片追加元素

func main() {
	s1 := []string{"北京", "上海", "深圳"}
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n ", s1, len(s1), cap(s1))
	// s1[3] = "广州"	index out of range 索引越界
	fmt.Println(s1)
	// 调用append函数必须使用原来的切片变量接收返回值	对于 s1 指向的 内存地址换了
	s2 := append(s1, "广州") //append追加元素,原来的底层数组放不下的时候, Go底层就会把底层数组换一个
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n ", s2, len(s2), cap(s2))
	// s1 = append(s1, "广州")
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n ", s1, len(s1), cap(s1)) // 原本的数组没有变
	// 调用append函数必须使用原来的切片变量接收返回值	对于 s1 指向的 内存地址换了不改变指向 还是原来那个s1
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n ", s1, len(s1), cap(s1)) // 原本的数组没有变
	// 必须用变量接收返回值
	s1 = append(s1, "广州")
	// 扩容逻辑		/usr/local/go/src/runtime/slice.go 下 小于1024 直接就翻倍 大于1024 1.25倍增加
	/*
		newcap := old.cap
		doublecap := newcap + newcap
		if cap > doublecap {
			newcap = cap
		} else {
			if old.cap < 1024 {
				newcap = doublecap
			} else {
				// Check 0 < newcap to detect overflow
				// and prevent an infinite loop.
				for 0 < newcap && newcap < cap {
					newcap += newcap / 4
				}
				// Set newcap to the requested cap when
				// the newcap calculation overflowed.
				if newcap <= 0 {
					newcap = cap
				}
			}
		}
	*/
	s1 = append(s1, "杭州", "成都")
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n ", s1, len(s1), cap(s1)) // 原本的数组没有变
	s1 = append(s1, "武汉", "西安", "苏州")                                  // 扩容了 cap 变12了
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n ", s1, len(s1), cap(s1)) // 原本的数组没有变
	s1 = append(s1, "武汉", "西安", "苏州")
	ss := []string{"武汉", "西安", "苏州"}                                   // 切片
	s1 = append(s1, ss...)                                             // ... 表示将切片拆开
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n ", s1, len(s1), cap(s1)) // 原本的数组没有变
}
