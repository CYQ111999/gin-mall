package model

// ProductReq 商品请求参数结构体
type ProductReq struct {
	Name        string  `json:"name" binding:"required"`  // 商品名称
	Price       float64 `json:"price" binding:"required"` // 商品价格 前端返回的是单位为元的浮点数，在api层需要转化为分单位的整数
	Stock       int64   `json:"stock" binding:"required"` //库存数量
	Description string  `json:"description"`              //商品详情
}
