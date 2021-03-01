package aga_array

import "testing"

/**
 * @author  wjj
 * @date  2020/8/22 7:22 下午
 * @description
 */

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 3, 4, 5} // 不想去数元素个数
	arr1[1] = 5
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr3)

	arr4 := [2][3]int{{1, 2, 0}, {3, 4, 9}}
	t.Log(arr4)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	for idx, e := range arr3 {
		t.Log(idx, e)
	}

	for _, e := range arr3 { // 不关心返回值结果 但是也是需要占位
		t.Log(e)
	}
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr3Sec := arr3[:3] // 不支持负索引
	t.Log(arr3Sec)
	arr3Sec = arr3[3:]
	t.Log(arr3Sec)
	arr3Sec = arr3[:]
	t.Log(arr3Sec)
}
