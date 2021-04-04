package main

import "fmt"

// fmt
func main() {
	fmt.Print("沙河")
	fmt.Print("娜扎")
	fmt.Println("----------")
	fmt.Println("沙河")
	fmt.Println("娜扎")
	// fmt.Printf("格式化字符串",值)
	// %T :查看类型
	// %d :十进制数
	// %b :二进制数
	// %o :八进制数
	// %x :十六进制数
	// %c :字符
	// %s :字符串
	// %v :值
	// %f :浮点数

	// 通用占位符
	// %v 值得默认格式表示
	// %+v 类似%v.但输出结构体时会添加字段名
	// %#v 值得Go语法表示
	// %T 打印值的类型
	// %% 百分号
	var m1 = make(map[string]int, 1)
	m1["理想"] = 100
	fmt.Printf("%v\n", m1)  // map[理想:100]
	fmt.Printf("%#v\n", m1) // map[string]int{"理想":100}
	printPercentage(99)

	// fmt.Printf("%s\n", 100)  打印成string类型的  不支持
	fmt.Printf("%v\n", 99) // 不知道类型 用%v准没错

	// 布尔型
	// %t true 或false

	// 整形
	// %b :二进制数
	// %c :改值对应的unicode码值
	// %d :十进制数
	// %o :八进制数
	// %x :十六进制数	使用a-f
	// %X :十六进制数 使用A-F
	// %U 表示为Unicode格式: U+1234,等价于  U+0041
	// %q 该值对应的单括号括起来的go语法字符字面值,必要时会采用安全的转义表示		'A'
	fmt.Println("整形")
	n := 65
	fmt.Printf("%b\n", n)
	fmt.Printf("%c\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	fmt.Printf("%X\n", n)
	fmt.Printf("%U\n", n)
	fmt.Printf("%q\n", n)

	// 浮点数和复数
	// %b 无小数部分\二进制指数的科学计数法,如8695310906796750p-46
	// %e 科学技术法 如1.235678e+02
	// %E 科学技术法 如1.235678E+02
	// %f 有小数部分但无指数部分,如123.567800
	// %F 等于与%f			,如123.567800
	// %g 根据实际情况使用%e或%f格式(以获得更简介\准确的输出) 如: 123.5678
	// %G 根据实际情况使用%E或%F格式(以获得更简介\准确的输出) 如: 123.5678
	f := 123.5678
	fmt.Printf("%b\n", f)
	fmt.Printf("%e\n", f)
	fmt.Printf("%E\n", f)
	fmt.Printf("%f\n", f)
	fmt.Printf("%F\n", f)
	fmt.Printf("%g\n", f)
	fmt.Printf("%G\n", f)

	// 字符串和[]byte
	// %s 直接输出字符串或者[]byte
	// %q 该值对应的双括号括起来的go语法字符串字面值,必要时会采用安全的转义表示
	// %x 每个字节用两字符十六进制数表示(使用a-f)
	// %X 每个字节用两字符十六进制数表示(使用A-F)
	s := "小王子"
	fmt.Printf("%s\n", s) // 小王子
	fmt.Printf("%q\n", s) // "小王子"
	fmt.Printf("%x\n", s) // e5b08fe78e8be5ad90	三个中文六个字节
	fmt.Printf("%X\n", s) // E5B08FE78E8BE5AD90
	b := [...]byte{61, 62, 63, 64, 65, 66, 65}
	fmt.Printf("%s\n", b) // =>?@ABA
	fmt.Printf("%q\n", b) // "=>?@ABA"
	fmt.Printf("%x\n", b) // 3d3e3f40414241
	fmt.Printf("%X\n", b) // 3D3E3F40414241

	// 指针	address of 0th element in base 16 notation, with leading 0x
	// %p
	fmt.Println("指针")
	fmt.Printf("%p\n", s)  // %!p(string=小王子)		想要取指针一定要 %p 跟&
	fmt.Printf("%p\n", &s) // 0xc000010220
	fmt.Println(&s)        // 0xc000010220

	in := 1
	fmt.Printf("%p\n", in)  // %!p(string=小王子)
	fmt.Printf("%p\n", &in) // 0xc000010220
	fmt.Println(&in)        // 0xc000010220

	// 宽度标识符		用于对齐
	// %f 默认宽度, 默认精度
	// %9f 宽度9, 默认精度
	// %.2f 默认宽度, 精度2
	// %9.2f 宽度9, 精度2
	// %9.f 宽度9, 精度0

	k := 12.34
	fmt.Printf("%f\n", k)    //12.340000
	fmt.Printf("%9f\n", k)   //12.340000
	fmt.Printf("%.2f\n", k)  //12.34
	fmt.Printf("%9.2f\n", k) //    12.34
	fmt.Printf("%9.f\n", k)  //       12

	m := 1234567890.34
	fmt.Printf("%9.f\n", m)  // 1234567890
	fmt.Printf("%9.2f\n", m) // 1234567890.34 大于了会全部显示

	// 其他
	// %+ 总是输出值得正负号; 对%q(%+q)会生成全部是ASCII字符的输出(通过转义);
	// % 对数值,正数前面加空格而复数前面加负号; 对字符串采用%x 或 %X 时(%x或%X)会给各打印的字节之间加空格;
	// %- 在输出右边填充空白而不是默认的左边(即从默认的右对齐切换为左对齐);
	// %# 八进制数前面加o(%#o),十六进制数前面加ox(%#x)或0X(%#X),指针去掉前面的ox(%#p) 对%q(%#q) %U(%#U)会输出空格和单引号括起来的go字面值;
	// %0 使用0而不是空格填充, 对于数值类型会把填充的0放在正负号后面
	fmt.Println("其他")
	s1 := "小王子"
	fmt.Printf("%s\n", s1)     //小王子
	fmt.Printf("%5s\n", s1)    //右对齐 总长度5 左边空两格	|  小王子
	fmt.Printf("%-5s\n", s1)   //左对齐 总长度5 右边空两格	|小王子
	fmt.Printf("%5.7s\n", s1)  //右对齐 总长度5 保留7个	    |  小王子
	fmt.Printf("%-5.7s\n", s1) //左对齐 总长度5 保留7个		|小王子
	s2 := "小王子啊啊啊啊啊啊"
	fmt.Printf("%5.7s\n", s2)  //右对齐 总长度5 保留7个	|小王子啊啊啊啊
	fmt.Printf("%-5.7s\n", s2) //左对齐 总长度5 保留7个	|小王子啊啊啊啊
	fmt.Printf("%5.2s\n", s1)  //右对齐 总共为5 只留2个  |   小王
	fmt.Printf("%05s\n", s1)   //右对齐 总共为5 其他臫0填充	|00小王子

	// Sprint Sprint系列函数会把传入的数据生成并返回一个字符串
	// fmt.Sprint
	// fmt.Sprintf
	// fmt.Sprintln

	s4 := fmt.Sprint("沙河小王子")
	name2 := "沙河小王子"
	age2 := 18
	s5 := fmt.Sprintf("name:%s,age:%d", name2, age2)
	s6 := fmt.Sprint("沙河小王子")
	fmt.Println(s4, s5, s6)

	// Errorf
	// Errof 函数根据format参数生成格式化字符串并返回一个包含包含该字符串的错误.

	// 获取输入 三个函数
	fmt.Println("获取输入")
	// fmt.Scan
	// fmt.Scanf
	// fmt.Scanln
	// Scan从标准输入扫描文本,读取由空白符分隔的值保存到传递给本函数的参数中,换行符视为空白符
	// 本函数返回成功扫描的数据个数和遇到的任何错误. 如果读取的数据个数比提供的参数少,会返回一个错误报告原因

	var pa string
	fmt.Scan(&pa)
	fmt.Println("用户输入的内容是:", pa)

	var (
		name   string
		name1  string
		age    int
		age1   int
		class  string
		class1 string
	)
	// fmt.Scanf("%s", &name)
	fmt.Scanf("%s %d %s\n", &name, &age, &class) // 不能换行输入
	fmt.Println(name, age, class)

	fmt.Scanln(&name1, &age1, &class1) // 这样只能一次加载内容 只填一部分 其他就默认
	//  scanln函数会识别空格左右的内容，但是一旦遇到换行符就会立即结束，不论后续还是否存在需要带输入的内容。
	//  这种写法的话必须把name和age一行输入，因为scanln是以回车为标志结束
	fmt.Println(name1, age1, class1)
	fmt.Scanln(&name1)
	fmt.Scanln(&age1)
	fmt.Scanln(&class1) // 可以分多次加载进内容
	fmt.Println(name1, age1, class1)

}
func printPercentage(num int) {
	fmt.Printf("%d\n", num)
	fmt.Printf("%d%%\n", num) // %% 转义 表示真正的百分号
}
