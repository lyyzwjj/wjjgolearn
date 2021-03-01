package bea_error

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

/**
 * @author  wjj
 * @date  2020/8/30 11:37 上午
 * @description 异常错误
 */

// 自定义异常
//type error interface {
//	Error() string
//}
var LessThanTwoError = errors.New("n should be not less than 2")
var LargerThanHundredError = errors.New("n should be not larger than 100")

func GetFibonacci(n int) ([]int, error) {
	/*if n < 0 || n > 100 {
		return nil, errors.New("n should be in [2,100]")
	}*/
	/*if n < 0 {
		return nil, errors.New("n should be not less than 2")
	}

	if n > 100 {
		return nil, errors.New("n should be not larger than 100")
	}*/
	if n < 0 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargerThanHundredError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	// t.Log(GetFibonacci(-100000))
	/*
		v, err := GetFibonacci(-10);
		if  err != nil {
			t.Error(err)
		} else {
			t.Log(v)
		}
	*/
	// 可以合并成一行
	if v, err := GetFibonacci(100); err != nil {
		if err == LessThanTwoError {
			fmt.Println("It is less.")
		}
		t.Error(err)
	} else {
		t.Log(v)
	}
}

//  不推荐
func GetFibonacci1(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err == nil {
		if list, err = GetFibonacci(i); err == nil {
			fmt.Println(list)
		} else {
			fmt.Println("Error", err)
		}
	} else {
		fmt.Println("Error", err)
	}
}

// 2相对于1 就没有那么多嵌套了
// 对于error 反方向操作  如果有error 则输入错误信息
func GetFibonacci2(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("Error", err)
		return
	}
	if list, err = GetFibonacci(i); err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(list)
}
