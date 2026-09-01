package dao

import (
	"gin-mall/pkg/e"
	"gin-mall/repository/db"
	"gin-mall/repository/db/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CreateCart 创建购物车函数
func CreateCart(cart *model.Cart) error {
	err := db.DB.Clauses(clause.OnConflict{
		//根据唯一联合索引判断是否冲突(存在同一用户购买同一商品就算冲突)
		Columns: []clause.Column{
			{Name: "user_id"},
			{Name: "product_id"},
		},
		// 如果冲突就用过去商品数量加新增加的数量作为新数量
		DoUpdates: clause.Set{
			{
				Column: clause.Column{Name: "quantity"},
				Value:  gorm.Expr("quantity + ?", cart.Quantity),
			},
			{
				Column: clause.Column{Name: "updated_at"},
				Value:  gorm.Expr("NOW()"), // 显式让数据库更新为当前时间
			},
		},
	}).Create(cart).Error
	return err
}

// DeleteCart 删除购物车函数
func DeleteCart(cartId, userId int64) error {
	// 校验用户id和购物车id确保是统一用户操作
	result := db.DB.Where("user_id = ? AND cart_id = ?", userId, cartId).
		Delete(&model.Cart{})
	// 如果是系统错误就返回错误信息
	if result.Error != nil {
		return result.Error
	}
	// 如果修改了0行数据说明购物车或者用户id有错，返回购物车不存在的自定义错误信息
	if result.RowsAffected == 0 {
		return e.NewError(e.ERROR_CART_NOT_EXIST)
	}
	return nil
}

// GetCartList 获取购物车列表（含分页）
func GetCartList(userId int64, page, pageSize int) ([]*model.CartList, int64, error) {
	// 现在创建切片用来存储当前页信息
	cartList := make([]*model.CartList, 0)
	// 创建total变量用来存储商品总数
	var total int64
	// 先统计商品总数
	if err := db.DB.Model(&model.Cart{}).Count(&total).Error; err != nil {
		// 如果出错说明总数为0返回查询信息为nil,商品总数为0和自定义错误码:商品不存在
		return nil, 0, err
	}
	err := db.DB.
		Model(&model.Cart{}).
		// Select选择返回字段信息
		Select("carts.cart_id, carts.user_id, carts.product_id, carts.quantity, "+
			"IFNULL(products.name, '') as product_name, IFNULL(products.price, 0) as product_price").
		// 利用购物车的商品id将购物车和商品两个表建立左连接
		Joins("left join products on carts.product_id = products.id").
		// 核对用户id
		Where("carts.user_id = ?", userId).
		// Offset跳过(当前页-1)*每页展示多少数据的数据
		Offset((page - 1) * pageSize).
		// Limit从当前位置开始截取pageSize大小的数据，Find存入数据
		Limit(pageSize).
		// 用Scan存入数据而不是Find，应为Find会在数据库里查询对应表结构，但是cartList是临时定义的结构体
		Scan(&cartList).Error
	// 返回当前页信息，商品总数，错误信息
	return cartList, total, err
}

// GetCartById 通过id获取购物车详情
func GetCartById(userId, cartId int64) (*model.CartDetail, error) {
	// 创建返回购物车时的临时结构体变量
	var cart model.CartDetail
	// 从数据库查询
	err := db.DB.
		Model(&model.Cart{}).
		// Select选择返回字段信息
		Select("carts.cart_id, carts.user_id, carts.product_id, carts.quantity, "+
			"IFNULL(products.name, '') as product_name, IFNULL(products.price, 0) as product_price, "+
			"carts.created_at, carts.updated_at").
		// 利用购物车的商品id将购物车和商品两个表建立左连接
		Joins("left join products on carts.product_id = products.id").
		// 核对用户id
		Where("carts.user_id = ? and carts.cart_id = ?", userId, cartId).
		// 用Scan存入数据而不是Find，应为Find会在数据库里查询对应表结构，但是cart是临时定义的结构体
		Scan(&cart).Error
	// 返回查询结果
	return &cart, err
}
