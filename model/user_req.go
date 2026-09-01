package model

// UserSignupReq UserReq 用户注册请求参数结构体
type UserSignupReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

// UserLogicReq 用户登录请求参数结构体
type UserLogicReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
