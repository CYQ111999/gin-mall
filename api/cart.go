package api

import (
	"gin-mall/model"
	"gin-mall/pkg/e"
	"gin-mall/pkg/util"
	dbmodel "gin-mall/repository/db/model"
	"gin-mall/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateCartHandler 创建购物车函数
func CreateCartHandler(c *gin.Context) {
	var req model.CartReq
	// 1.参数校验
	if err := c.ShouldBindJSON(&req); err != nil {
		// 参数校验错误
		util.ParamError(c)
		return
	}
	// 校验购买商品数是否大于等于1
	if req.Quantity < 1 {
		// 返回自定义错误购物数量不能小于1
		util.HandleError(c, e.NewError(e.ERROR_QUANTITY_WRONG))
		return
	}
	// 2.业务逻辑
	// 创建数据库内购物车的结构体，赋值后传给service层
	var cart dbmodel.Cart
	// 要从中间件中获取UserID
	userid, exists := c.Get("user_id")
	if !exists {
		// 如果exists是false说明用户是未登录状态,返回用户未登录的响应
		util.HandleError(c, e.NewError(e.ERROR_NOT_LOGIC))
		return
	}
	//如果已登录就继续
	// 判断商品id是否存在
	_, err := service.GetProductById(req.ProductID)
	if err != nil {
		// 如果不存在接收返回的商品不存在的自定义信息并返回响应
		util.HandleError(c, err)
		return
	}
	// 如果商品存在就赋值后传给service层
	cart.UserID = userid.(int64)
	cart.ProductID = req.ProductID
	cart.Quantity = req.Quantity
	err = service.CreateCart(&cart)
	// 创建失败返回错误响应
	if err != nil {
		util.HandleError(c, err)
		return
	}
	// 3.返回响应
	util.Success(c, "购物车添加成功", nil)
}

// DeleteCartHandler 删除购物车函数
func DeleteCartHandler(c *gin.Context) {
	// 1.参数校验
	// 从路由获取要删除的购物车id
	idStr := c.Param("cart_id")
	cartId, err := strconv.ParseInt(idStr, 10, 64)
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
	// 2.业务逻辑
	// 传递到service层
	err = service.DeleteCart(cartId, userId.(int64))
	if err != nil {
		util.HandleError(c, err)
		return
	}
	// 3.返回响应
	util.Success(c, "购物车删除成功", nil)
}

// GetCartListHandler 获取购物车列表（含分页）
func GetCartListHandler(c *gin.Context) {
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
	cartList, total, err := service.GetCartList(userId.(int64), req.Page, req.PageSize)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	// 3.返回响应
	util.Success(c, "购物车列表查询成功", gin.H{
		"list":  cartList,
		"total": total,
	})
}

// GetCartByIdHandler 通过id获取购物车详情
func GetCartByIdHandler(c *gin.Context) {
	// 1.参数校验
	// 动态路由获取购物车id
	idStr := c.Param("cart_id")
	// 类型转化
	cartId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
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
	// 将id传给service层并接收返回信息
	cart, err := service.GetCartById(userId.(int64), cartId)
	if err != nil {
		// 如果出错返回错误响应
		util.HandleError(c, err)
		return
	}
	// 3.返回响应
	util.Success(c, "购物车详情查询成功", cart)
}
