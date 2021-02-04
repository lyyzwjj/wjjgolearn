package bba_interface

import (
	"fmt"
	"testing"
)

/**
 * @author  wjj
 * @date  2020/8/30 1:21 上午
 * @description 空接口和断言  空接口类似Object  可以表示任何类型
 */

func DoSomething(p interface{}) {
	/*
		if i, ok := p.(int); ok {
			fmt.Println("Integer", i)
			return
		}
		if s, ok := p.(string); ok {
			fmt.Println("string", s)
			return
		}
		fmt.Println("Unknown Type")
	*/
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("Unknown Type")
	}
}
func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
}
