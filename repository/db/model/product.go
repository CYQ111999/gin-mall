package model

// Product 定义商品结构体
type Product struct {
	Id          int64  `gorm:"column:id;primary_key;comment:'商品id'"`
	Name        string `gorm:"column:name;comment:'商品名称'"`
	Price       int64  `gorm:"column:price;comment:'价格(单位为分)'"` // 为防止精度问题采用int和单位分，在
	Stock       int64  `gorm:"colum:stock;comment:'库存数量'"`
	Description string `gorm:"colum:description;comment:'商品详情'"`
}
