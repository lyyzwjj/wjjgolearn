package aha_slice

import "testing"

/**
 * @author  wjj
 * @date  2020/8/23 11:04 下午
 * @description  slice结构体  每个结构体包裹一块小数组
 */

func TestSliceInit(t *testing.T) {
	var s0 []int // 可变长数组
	t.Log(s0)
	t.Log(len(s0), cap(s0)) // 看一下长度 和 容量
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))
	s1 := []int{1, 2, 3, 4}
	t.Log(s1)
	s2 := make([]int, 3, 5) //创建一个数组 长度是3 容量是5
	t.Log(s2)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2]) // 0 0 0
	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2], s2[3], s2[4])
}

func TestSliceGrowing(t *testing.T) {
	// s := []int{}
	var s []int
	for i := 0; i < 100; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s)) // cap容量 是按照2的
	}
}

// 共享内存存储空间
func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "June", "July", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2)
	t.Log(len(Q2), cap(Q2)) // length 3 cap 9  (Apr开始到左后有9个) 和year共享
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknow"
	t.Log(Q2)
	t.Log(year)
}

func TestSliceComparing(t *testing.T) {
	// a := []int{1, 2, 3, 4}
	// b := []int{1, 2, 3, 4}
	// if a == b {    // slice只能和nil比较
	//
	// }
}
