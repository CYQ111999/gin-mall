package middleware

import (
	"gin-mall/pkg/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 根据token鉴定用户权限的中间件(需要登录才能执行之后的逻辑)
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1.从请求头中获取Authorization
		authHeader := c.GetHeader("Authorization")
		// 如果为空说明token不存在用户未登录
		if authHeader == "" {
			// 返回未登录
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "未登录，请先登录或注册",
			})
			// 终止后续业务
			c.Abort()
			// 返回
			return
		}
		// 如果不为空

		// 2.解析token格式
		// 按照空格分割为2部分
		parts := strings.SplitN(authHeader, " ", 2)
		// 如果不满足2部分和第一部分是"Bearer"
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			// 返回token格式错误
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "token格式错误",
			})
			// 终止后续业务
			c.Abort()
			// 返回
			return
		}

		// 3.解析验证token是否正确
		tokenString := parts[1]
		claims, err := util.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "token无效或过期",
			})
			c.Abort()
			return
		}
		// 4.通过验证将用户信息存入上下文
		c.Set("user_id", claims.UserId)
		// 执行后续业务逻辑
		c.Next()
	}
}
