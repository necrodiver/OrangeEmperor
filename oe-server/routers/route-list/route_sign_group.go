package routes

import (
	"net/http"
	"oe_server/helpers/msgcodes"
	"oe_server/helpers/utils"
	"oe_server/models"
	"oe_server/service"
	"time"

	"github.com/gin-gonic/gin"
)

type login struct {
	UserName string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

// GetTime 请求index路由
func GetTime(r *gin.Engine) {
	userList := models.StartDb()
	r.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"time": time.Now().Format("2006-01-02 15:04:05"),
			"msg":  userList,
		})
	})
}

// SignGroup 签名组:登录/登出/注册
func SignGroup(r *gin.Engine) {
	group := r.Group("/Sign")
	{
		group.POST("/In", signIn)
		group.GET("/Out", signOut)
		group.POST("/Register", signRegister)
		group.DELETE("/LogOff", signLogoff)
	}
}

// 登录
func signIn(c *gin.Context) {
	var json login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	authInfo, err := service.SignIn(json.UserName, json.Password)
	if err != nil {
		c.JSON(401, gin.H{
			"code":   msgcodes.Fail,
			"status": "请检查你的用户名或密码是否正确",
		})
		return
	}
	if authInfo.Id > 0 {
		token, err := utils.CreateToken(authInfo.Id)
		if err != nil {
			c.SecureJSON(200, gin.H{"code": msgcodes.Fail, "msg": "登录失败，请重试"})
			return
		}
		c.Header("atom", token)
		// c.Request.Response.Header.Add("atom", token)
		c.JSON(200, gin.H{"code": msgcodes.Success})
		return
	}
	// if json.UserName != "zhangsan" || json.Password != "123456" {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	// 	return
	// }
	c.SecureJSON(200, gin.H{"code": msgcodes.Fail, "msg": "登录失败，请重试"})
}

// 登出
func signOut(c *gin.Context) {

}

// 注册
func signRegister(c *gin.Context) {

}

// 注销
func signLogoff(c *gin.Context) {

}
