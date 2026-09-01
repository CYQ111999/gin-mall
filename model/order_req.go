package model

import (
	"gin-mall/repository/db/model"
	"time"
)

// OrderCreateReq 订单创建请求参数结构体
type OrderCreateReq struct {
	AddressID   int64   `json:"address_id" binding:"required"`    // 收货地址ID
	CartItemIDs []int64 `json:"cart_item_ids" binding:"required"` // 选中的购物车记录ID切片
}

// OrderResponse 订单主表返回响应的结构体
type OrderResponse struct {
	OrderSN         string                `json:"order_sn"`         // 业务订单号
	TotalAmount     int64                 `json:"total_amount"`     // 总金额
	Status          int                   `json:"status"`           // 支付状态
	CreateAt        time.Time             `json:"create_at"`        // 创建时间
	AddressSnapshot model.AddressSnapshot `json:"address_snapshot"` // 地址快照
	Items           []OrderItemResponse   `json:"items"`            // 附表连接
}

// OrderItemResponse 订单附表返回响应的结构体
type OrderItemResponse struct {
	ProductID   int64  `json:"product_id"`   // 商品id
	ProductName string `json:"product_name"` // 商品名
	Quantity    int32  `json:"quantity"`     // 商品数量
	Price       int64  `json:"price"`        // 下单时的商品单价
}
