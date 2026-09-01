package dao

import (
	"gin-mall/repository/db"
	"gin-mall/repository/db/model"
)

// CreateProduct 创建商品函数
func CreateProduct(product *model.Product) error {
	err := db.DB.Create(product).Error
	// 创建失败返回err
	if err != nil {
		return err
	}
	// 成功返回nil
	return nil
}

// GetProductList 获取商品列表(含分页)
// 根据接收的第几页，每页展示多少商品，返回当前页信息，商品总数(让前端计算共有多少页)，错误信息
func GetProductList(page int, pageSize int) ([]*model.Product, int64, error) {
	// 现在创建切片用来存储当前页信息
	productList := make([]*model.Product, 0)
	// 创建total变量用来存储商品总数
	var total int64
	// 先统计商品总数
	if err := db.DB.Model(&model.Product{}).Count(&total).Error; err != nil {
		// 如果出错说明总数为0返回查询信息为nil,商品总数为0和自定义错误码:商品不存在
		return nil, 0, err
	}
	// Select控制不返回商品详情字段信息,Offset跳过(当前页-1)*每页展示多少数据的数据;
	//Limit从当前位置开始截取pageSize大小的数据，Find存入productList
	err := db.DB.Select("id", "name", "price", "stock").Offset((page - 1) * pageSize).
		Limit(pageSize).Find(&productList).Error
	// 返回当前页信息，商品总数，错误信息
	return productList, total, err
}

// GetProductById 根据商品id返回商品具体信息
func GetProductById(id int64) (*model.Product, error) {
	var product model.Product
	// 根据id查找商品
	err := db.DB.Where("id=?", id).First(&product).Error
	if err != nil {
		// 没查到就返回nil和err
		return nil, err
	}
	// 查到返回商品信息和nil
	return &product, nil
}
