package main

import "fmt"

// 接口示例2
// 不管是什么牌子的车,都能跑
type Car interface {
	run()
}

type Ferrari struct {
	Brand string
}
type Porsche struct {
	Brand string
}

func (f Ferrari) run() {
	fmt.Printf("%s速度70迈~\n", f.Brand)
}
func (p Porsche) run() {
	fmt.Printf("%s速度70迈~\n", p.Brand)
}

func drive(car Car) {
	car.run()
}

func main() {
	var f1 = Ferrari{
		Brand: "法拉利",
	}
	var p1 = Porsche{
		Brand: "保时捷",
	}
	drive(f1)
	drive(p1)
}
