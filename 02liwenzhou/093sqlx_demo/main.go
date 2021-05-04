package main

// go sqlx demo
// 下载第三方库
// go get -u github.com/jmoiron/sqlx
// 不同的库占位符不同

// MySQL			?
// PostgreSQL		$1,$2等
// SQLite			?和$1
// Oracle 			:name
import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 自动执行init方法
	"github.com/jmoiron/sqlx"
)

type user struct {
	ID   int
	Name string
	Age  int
}

var db *sqlx.DB // 一个连接池对象
func initDB() (err error) {
	// 数据库信息 DSN Data Source Name
	dsn := "root:Wzzst310@163.com@(wjjzst.com:3306)/test"
	// 连接数据库
	// db, err = sql.Open("mysql", dsn) // 不会校验用户名和密码是否正确 只会校验dsn格式是否正确
	db, err = sqlx.Connect("mysql", dsn) // 连接的是否并且ping一下
	if err != nil {                      // dsn格式不正确的时候会报错
		// fmt.Printf("dsn:%s invalid, err:%v\n", dsn, err)
		return
	}
	//err = db.Ping()
	//if err != nil {
	//	// fmt.Printf("open %s failed, err:%v\n", dsn, err)
	//	return
	//}
	// 设置数据库连接池的最大连接数
	// db.SetMaxOpenConns(1)
	db.SetMaxOpenConns(10)
	// 设置最大空闲连接数
	db.SetMaxIdleConns(5)
	return
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
		return
	}
	fmt.Println("连接数据库成功")
	sqlStr1 := "select id, name, age from user where id = 2"
	var u user
	db.Get(&u, sqlStr1)
	fmt.Printf("u:%#v\n", u)
	// var userList = make([]user, 0, 10) // slice 本来就是引用
	var userList []user // slice 本来就是引用
	sqlStr2 := "select id, name, age from user"
	err = db.Select(&userList, sqlStr2)
	if err != nil {
		fmt.Printf("select failed, err:%v\n", err)
	}
	fmt.Printf("userList:%#v\n", userList)
}
