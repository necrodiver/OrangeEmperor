package route

import (
	"oe_server/helpers/middleware"
	"oe_server/helpers/setting"
	routes "oe_server/routers/route-list"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化Router
func InitRouter() {
	router := gin.Default()
	// 中间件-跨域验证
	router.Use(middleware.CorsHeaderMiddleware())
	// 中间件-加密验证
	router.Use(middleware.ValidateMiddleware())

	routes.GetTime(router)
	routes.SignGroup(router)
	router.Run(":" + setting.HTTP_PORT)
}
