package main

// sql 注入
// 防止sql注入  预编译问题
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

func sqlInjectDemo(name string) {
	// 自己拼接SQL语句的字符串
	sqlStr := fmt.Sprintf("select id, name, age from user where name = '%s'", name)
	fmt.Printf("SQL:%s\n", sqlStr)
	var users []user
	err := db.Select(&users, sqlStr)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	for _, u := range users {
		fmt.Printf("user:%#v\n", u)
	}
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
	// SQL 注入的几种示例
	sqlInjectDemo("天说")
	sqlInjectDemo("xxx' or 1 = 1 #") // SQL:select id, name, age from user where name = 'xxx' or 1 = 1 #'
	// sqlInjectDemo("xxx' union select * from user #") // SQL:select id, name, age from user where name = 'xxx' union select * from user #'
	sqlInjectDemo("xxx' union select  id, name, age from user #") // SQL:select id, name, age from user where name = 'xxx' union select  id, name, age from user #'
	sqlInjectDemo("xxx' and (select count(*) from user) < 10 #")  // SQL:select id, name, age from user where name = 'xxx' and (select count(*) from user) < 10 #'
}
