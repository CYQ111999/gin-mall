package e

// MsgFlags 错误码对应的错误信息
var MsgFlags = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",

	ERROR_USER_NOT_EXIST:         "用户不存在",
	ERROR_PASSWORD_WRONG:         "密码错误",
	ERROR_TOKEN_WRONG:            "认证失败",
	ERROR_SHOULDBINDJSON_WRONG:   "请求参数错误",
	ERROR_USERNAME_EXIST:         "用户名已存在",
	ERROR_PRODUCT_NOT_EXIST:      "商品不存在",
	ERROR_QUANTITY_WRONG:         "购物数量不能小于1",
	ERROR_NOT_LOGIC:              "用户未登录",
	ERROR_CART_NOT_EXIST:         "购物车不存在",
	ERROR_ADDRESS_NOT_EXIST:      "地址不存在",
	ERROR_PRODUCT_STOCK_SHORTAGE: "商品库存不足",
	ERROR_ORDER_NOT_EXIST:        "订单不存在",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
