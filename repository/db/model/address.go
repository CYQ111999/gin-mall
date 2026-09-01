package model

import "time"

// Address 地址结构体   （注意操作用户和收货人可能不同）
type Address struct {
	AddressID int64     `gorm:"column:address_id;primary_key;AUTO_INCREMENT"` // 地址id
	UserID    int64     `gorm:"column:user_id"`                               // 用户id
	Name      string    `gorm:"column:name"`                                  // 收货人名
	Phone     string    `gorm:"column:phone"`                                 // 收货人电话
	Address   string    `gorm:"column:address"`                               // 收货人地址
	CreatedAt time.Time `gorm:"column:created_at"`                            // 地址创建时间
	UpdatedAt time.Time `gorm:"column:updated_at"`                            // 地址修改时间
}
