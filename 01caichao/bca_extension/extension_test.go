package bca_extension

import (
	"fmt"
	"testing"
)

/**
 * @author  wjj
 * @date  2020/8/29 11:57 下午
 * @description
 */

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}
func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

/*
type Dog struct {
	p *Pet
}

func (d *Dog) Speak() {
	// d.p.Speak()
	fmt.Println("Wang")
}
func (d *Dog) SpeakTo(host string) {
	// d.p.SpeakTo(host)
	d.Speak()
	fmt.Println(" ", host)
}
*/

type Dog struct {
	Pet
}

func (d *Dog) Speak() {
	// d.p.Speak()  // 没有p了
	fmt.Println("Wang")
}
func TestDog(t *testing.T) {
	// dog := new(Pet)
	// dog := new(Dog)
	// dog.Speak()
	// dog.SpeakTo("Chao")

	// var dog Pet := new(Dog) // 不支持显示类型转换 所以报错

	// var dog *Dog = new(Dog)
	// var p = (*Pet)(dog)   // 无法支持lsp 强制转换也不行 不支持重载
	dog := new(Dog)
	// ...  Chao 内嵌结构类型完全不能当做继承来用的 此处Dog哪怕复写了speak speakTo还是调用了Pet的 speak方法
	dog.SpeakTo("Chao")
}
