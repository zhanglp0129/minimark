package routers

import (
	"github.com/gin-gonic/gin"
	"minimark/service"
)

func UserRouters(r *gin.RouterGroup) {
	r.POST("/login", UserLogin)
}

func UserLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	token, err := service.UserLogin(username, password)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	c.JSON(200, gin.H{"token": token})
	
}
