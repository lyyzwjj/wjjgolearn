package main

// 包的路径从GOPATH/src后面的路径开始写起,路径分隔符用/
// 想被别的包调用的标识符都要首字母大写!
// 单行导入和多行导入
// 导入包的时候可以指定别名
// 导入包不想使用包内部的标识符,需要使用匿名导入
// 每个包导入的时候会自动执行一个名为init()的函数,他没有参数也没有返回值也不能被手动调用
// 多个包都定义了`init()`,则它们的执行顺序见package2图
import (
	"fmt"
	myCalc "wjjgolearn/02liwenzhou/052package/10calc" // 包名字不合法时候  取别名
	// "github.com/wzzst310/wjjgolearn/02liwenzhou/052package/calc"
	"wjjgolearn/02liwenzhou/052package/calc" // 同一个模块就没必要加域名引用
	// _ "github.com/mailru/easyjson"	// 只导包 不用方法	导包会触发调用init()方法  mysql驱动包
	// "github.com/wzzst310/goprojecttest/haha" // 加了github域名的会从github上面找不管本地有没有
)

func main() {
	ret := calc.Add(1, 2)
	fmt.Println(ret)
	ret = myCalc.Sub(2, 1)
	fmt.Println(ret)
}
