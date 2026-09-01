package service

import (
	"gin-mall/pkg/e"
	"gin-mall/repository/db/dao"
	"gin-mall/repository/db/model"
)

// CreateCart 创建购物车函数
func CreateCart(cart *model.Cart) error {
	// 将api层获取的购物车信息传给dao层创建
	err := dao.CreateCart(cart)
	if err != nil {
		return err
	}
	return nil
}

// DeleteCart 删除购物车函数
func DeleteCart(cartId, userId int64) error {
	// 从api层获取的数据传递到dao层删除数据
	err := dao.DeleteCart(cartId, userId)
	if err != nil {
		return err
	}
	return nil
}

// GetCartList 获取购物车列表（含分页）
func GetCartList(userId int64, page, size int) ([]*model.CartList, int64, error) {
	// 从dao层接收购物车列表，购物车总数，错误信息
	cartList, total, err := dao.GetCartList(userId, page, size)
	if err != nil {
		// 如果出错,说明没查询到购物车，返回购物车列表为nil, 购物车总数为0，购物车不存在的错误码
		return nil, 0, e.NewError(e.ERROR_CART_NOT_EXIST)
	}
	// 如果查询成功就返回列表信息，购物车总数， 为nil的错误信息
	return cartList, total, nil
}

// GetCartById 通过id获取购物车详情
func GetCartById(userId, cartId int64) (*model.CartDetail, error) {
	// 从dao层获取数据
	cart, err := dao.GetCartById(userId, cartId)
	if err != nil {
		// 如果出错,说明没查询到购物车，返回购物车列表为nil, 购物车总数为0，购物车不存在的错误码
		return nil, e.NewError(e.ERROR_CART_NOT_EXIST)
	}
	// 如果查询成功就返回列表信息，购物车总数， 为nil的错误信息
	return cart, nil
}
