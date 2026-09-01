package dao

import (
	"gin-mall/pkg/e"
	"gin-mall/repository/db"
	"gin-mall/repository/db/model"
	"time"

	"gorm.io/gorm"
)

// CreateOrder 订单主表创建函数
func CreateOrder(db *gorm.DB, order *model.Order) error {
	return db.Create(order).Error
}

// CreateOrderItems 批量插入订单商品细明（订单附表）函数
func CreateOrderItems(db *gorm.DB, items []*model.OrderItem) error {
	return db.Create(items).Error
}

// ReduceProductStock 减少商品库存函数 通过商品id找到订单买的商品然后商品库存(stock)减去购买商品数量(quantity)
func ReduceProductStock(db *gorm.DB, productId int64, quantity int64) error {
	// 判断商品库存(stock)需大于于等于购买商品数量(quantity)
	result := db.Model(&model.Product{}).
		Where("id = ? AND stock >= ?", productId, quantity).
		Update("stock", gorm.Expr("stock - ?", quantity))
	if result.Error != nil {
		// 如果有系统错误直接返回

		return result.Error
	}
	if result.RowsAffected == 0 {
		// 如果影响行数为0，说明商品不存在或者商品库存不足,这里统一返回商品库存不足的自定义错误码

		return e.NewError(e.ERROR_PRODUCT_STOCK_SHORTAGE)
	}
	return nil
}

// GetOrderList 获取订单列表（含分页）
func GetOrderList(userId int64, page, pageSize int) ([]*model.Order, int64, error) {
	// 创建切片存储当前页订单信息
	orderList := make([]*model.Order, 0)
	// 创建变量用来存储订单总数
	var total int64
	// 先统计用户的订单总数
	if err := db.DB.
		Model(&model.Order{}).
		Where("user_id = ?", userId).
		Count(&total).
		Error; err != nil {
		// 如果查找出来的总数为0，则返回总数为0和错误信息
		return nil, 0, err
	}
	err := db.DB.
		Model(&model.Order{}).
		// 核对用户id
		Where("user_id = ?", userId).
		// 查询附表信息
		Preload("Items").
		// 按照创建时间由新到旧排序
		Order("create_at DESC").
		// Offset跳过(当前页-1)*每页展示多少数据的数据
		Offset((page - 1) * pageSize).
		// Limit从当前位置开始截取pageSize大小的数据，Find存入数据
		Limit(pageSize).
		// 用Scan存入数据而不是Find，应为Find会在数据库里查询对应表结构，但是cartList是临时定义的结构体
		Find(&orderList).Error
	return orderList, total, err
}

// UpdateOrder 更新订单信息
func UpdateOrder(db *gorm.DB, userId, orderId int64, status int) error {
	// 修改用户订单状态，并且返回错误信息
	result := db.Model(&model.Order{}).
		Where("user_id = ? AND order_id = ?", userId, orderId).
		Update("status", status)
	if result.Error != nil {
		// 如果是系统出错返回错误信息
		return result.Error
	}
	if result.RowsAffected == 0 {
		// 如果是业务逻辑出错返回自定义错误信息：订单不存在
		return e.NewError(e.ERROR_ORDER_NOT_EXIST)
	}
	return nil
}

// IncreaseProductStock 恢复商品库存 通过商品id查找订单买的商品然后商品库存加上购买商品数量
func IncreaseProductStock(db *gorm.DB, productId int64, quantity int64) error {
	result := db.Model(&model.Product{}).
		Where("id = ?", productId).
		Update("stock", gorm.Expr("stock + ?", quantity))
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		// 如果修改了0行数据说明订单不存在，返回订单不存在的错误码
		return e.NewError(e.ERROR_PRODUCT_NOT_EXIST)
	}
	return nil
}

// PayOrder 订单支付函数
func PayOrder(db *gorm.DB, userId, orderId int64) error {
	// 修改订单状态并返回错误信息
	result := db.Model(&model.Order{}).
		// 要确保是同一用户同一订单，同时订单状态是待支付
		Where("user_id = ? AND order_id = ? AND status = 1", userId, orderId).
		// 修改订单信息
		Updates(map[string]interface{}{
			"status": 2, //更新为已支付的状态
			"pay_at": time.Now(),
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return e.NewError(e.ERROR_ORDER_NOT_EXIST)
	}
	return nil
}
