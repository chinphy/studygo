package models

import (
	// "database/sql"
	"github.com/jinzhu/gorm"
	// 引入并注册驱动
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// 引入并注册驱动
	// _ "github.com/go-sql-driver/mysql"
)

// DB 数据库对象
var DB *gorm.DB

// InitDB InitDB
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		return nil, err
	}
	DB = db
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	return db, nil
}

// User 对象
type User struct {
	ID       uint64 `gorm:"primary_key;AUTO_INCREMENT"`
	Username string `gorm:"type:varchar(30)"`
	Age      uint8  `gorm:"type:int(11)"`
	Password string `gorm:"type:varchar(32)"`
	Salt     string `gorm:"type:varchar(6)"`
	Status   uint8  `gorm:"type:tinyint"`
}
