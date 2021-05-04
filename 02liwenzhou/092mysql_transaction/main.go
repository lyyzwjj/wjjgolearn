package main

// go 连接mysql示例

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 自动执行init方法
)

/*
查看mysql的 driver.go源码 有把mysql注册上去的方法
func init() {
	sql.Register("mysql", &MySQLDriver{})
}
*/

type user struct {
	id   int
	name string
	age  int
}

var db *sql.DB // 一个连接池对象

func initDB() (err error) {
	// 数据库信息 DSN Data Source Name
	dsn := "root:Wzzst310@163.com@(wjjzst.com:3306)/test"
	// 连接数据库
	db, err = sql.Open("mysql", dsn) // 不会校验用户名和密码是否正确 只会校验dsn格式是否正确
	if err != nil {                  // dsn格式不正确的时候会报错
		// fmt.Printf("dsn:%s invalid, err:%v\n", dsn, err)
		return
	}
	err = db.Ping()
	if err != nil {
		// fmt.Printf("open %s failed, err:%v\n", dsn, err)
		return
	}
	// 设置数据库连接池的最大连接数
	// db.SetMaxOpenConns(1)
	db.SetMaxOpenConns(10)
	// 设置最大空闲连接数
	db.SetMaxIdleConns(5)
	return
}

func transactionDemo() {
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("begin failed,err:%v\n", err)
		return
	}
	// 执行多个SQL操作
	sqlStr1 := "update user set age=age-2 where id=2"
	// sqlStr2 := "update user set age=age+2 where id=3"
	sqlStr2 := "update users set age=age+2 where id=3"
	// 执行SQL1
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		// 要回滚
		tx.Rollback()
		fmt.Println("执行SQL1出错啦, 要回滚! ")
		return
	}
	// 执行SQL2
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		// 要回滚
		tx.Rollback()
		fmt.Println("执行SQL2出错啦, 要回滚! ")
		return
	}
	// 上面两步SQL都执行成功,就提交本次事务
	err = tx.Commit()
	if err != nil {
		// 要回滚
		tx.Rollback()
		fmt.Println("提交出错啦, 要回滚! ")
		return
	}
	fmt.Println("事务执行成功")

}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
		return
	}
	fmt.Println("连接数据库成功")
	transactionDemo()
}
