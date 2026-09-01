package api

import (
	"gin-mall/model"
	"gin-mall/pkg/util"
	"gin-mall/service"

	"github.com/gin-gonic/gin"
)

// SignupHandler 用户注册
func SignupHandler(c *gin.Context) {
	// 1.参数校验
	// 创建用户请求参数结构体
	var user model.UserSignupReq
	// 从前端获取注册用户信息
	if err := c.ShouldBindJSON(&user); err != nil {
		util.ParamError(c)
		return
	}
	// 2. 业务处理
	err := service.Signup(user.Username, user.Password, user.Email)
	// 返回错误
	if err != nil {
		util.HandleError(c, err)
		return
	}
	// 3. 返回响应
	util.Success(c, "注册成功", nil)
}

// LogicHandler 用户登录
func LogicHandler(c *gin.Context) {
	// 1.参数校验
	var user model.UserLogicReq
	// 从前端获取登录用户信息
	if err := c.ShouldBind(&user); err != nil {
		util.ParamError(c)
		return
	}
	// 2.业务处理
	token, err := service.Logic(user.Username, user.Password)
	// 返回错误
	if err != nil {
		util.HandleError(c, err)
		return
	}
	// 3.返回响应
	util.Success(c, "登录成功", gin.H{"token": token})
}

// UpdateUserHandler 修改用户信息
func UpdateUserHandler(c *gin.Context) {
	// 1. 参数校验
	var user model.UserSignupReq
	if err := c.ShouldBindJSON(&user); err != nil {
		util.ParamError(c)
		return
	}
	// 2. 业务处理
	err := service.UpdateUser(user.Username, user.Password, user.Email)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	// 3. 返回响应
	util.Success(c, "修改成功", nil)
}
