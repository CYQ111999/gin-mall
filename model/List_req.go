package model

// ListReq 分页请求参数结构体
type ListReq struct {
	Page     int `form:"page" binding:"required"`      // 当前分页
	PageSize int `form:"page_size" binding:"required"` // 每页信息大小
}
