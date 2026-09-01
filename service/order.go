package service

import (
	"fmt"
	"gin-mall/model"
	"gin-mall/pkg/e"
	"gin-mall/repository/db"
	"gin-mall/repository/db/dao"
	dbmodel "gin-mall/repository/db/model"
	"time"

	"gorm.io/gorm"
)

// CreateOrder 创建订单信息函数
func CreateOrder(userId int64, req model.OrderCreateReq) (string, error) {
	// 1. 获取地址信息，并校验地址是否属于当前用户
	address, err := dao.GetAddressById(req.AddressID)
	if err != nil || address.UserID != userId {
		// 如果返回错误或者用户id不匹配就返回地址不存在的错误信息
		return "", e.NewError(e.ERROR_ADDRESS_NOT_EXIST)
	}
	// 构建地址快照
	addressSnapshot := dbmodel.AddressSnapshot{
		Name:    address.Name,
		Phone:   address.Phone,
		Address: address.Address,
	}
	// 2. 根据前端传的购物车ID列表，循环查出每个购物车的信息
	// 先创建购物车详情切片
	var cartItems []*dbmodel.CartDetail
	for _, id := range req.CartItemIDs {
		var item *dbmodel.CartDetail
		item, err = dao.GetCartById(userId, id)
		if err != nil {
			return "", err
		}
		// 把每个id查找到的信息拼接到切片
		cartItems = append(cartItems, item)
	}
	// 如果最后切片长度为0，说明购物车不存在，返回错误信息
	if len(cartItems) == 0 {
		return "", e.NewError(e.ERROR_CART_NOT_EXIST)
	}
	// 3. 组装订单主表和明细切片，同时计算总价
	var totalAmount int64
	var orderItems []*dbmodel.OrderItem
	for _, item := range cartItems {
		// 总价是循环累加每个购物车详情里的商品数乘商品价格
		totalAmount += item.Quantity * item.ProductPrice
		// 拼接每个订单详情
		orderItems = append(orderItems, &dbmodel.OrderItem{
			ProductID:    item.ProductID,
			ProductName:  item.ProductName,
			Quantity:     item.Quantity,
			PriceAtOrder: item.ProductPrice,
			CreateAt:     time.Now(),
		})
	}
	// 生成唯一的业务订单号（雪花算法）
	orderSN := fmt.Sprintf("%d%d", time.Now().Unix(), userId)
	// 赋值完善订单信息
	order := &dbmodel.Order{
		OrderSN:         orderSN,
		UserID:          userId,
		AddressSnapshot: addressSnapshot,
		TotalAmount:     totalAmount,
		Status:          1,          // 待支付
		CreateAt:        time.Now(), // 手动赋值当前时间
		UpdateAt:        time.Now(), // 手动赋值当前时间
	}
	// 4.将相关逻辑打包成事务
	err = db.DB.Transaction(func(tx *gorm.DB) error {
		// 4.1 插入订单主表
		if err = dao.CreateOrder(tx, order); err != nil {
			return err
		}

		// 4.2 填充 OrderID 并插入订单明细表
		for _, item := range orderItems {
			item.OrderID = order.OrderID
		}
		if err = dao.CreateOrderItems(tx, orderItems); err != nil {
			return err
		}

		// 4.3 循环扣减对应商品的库存
		for _, item := range cartItems {
			if err = dao.ReduceProductStock(tx, item.ProductID, item.Quantity); err != nil {
				// 如果出错说明是库存不足返回从dao层获取的库存不足的错误信息直接返回
				return err
			}
		}
		return nil // 返回 nil 自动提交事务
	})
	return orderSN, err
}

// GetOrderList 获取订单列表函数（含分页）
func GetOrderList(userId int64, page, pageSize int) ([]*model.OrderResponse, int64, error) {
	// 从dao层获取当前页的订单信息
	orderList, total, err := dao.GetOrderList(userId, page, pageSize)
	if err != nil {
		// 将返回的错误信息包装成自定义错误：订单不存在
		return nil, 0, e.NewError(e.ERROR_ORDER_NOT_EXIST)
	}
	// 如果没出错返回当前页订单信息，订单总数，nil

	// 为了防止返回过多信息，我们创建新的结构体储存需要返回的变量，并赋值
	var orders = make([]*model.OrderResponse, 0)
	for _, order := range orderList {
		// 1. 组装主表字段
		orderResp := &model.OrderResponse{
			OrderSN:         order.OrderSN,
			TotalAmount:     order.TotalAmount,
			Status:          order.Status,
			CreateAt:        order.CreateAt,
			AddressSnapshot: order.AddressSnapshot,
		}

		// 2. 循环组装附表商品明细
		for _, item := range order.Items {
			itemResp := model.OrderItemResponse{
				ProductID:   item.ProductID,
				ProductName: item.ProductName,
				Quantity:    int32(item.Quantity),
				Price:       item.PriceAtOrder,
			}
			orderResp.Items = append(orderResp.Items, itemResp)
		}

		orders = append(orders, orderResp)
	}
	return orders, total, nil
}

// CancelOrder 取消用户订单函数
func CancelOrder(userId, orderId int64) error {
	// 开启取消订单和恢复商品库存的事务
	return db.DB.Transaction(func(tx *gorm.DB) error {
		// 1. 查询校验需要取消的订单（订单需要是用户自己的，并且状态是未支付的
		var order dbmodel.Order
		if err := tx.Where("user_id = ? AND order_id = ? AND status = 1", userId, orderId).
			First(&order).Error; err != nil {
			// 如果出错返回订单不存在的错误信息并回滚
			return e.NewError(e.ERROR_ORDER_NOT_EXIST)
		}
		// 2. 修改订单状态
		if err := dao.UpdateOrder(tx, userId, orderId, 3); err != nil {
			return err
		}
		// 3. 循环恢复订单里所有商品的商品库存
		// 先取出订单附表里的商品细明
		var orderItems []*dbmodel.OrderItem
		if err := tx.Where("order_id = ?", orderId).
			Find(&orderItems).Error; err != nil {
			return err
		}
		// 再开始循环恢复商品库存
		for _, item := range orderItems {
			if err := dao.IncreaseProductStock(tx, item.ProductID, item.Quantity); err != nil {
				return err
			}
		}
		return nil
	})
}

// PayOrder 订单支付函数
func PayOrder(userId, orderId int64) error {
	// 这里支付是假的支付，不过也写成事务以后可以延伸拓展
	return db.DB.Transaction(func(tx *gorm.DB) error {
		err := dao.PayOrder(tx, userId, orderId)
		if err != nil {
			return err
		}
		return nil
	})
}
