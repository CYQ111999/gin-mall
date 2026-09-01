package model

// AddressReq 地址请求参数结构体
type AddressReq struct {
	AddressID int64  `json:"address_id"`
	Address   string `json:"address" binding:"required"` // 收货人地址
	Name      string `json:"name" binding:"required"`    // 收货人名
	Phone     string `json:"phone" binding:"required"`   // 收货人电话
}
