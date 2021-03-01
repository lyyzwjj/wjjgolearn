package afb_condition

import (
	"fmt"
	"runtime"
	"testing"
)

/**
 * @author  wjj
 * @date  2020/8/21 1:53 上午
 * @description
 */

func TestIfMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("1==1")
	}
	// 执行一个方法有多个返回值 如果有err则执行 没有err则执行另外一个
	//if v, err := someFun(); err == nil {
	//
	//} else {
	//
	//}
}
func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}
}

// if else 可以用case写
func TestSwitchMultiCondtion(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("unknow")
		}
	}
}
