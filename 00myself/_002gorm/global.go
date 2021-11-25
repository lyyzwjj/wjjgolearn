package _002gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() {
	dsn := "root:Wzzst310@163.com@tcp(wjjzst.com:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	var err error
	if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		panic("failed to connect database")
	}
}
