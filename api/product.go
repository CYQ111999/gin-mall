package api

import (
	"gin-mall/model"
	"gin-mall/pkg/util"
	dbmodel "gin-mall/repository/db/model"
	"gin-mall/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateProductHandler 创建商品
func CreateProductHandler(c *gin.Context) {
	var req model.ProductReq
	// 1.参数校验
	if err := c.ShouldBindJSON(&req); err != nil {
		// 参数校验错误
		util.ParamError(c)
		return
	}
	// 2.业务处理
	// 创建数据库结构体变量，将参数结构体变量数据处理后赋值
	var product dbmodel.Product
	product.Name = req.Name
	product.Price = (int64)(req.Price * 100) //需要将元单位浮点数转化成分单位整数
	product.Stock = req.Stock
	product.Description = req.Description
	// 将赋值后的结构体变量传给service层
	err := service.CreateProduct(&product)
	// 创建失败返回错误响应
	if err != nil {
		util.HandleError(c, err)
		return
	}
	// 3.返回响应
	util.Success(c, "商品添加成功", nil)
}

// GetProductListHandler 获取商品列表(含分页)
func GetProductListHandler(c *gin.Context) {
	// 建立分页请求参数的变量，包含当前页数和分页大小
	var req model.ListReq
	// 1.参数校验
	// 直接从路由获取page和pagesize
	if err := c.ShouldBindQuery(&req); err != nil {
		// 出错则返回请求参数错误的自定义响应
		util.ParamError(c)
		return
	}
	// 2.业务处理
	// 将前端传来的分页数据传给service层,从service层接收当前页数据，商品总数，和错误信息
	productList, total, err := service.GetProductList(req.Page, req.PageSize)
	if err != nil {
		// 查询失败返回错误
		util.HandleError(c, err)
		return
	}
	// 3.返回响应
	util.Success(c, "商品列表查询成功", gin.H{
		"list":  productList,
		"total": total,
	})
}

// GetProductByIdHandler 根据商品id获取商品信息
func GetProductByIdHandler(c *gin.Context) {
	// 1.参数校验
	// 动态路由获取商品id
	idStr := c.Param("productid")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		util.ParamError(c)
		return
	}
	// 2.业务处理
	// 将从前端解析的商品id传给service层并接收返回的商品详情和错误信息
	product, err := service.GetProductById(id)
	if err != nil {
		// 查询失败返回错误信息
		util.HandleError(c, err)
		return
	}
	// 3.返回响应
	util.Success(c, "商品详情查询成功", product)
}
