package bba_interface

import (
	"io"
	"testing"
)

/**
 * @author  wjj
 * @date  2020/8/29 11:25 下午
 * @description 接口  Duck typeing
 * “当看到一只鸟走起来像鸭子、游泳起来像鸭子、叫起来也像鸭子，那么这只鸟就可以被称为鸭子。”
 * 我们并不关心对象是什么类型，到底是不是鸭子，只关心行为。
	类似python中
	#coding=utf-8
	class Duck:
		def quack(self):
			print "Quaaaaaack!"

	class Bird:
		def quack(self):
			print "bird imitate duck."

	class Doge:
		def quack(self):
			print "doge imitate duck."

	def in_the_forest(duck):
		duck.quack()

	duck = Duck()
	bird = Bird()
	doge = Doge()
	for x in [duck, bird, doge]:
		in_the_forest(x)
*/

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

// 只要方法签名保持一致 就认为是接口的实现
func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}

// 大接口可以由多个较小接口组合而成
type ReadWriter interface {
	io.Reader
	io.Writer
}
