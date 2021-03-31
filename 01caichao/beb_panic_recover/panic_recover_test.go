package beb_panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

/**
 * @author  wjj
 * @date  2020/8/30 2:27 下午
 * @description panic  recover
 */

func TestPanicVxExit(t *testing.T) {
	/*defer func() {
		fmt.Println("Finally")
	}()*/
	// 此处的defer 函数recover 等于是手动catch住了错误信息
	defer func() {
		if err := recover(); err != nil {
			// 恢复错误
			fmt.Println("recovered from ", err)
			// log.Error("recovered panic",err)  正常开发
		}
	}()
	fmt.Println("Start")
	// panic 如果不抓住会打印堆栈信息
	panic(errors.New("Something wrong!"))
	//os.Exit(-1)
	// fmt.Println("End")
}
