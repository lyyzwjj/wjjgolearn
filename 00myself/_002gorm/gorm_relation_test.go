package _002gorm

import (
	"gorm.io/gorm"
	"testing"
)

// Belongs To
// `User` 属于 `Company`，`CompanyID` 是外键
type User struct {
	gorm.Model
	Name      string
	CompanyID int
	Company   Company
}

type Company struct {
	ID   int
	Name string
}

// 重写外键
//type User struct {
//	gorm.Model
//	Name         string
//	CompanyRefer int
//	Company      Company `gorm:"foreignKey:CompanyRefer"`
//	// 使用 CompanyRefer 作为外键
//}

// 重写引用
//type User struct {
//	gorm.Model
//	Name      string
//	CompanyID string
//	Company   Company `gorm:"references:Code"` // 使用 Code 作为引用
//}
//
//type Company struct {
//	ID   int
//	Code string
//	Name string
//}

func TestGormBelongsTo(t *testing.T) {
	Init()
	// 迁移 schema
	db.AutoMigrate(&User{})
}
