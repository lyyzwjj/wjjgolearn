package afa_loop

import "testing"

/**
 * @author  wjj
 * @date  2020/8/21 1:49 上午
 * @description 循环
 */

func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
	// 无限循环
	for {
		t.Log(n)
	}
}
