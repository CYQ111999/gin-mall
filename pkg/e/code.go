package e

// 设置错误码
const (
	SUCCESS = 200
	ERROR   = 500

	ERROR_USER_NOT_EXIST         = 1000 // 用户不存在
	ERROR_PASSWORD_WRONG         = 1001 // 密码错误
	ERROR_TOKEN_WRONG            = 1002 // 认证失败
	ERROR_SHOULDBINDJSON_WRONG   = 1003 // 请求参数错误
	ERROR_USERNAME_EXIST         = 1004 // 用户名已存在
	ERROR_PRODUCT_NOT_EXIST      = 1005 // 商品不存在
	ERROR_QUANTITY_WRONG         = 1006 // 购物数量不能小于1
	ERROR_NOT_LOGIC              = 1007 // 用户未登录
	ERROR_CART_NOT_EXIST         = 1008 // 购物车不存在
	ERROR_ADDRESS_NOT_EXIST      = 1009 // 地址不存在
	ERROR_PRODUCT_STOCK_SHORTAGE = 1010 // 商品库存不足
	ERROR_ORDER_NOT_EXIST        = 1011 // 订单不存在
)
