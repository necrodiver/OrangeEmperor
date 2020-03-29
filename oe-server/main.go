package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping/:ID", func(c *gin.Context) {
		ID := c.Param("ID")
		t := time.Now()
		c.JSON(200, gin.H{
			"message": "pong",
			"id":      ID,
			"nowTime": t.String(),
		})
	})
	api := router.Group("/api")
	{
		api.GET("/login", loginEndpoint)
	}
	router.Run() // listen and serve on 0.0.0.0:8080
}

// 登录测试
func loginEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "确定了",
		"time":    time.Now().String(),
	})
}
