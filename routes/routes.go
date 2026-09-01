package routes

import (
	"gin-mall/api"
	"gin-mall/middleware"

	"github.com/gin-gonic/gin"
)

// Routes 路由
func Routes() *gin.Engine {
	r := gin.Default()
	
	// 用户路由分组
	userGroup := r.Group("/user")
	{
		// 用户注册路由
		userGroup.POST("/signup", api.SignupHandler)
		// 用户登录路由
		userGroup.POST("/logic", api.LogicHandler)
		// 用户信息修改路由
		userGroup.PUT("/Update", api.UpdateUserHandler)
	}

	// 商品路由分组
	productGroup := r.Group("/product")
	{
		// 商品创建路由  需身份认证
		productGroup.POST("/create", middleware.JWTAuthMiddleware(), api.CreateProductHandler)
		// 商品列表查询路由
		productGroup.GET("/list", api.GetProductListHandler)
		// 商品详情查询路由
		productGroup.GET("/:productid", api.GetProductByIdHandler)
	}

	// 购物车路由分组  路由均需身份验证
	cartGroup := r.Group("/cart", middleware.JWTAuthMiddleware())
	{
		// 购物车创建路由
		cartGroup.POST("/create", api.CreateCartHandler)
		// 购物车删除路由
		cartGroup.POST("/delete/:cart_id", api.DeleteCartHandler)
		// 购物车列表查询路由
		cartGroup.GET("/list", api.GetCartListHandler)
		// 购物车详情查询路由
		cartGroup.GET("/:cart_id", api.GetCartByIdHandler)
	}

	// 地址路由分组   均需身份验证
	AddressGroup := r.Group("/address", middleware.JWTAuthMiddleware())
	{
		// 地址创建路由
		AddressGroup.POST("/create", api.CreateAddressHandler)
		// 地址删除路由
		AddressGroup.POST("/delete/:address_id", api.DeleteAddressHandler)
		// 地址修改路由
		AddressGroup.POST("/update", api.UpdateAddressHandler)
	}

	// 订单路由分组   均需身份验证
	OrderGroup := r.Group("/order", middleware.JWTAuthMiddleware())
	{
		// 创建订单信息
		OrderGroup.POST("/create", api.CreateOrderHandler)
		// 查询订单分页信息
		OrderGroup.GET("/list", api.GetOrderListHandler)
		// 取消订单
		OrderGroup.POST("/cancel/:order_id", api.CancelOrderHandler)
		// 支付订单
		OrderGroup.POST("/pay/:order_id", api.PayOrderHandler)
	}
	return r
}
