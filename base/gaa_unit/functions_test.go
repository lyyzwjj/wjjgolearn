package gaa_unit

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * @author  wjj
 * @date  2020/9/1 11:59 下午
 * @description  test assert
 */

func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := squares(inputs[i])
		if ret != expected[i] {
			// t.Errorf("input is %d,the expected is %d,the actual %d", inputs[i], expected[i], ret)  //本次测试 本测试不会终止 其他测试也能执行
			t.Fatalf("input is %d,the expected is %d,the actual %d", inputs[i], expected[i], ret) //本次测试 本测试终止 其他测试也能执行
		}
	}
}

func TestErrorInCode(t *testing.T) {
	fmt.Println("Start")
	t.Error("Error")
	fmt.Println("End")
}

func TestFatalInCode(t *testing.T) {
	fmt.Println("Start")
	t.Fatal("Fatal")
	fmt.Println("End")
}
func TestSquareAssert(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := squares(inputs[i])
		assert.Equal(t, expected[i], ret)
	}
}
