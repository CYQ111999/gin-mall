package model

import "time"

// Cart 购物车结构体
type Cart struct {
	CartID int64 `gorm:"column:cart_id;primary_key;autoIncrement"` // 购物车id
	// 用户和商品组成联合索引，这样一个用户买同一个商品只会在同一个购物车上累加商品数而不是多一个购物车
	UserID    int64     `gorm:"column:user_id;uniqueIndex:idx_user_product"`    // 用户id
	ProductID int64     `gorm:"column:product_id;uniqueIndex:idx_user_product"` // 商品id
	Quantity  int64     `gorm:"column:quantity;default:1"`                      // 购物数量  默认1
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`               // 购物车创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`               // 购物车更新时间
}

// CartList 列表返回购物车的临时结构体包含购物车内商品的简略信息
type CartList struct {
	CartID       int64  `json:"cart_id"`       // 购物车id
	UserID       int64  `json:"user_id"`       // 用户id
	ProductID    int64  `json:"product_id"`    // 商品id
	Quantity     int64  `json:"quantity"`      // 购物数量
	ProductName  string `json:"product_name"`  // 商品名
	ProductPrice int64  `json:"product_price"` // 商品价格
}

// CartDetail 返回购物车时的临时结构体包含购物车和商品的详细信息
type CartDetail struct {
	CartID       int64     `json:"cart_id"`       // 购物车id
	UserID       int64     `json:"user_id"`       // 用户id
	ProductID    int64     `json:"product_id"`    // 商品id
	Quantity     int64     `json:"quantity"`      // 购物数量
	ProductName  string    `json:"product_name"`  // 商品名
	ProductPrice int64     `json:"product_price"` // 商品价格
	CreatedAt    time.Time `json:"created_at"`    // 购物车创建时间
	UpdatedAt    time.Time `json:"updated_at"`    // 购物车更新时间
}
