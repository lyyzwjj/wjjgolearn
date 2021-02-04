package bda_polymorphism

import (
	"fmt"
	"testing"
)

/**
 * @author  wjj
 * @date  2020/8/30 1:04 上午
 * @description  多态   就是类似python那样动态类型语言 不管什么类型直接调方法
 */

type Code string
type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

// 只要方法签名保持一致 就认为是接口的实现
func (g *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World\")"
}

type JavaProgrammer struct {
}

// 只要方法签名保持一致 就认为是接口的实现
func (g *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.println(\"Hello World\")"
}
func WriteFirstProgram(p Programmer) {
	// %T 输出实例类型
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}
func TestPolymorphism(t *testing.T) {
	goProg := new(GoProgrammer)
	javaProg := new(JavaProgrammer)
	WriteFirstProgram(goProg)
	WriteFirstProgram(javaProg)
	// prog := GoProgrammer{}    // 这样写不行
	prog := &GoProgrammer{}
	WriteFirstProgram(prog)
}
