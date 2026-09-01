package util

import (
	"errors"
	"gin-mall/pkg/e"

	"github.com/gin-gonic/gin"
)

// ParamError 请求参数错误响应
func ParamError(c *gin.Context) {
	// 如果解析参数时返回err,则返回请求参数错误的json信息
	c.JSON(400, gin.H{
		"code": e.ERROR_SHOULDBINDJSON_WRONG,
		"msg":  e.GetMsg(e.ERROR_SHOULDBINDJSON_WRONG),
	})
}

// HandleError 按照自定义错误和系统错误返回不同响应
func HandleError(c *gin.Context, err error) {
	var appErr *e.Apperror
	// 如果是自定义错误类型
	if errors.As(err, &appErr) {
		c.JSON(400, gin.H{"code": appErr.Code, "msg": appErr.Msg})
		return
	}
	// 如果是系统错误
	c.JSON(e.ERROR, gin.H{
		"error": "系统出错",
	})
}

// Success 返回成功响应
func Success(c *gin.Context, msg string, data interface{}) {
	resp := gin.H{"code": e.SUCCESS, "msg": msg}
	// 如果data不为空就赋值
	if data != nil {
		resp["data"] = data
	}
	c.JSON(e.SUCCESS, resp)
}
