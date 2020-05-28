package middleware

import (
	"github.com/gin-gonic/gin"
)

var (
	noAtom      = "违规登录"   // 无atom参数导致违法登录
	dataErr     = "参数获取错误" // 获取请求参数错误
	verifiyFail = "校验失败"   // 参数的加密和atom对比不匹配
)

// ValidateMiddleware 验证请求参数是否进行加密
func ValidateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		atomStr := c.GetHeader("atom")
		if atomStr == "" {
			c.JSON(401, gin.H{"error": noAtom})
			c.Abort()
			return
		}
		// var body = c.Request.Body
		// getData, err := ioutil.ReadAll(body)
		// if err != nil {
		// 	c.JSON(400, gin.H{"error": dataErr})
		// 	c.Abort()
		// 	return
		// }
		c.Next()
	}
}
