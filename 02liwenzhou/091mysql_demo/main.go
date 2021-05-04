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

//查询
func queryOne(id int) {
	// 1. 写查询单条记录的sql语句
	sqlStr := "select id, name, age from user where id=?;"
	// 2. 执行
	// 3. 拿到结果
	var u1 user
	db.QueryRow(sqlStr, id).Scan(&u1.id, &u1.name, &u1.age) // 从连接池拿一个连接出来去数据库查询单条记录
	// fmt.Println(rowObj)
	// rowObj.Scan(&u1.id, &u1.name, &u1.age) // 一定要调用Scan方法释放连接    Scan 会有释放连接的方法    r.rows.Close()
	// 4. 打印结果
	fmt.Printf("u1:%#v\n", u1)

	// rowObj = db.QueryRow(sqlStr, 2) //  db.SetMaxOpenConns(1) 连接池只设置1个时 拿不到就会阻塞
	// fmt.Println("u1:%#v\n", u1)
	/*for i := 0; i < 20; i++ { // 或者这样来验证
		fmt.Printf("开始第%d次查询\n", i)
		db.QueryRow(sqlStr, 1)
	}*/
}

func queryMore(n int) {
	// 1. SQL语句
	sqlStr := "select id, name, age from user where id > ?;"
	// 2. 执行
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Println("exec %s query failed, err:%v\n", err)
		return
	}
	// 3. 一定要关闭rows
	defer rows.Close()
	for rows.Next() {
		var u1 user
		err := rows.Scan(&u1.id, &u1.name, &u1.age)
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
		}
		fmt.Printf("u1:%#v\n", u1)
	}
}

func insert() {
	// 1. 写SQL语句
	sqlStr := `insert into user(name,age) values("图朝阳",28)`
	// 2. exec
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	// 如果是插入数据的操作,能够拿到插入数据的id
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed, err:%v\n", err)
		return
	}
	fmt.Println("id: ", id)
}

func update(newAge int, id int) {
	sqlStr := `update user set age = ? where id > ?`
	ret, err := db.Exec(sqlStr, newAge, id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	// 如果是插入数据的操作,能够拿到插入数据的id
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id failed, err:%v\n", err)
		return
	}
	fmt.Printf("更新了%d行数据\n", n)
}

func delete(id int) {
	sqlStr := `delete from user where id = ?`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	// 如果是插入数据的操作,能够拿到插入数据的id
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id failed, err:%v\n", err)
		return
	}
	fmt.Printf("删除了%d行数据\n", n)
}

// 预处理方式插入多条数据
func prepareInsert() {
	sqlStr := `insert into user(name, age) values(?, ?)`
	stmt, err := db.Prepare(sqlStr) // 把SQL语句先发给我MySQL预处理一下
	if err != nil {
		fmt.Printf("parepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	// 后续只需要拿到stmt去执行一些操作
	var m = map[string]int{
		"六七强": 30,
		"王相机": 32,
		"天说":  72,
		"白慧姐": 40,
	}
	for k, v := range m {
		stmt.Exec(k, v) // 后续只需要传值
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
		return
	}
	fmt.Println("连接数据库成功")
	// queryOne(2)
	// queryMore(0)
	// insert()
	// update(12,1)
	// delete(1)
	// prepareInsert()
	queryMore(0)
}
