package model

import (
	"database/sql/driver"
	"errors"
	"time"

	"github.com/goccy/go-json"
)

// Order 订单主表结构体
type Order struct {
	OrderID         int64           `gorm:"column:order_id;primary_key;auto_increment"` // 订单主表id
	OrderSN         string          `gorm:"column:order_sn;"`                           // 业务订单号
	UserID          int64           `gorm:"column:user_id;"`                            // 用户id
	AddressSnapshot AddressSnapshot `gorm:"column:address_snapshot;"`                   // 地址快照
	TotalAmount     int64           `gorm:"column:total_amount;"`                       // 总金额（分）
	Status          int             `gorm:"column:status;"`                             // 支付状态：1待支付，2已支付，3已取消，4已完成
	CreateAt        time.Time       `gorm:"column:create_at;"`                          // 创建时间
	UpdateAt        time.Time       `gorm:"column:update_at;"`                          // 修改时间
	PayAt           time.Time       `gorm:"column:pay_at"`                              // 支付时间
	Items           []OrderItem     `gorm:"foreignKey:OrderID" json:"-"`                // 通过主表id和附表建立连接
}

// OrderItem 订单附表结构体（订单商品细明）
type OrderItem struct {
	ItemID       int64     `gorm:"column:item_id;primary_key;auto_increment"` // 订单附表id
	OrderID      int64     `gorm:"column:order_id"`                           // 订单主表id
	ProductID    int64     `gorm:"column:product_id"`                         // 商品id
	ProductName  string    `gorm:"column:product_name"`                       // 商品名
	Quantity     int64     `gorm:"column:quantity"`                           // 商品数量
	PriceAtOrder int64     `gorm:"column:price_at_order"`                     // 下单时的商品单价（分）
	CreateAt     time.Time `gorm:"column:create_at;"`                         // 创建时间
}

// AddressSnapshot 收获地址快照结构体(存json)
type AddressSnapshot struct {
	Name    string `json:"name"`    // 收货人姓名
	Phone   string `json:"phone"`   // 收货人手机号
	Address string `json:"address"` // 收货地址
}

// Value 实现 driver.Valuer 接口：保存到数据库时，将结构体转为 JSON 字符串
func (a AddressSnapshot) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan 实现 sql.Scanner 接口：从数据库读取时，将 JSON 字符串解析回结构体
func (a *AddressSnapshot) Scan(value interface{}) error {
	if value == nil {
		*a = AddressSnapshot{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("AddressSnapshot Scan 失败: 数据格式不匹配")
	}
	return json.Unmarshal(bytes, a)
}
