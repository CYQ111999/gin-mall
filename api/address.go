package api

import (
	"gin-mall/model"
	"gin-mall/pkg/e"
	"gin-mall/pkg/util"
	db "gin-mall/repository/db/model"
	"gin-mall/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateAddressHandler CreateAddress 创建地址函数
func CreateAddressHandler(c *gin.Context) {
	var req model.AddressReq
	// 1.参数校验
	if err := c.ShouldBindJSON(&req); err != nil {
		// 参数校验错误
		util.ParamError(c)
		return
	}
	// 要从中间件中获取UserID
	userid, exists := c.Get("user_id")
	if !exists {
		// 如果exists是false说明用户是未登录状态,返回用户未登录的响应
		util.HandleError(c, e.NewError(e.ERROR_NOT_LOGIC))
		return
	}
	// 2.业务逻辑
	// 创建地址结构体变量
	var address db.Address
	// 赋值
	address.UserID = userid.(int64)
	address.Address = req.Address
	address.Name = req.Name
	address.Phone = req.Phone
	// 将信息传给service层并接受返回信息
	err := service.CreateAddress(&address)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	// 返回响应
	util.Success(c, "地址创建成功", gin.H{"address": address})
}

// DeleteAddressHandler 地址信息删除函数
func DeleteAddressHandler(c *gin.Context) {
	// 1. 参数校验
	// 直接从路由获取要删除的地址id,并转化类型
	addressIdStr := c.Param("address_id")
	addressId, err := strconv.ParseInt(addressIdStr, 10, 64)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	// 要从中间件中获取userId
	userId, exists := c.Get("user_id")
	if !exists {
		// 如果exists是false说明用户是未登录状态,返回用户未登录的响应
		util.HandleError(c, e.NewError(e.ERROR_NOT_LOGIC))
		return
	}
	// 2. 业务逻辑
	err = service.DeleteAddress(userId.(int64), addressId)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	// 返回响应
	util.Success(c, "地址删除成功", nil)
}

// UpdateAddressHandler 修改地址信息
func UpdateAddressHandler(c *gin.Context) {
	var req model.AddressReq
	// 1. 参数校验
	if err := c.ShouldBindJSON(&req); err != nil {
		util.ParamError(c)
		return
	}
	userId, exists := c.Get("user_id")
	if !exists {
		util.HandleError(c, e.NewError(e.ERROR_NOT_LOGIC))
		return
	}
	// 2.业务逻辑
	var address db.Address
	address.AddressID = req.AddressID
	address.Name = req.Name
	address.Phone = req.Phone
	address.Address = req.Address
	err := service.UpdateAddress(userId.(int64), &address)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	// 3. 返回响应
	util.Success(c, "地址修改成功", nil)
}
