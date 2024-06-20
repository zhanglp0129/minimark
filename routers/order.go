package routers

import (
	"github.com/gin-gonic/gin"
	"minimark/pojo"
	"minimark/service"
)

func OrderRouters(r *gin.RouterGroup) {
	r.POST("/", OrderCreate)
}

func OrderCreate(c *gin.Context) {
	dto := pojo.OrderCreateDTO{}
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}

	err = service.OrderCreate(&dto)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	c.String(200, "")
}
