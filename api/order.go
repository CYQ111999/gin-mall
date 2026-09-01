package api

import (
	"gin-mall/model"
	"gin-mall/pkg/e"
	"gin-mall/pkg/util"
	"gin-mall/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateOrderHandler 创建订单信息函数
func CreateOrderHandler(c *gin.Context) {
	var req model.OrderCreateReq
	// 1. 参数校验
	if err := c.ShouldBindJSON(&req); err != nil {
		util.ParamError(c)
		return
	}
	// 从中间件获取当前登录用户ID
	userIdVal, exists := c.Get("user_id")
	if !exists {
		util.HandleError(c, e.NewError(e.ERROR_NOT_LOGIC))
		return
	}
	userId := userIdVal.(int64)
	// 2.调用Service层创建订单
	orderSN, err := service.CreateOrder(userId, req)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	// 3.返回成功响应
	util.Success(c, "订单创建成功", gin.H{
		"order_sn": orderSN,
	})
}

// GetOrderListHandler 获取订单列表（含分页）
func GetOrderListHandler(c *gin.Context) {
	// 建立分页请求参数的变量，包含当前页数和分页大小
	var req model.ListReq
	// 1.参数校验
	// 直接从路由获取page和pagesize
	if err := c.ShouldBindQuery(&req); err != nil {
		// 出错则返回请求参数错误的自定义响应
		util.ParamError(c)
		return
	}
	// 要从中间件中获取userId
	userId, exists := c.Get("user_id")
	if !exists {
		// 如果exists是false说明用户是未登录状态,返回用户未登录的响应
		util.HandleError(c, e.NewError(e.ERROR_NOT_LOGIC))
		return
	}
	// 2.业务逻辑
	// 将信息传给service层执行并接收返回信息
	orderList, total, err := service.GetOrderList(userId.(int64), req.Page, req.PageSize)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	// 返回响应
	util.Success(c, "订单列表查询成功", gin.H{
		"order_list": orderList,
		"total":      total,
	})
}

// CancelOrderHandler 取消订单函数
func CancelOrderHandler(c *gin.Context) {
	// 1. 参数校验
	// 从动态路由获取订单id
	IdStr := c.Param("order_id")
	orderId, err := strconv.ParseInt(IdStr, 10, 64)
	if err != nil {
		util.ParamError(c)
		return
	}
	// 再从中间件中获取用户id
	userIdVal, exists := c.Get("user_id")
	if !exists {
		util.HandleError(c, e.NewError(e.ERROR_NOT_LOGIC))
		return
	}
	// 2. 业务逻辑
	err = service.CancelOrder(userIdVal.(int64), orderId)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	// 3. 返回响应
	util.Success(c, "订单取消成功", nil)
}

// PayOrderHandler 订单支付函数
func PayOrderHandler(c *gin.Context) {
	// 1. 参数校验
	// 从路由获取要支付的订单id
	idStr := c.Param("order_id")
	orderId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		util.ParamError(c)
		return
	}
	// 从中间件中获取用户id
	userIdVal, exists := c.Get("user_id")
	if !exists {
		util.HandleError(c, e.NewError(e.ERROR_NOT_LOGIC))
		return
	}
	userId := userIdVal.(int64)
	// 业务逻辑
	err = service.PayOrder(userId, orderId)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	// 返回响应
	util.Success(c, "支付成功", nil)
}
