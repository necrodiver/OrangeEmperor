package middleware

import (
	"github.com/gin-gonic/gin"
	// "oe_server/helpers/setting"
	// "oe_server/helpers/utils"
)

var (
	noAtom      = "违规登录"   // 无atom参数导致违法登录
	dataErr     = "参数获取错误" // 获取请求参数错误
	verifiyFail = "校验失败"   // 参数的加密和atom对比不匹配
)

// ValidateMiddleware 验证请求参数是否进行加密
func ValidateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 这里需要进行传入对应url判断是否需要atom这个token来约束，如果是登录接口，起初是没有atom
		//setting.NoValidateUrls
		// if(setting)
		// atomStr := c.GetHeader("atom")

		// if atomStr == nil {
		// 	c.JSON(401, gin.H{"error": noAtom})
		// 	c.Abort()
		// 	return
		// }
		// atom := utils.RSADecrypt(byte[](atomStr))

		c.Next()
	}
}
