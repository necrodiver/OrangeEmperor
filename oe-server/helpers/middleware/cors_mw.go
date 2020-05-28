package middleware

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"oe_server/helpers/utils"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var corsUrls []string = []string{"*", "https://fanyi.baidu.com", "https://www.baidu.com"}

// CorsHeaderMiddleware 跨域请求配置/中间件
func CorsHeaderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cors.New(cors.Config{
			AllowOrigins:     corsUrls,
			AllowMethods:     []string{"PUT", "GET", "POST", "UPDATE", "DELETE"},
			AllowHeaders:     []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization"},
			ExposeHeaders:    []string{"Content-Length", "Token"},
			AllowCredentials: true,
			// 以下AllowOriginFunc 若设置了，那么以上的AllowOrigins将失效
			// AllowOriginFunc: func(origin string) bool {
			// 	return origin == "https://github.com"
			// },
			MaxAge: 12 * time.Hour,
		})

		method := c.Request.Method
		// origin := c.Request.Header.Get("Origin")
		// if origin != "" {
		// 	// c.Header("Access-Control-Allow-Origin", origin)
		// 	// c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		// 	// c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		// 	// c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		// 	// c.Header("Access-Control-Allow-Credentials", "true")
		// 	c.Set("content-type", "application/json")
		// }
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		getData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(400, gin.H{"error": dataErr})
			c.Abort()
			return
		}

		buf := bytes.NewBuffer(getData)
		c.Request.Body = ioutil.NopCloser(buf)
		// 日志记录一下请求
		utils.ZapLogger.Info("     Origin:", c.Request.Header.Get("Origin"), "     url:", c.Request.RequestURI, "     header:", c.Request.Header, "     Body:", string(getData[:]))
		c.Next()
	}
}
