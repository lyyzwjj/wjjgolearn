package a_series

import "fmt"

/**
 * @author  wjj
 * @date  2020/8/30 3:33 下午
 * @description
 */
func init() {
	fmt.Println("init1")
}

// 可以定义多个init
func init() {
	fmt.Println("init2")
}

// 大写方法名可以在包外访问
func GetFibonacciSerie(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}

// 小写方法名无法在包外访问
// func Square(n int) int {
func square(n int) int {
	return n * n
}
