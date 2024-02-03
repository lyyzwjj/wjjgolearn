package gitee

import "fmt"

const TestParam = "TestParam"

func FunTest() {
	println("FunTest")
}
func FunParamsTest[T any](t T) {
	println("FunParamsTest")
	println(t)
}

func init() {
	fmt.Println("import gitee_private_test_demo.go 我自动执行...")
}
