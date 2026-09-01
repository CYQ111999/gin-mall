package db

import (
	"gin-mall/repository/db/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init 初始化表结构
func Init() error {
	dsn := "root:cyq20070115@tcp(127.0.0.1:3306)/mall?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return DB.AutoMigrate(&model.User{}, &model.Product{}, &model.Cart{}, &model.Address{}, &model.Order{},
		&model.OrderItem{})
}
