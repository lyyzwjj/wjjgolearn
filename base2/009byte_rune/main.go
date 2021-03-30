package main

import "fmt"

// byte和rune类型
// Go语言中为了处理非ASCII码类型的字符 中文日文等 定义了新的rune类型 实际是一个int32
func main() {
	s := "Hello沙河" // 5*8 +2* 3*8  一个ASCII码 一个字节 一个中文 3个字节
	// len() 求得是byte字节的数量 5+2*3 = 11
	n := len(s) // 求字符串s的长度,把长度保存到变量n中
	fmt.Println(n)

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
		fmt.Printf("%c\n", s[i])         // %c:字符
		fmt.Printf("%v(%c)", s[i], s[i]) // %c:字符
	}
	for _, c := range s { // 从字符串中拿出具体的字符
		fmt.Println(c)
		fmt.Printf("%v(%c)", c) // %c:字符
	}
	for _, c := range s { // 从字符串中拿出具体的字符
		fmt.Printf("%c\n", c)
	}
	stringChange()
	typeConvert()
}

// 字符串修改
func stringChange() {
	// "Hello" => 'H' 'e' 'l' 'l' 'o'
	s2 := "白萝卜"      // [ 白 萝 卜 ]
	s3 := []rune(s2) // 把字符串强制转换成一个rune切片
	s3[0] = '红'
	fmt.Println(string(s3)) // 把rune切片强制转换成字符串

	c1 := "红"
	c2 := '红' // rune(int32)
	fmt.Printf("c1:%T c2:%T\n", c1, c2)
	c3 := "H" // string
	c4 := 'H' // int32
	c5 := byte('H')
	fmt.Printf("c3:%T c4:%T\n", c3, c4)
	fmt.Printf("c4:%d\n", c4)
	fmt.Printf("c5:%T", c5) // uint8
}

// 类型转换
func typeConvert() {
	n1 := 10
	var f float64
	f = float64(n1)
	fmt.Println(f)
	fmt.Printf("f:%T", f)

}
