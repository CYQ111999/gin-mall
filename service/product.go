package service

import (
	"gin-mall/pkg/e"
	"gin-mall/repository/db/dao"
	"gin-mall/repository/db/model"
)

// CreateProduct 创建商品函数
func CreateProduct(product *model.Product) error {
	// 将api层拿到的商品信息传给dao层
	err := dao.CreateProduct(product)
	// 返回err说明创建失败
	if err != nil {
		return err
	}
	// 创建成功
	return nil
}

// GetProductList 获取商品列表(含分页)
func GetProductList(page int, pageSize int) ([]*model.Product, int64, error) {
	// 从dao层获取信息
	productList, total, err := dao.GetProductList(page, pageSize)
	if err != nil {
		// 出错说明商品信息不存在，返回商品不存在的自定义错误
		return nil, 0, e.NewError(e.ERROR_PRODUCT_NOT_EXIST)
	}
	// 查到了就返回信息
	return productList, total, nil
}

// GetProductById 根据商品id返回商品具体信息
func GetProductById(id int64) (*model.Product, error) {
	// 用商品id从dao层获取商品具体信息
	product, err := dao.GetProductById(id)
	if err != nil {
		// 如果返回err说明没查询到商品，返回商品不存在的自定义信息
		return nil, e.NewError(e.ERROR_PRODUCT_NOT_EXIST)
	}
	// 成功就返回商品信息和nil
	return product, nil
}
