package model

// CartReq 购物车请求参数结构体
type CartReq struct {
	ProductID int64 `json:"product_id" binding:"required"` // 商品id
	Quantity  int64 `json:"quantity" binding:"required"`   // 商品数量
}
