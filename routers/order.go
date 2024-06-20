package routers

import (
	"github.com/gin-gonic/gin"
	"minimark/pojo"
	"minimark/service"
	"strconv"
)

func OrderRouters(r *gin.RouterGroup) {
	r.POST("/", OrderCreate)
	r.GET("/", OrderPage)
	r.GET("/:id", OrderFind)
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

func OrderPage(c *gin.Context) {
	var dto pojo.OrderPageDTO
	err := c.ShouldBindQuery(&dto)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}

	orders, err := service.OrderPage(&dto)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	c.JSON(200, orders)
}

func OrderFind(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}

	order, err := service.OrderFind(id)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	c.JSON(200, order)
}
