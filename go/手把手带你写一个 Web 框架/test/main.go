package main

import (
  "gorm.io/gorm"
  "gorm.io/driver/mysql"
  
)

// 定义一个 gorm 类
type UserInfo struct {
   ID           uint
   Name         string
}

func main() {
	// 创建 mysql 连接
	dsn := "root:123456@tcp(192.168.0.101:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	// 插入一条数据
	db.AutoMigrate(&UserInfo{})
	u1 := UserInfo{ID: 2, Name:"angel"}
	db.Create(&u1)

}