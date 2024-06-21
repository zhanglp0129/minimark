package routers

import (
	"github.com/gin-gonic/gin"
	"minimark/dao"
	"minimark/pojo"
	"minimark/service"
	"strconv"
)

func OrderRouters(r *gin.RouterGroup) {
	r.POST("/", OrderCreate)
	r.GET("/", OrderPage)
	r.GET("/:id", OrderFind)
	r.PUT("/:id", OrderUpdate)
	r.POST("/add-goods/:id", OrderAddGoods)
	r.PATCH("/update-goods/:id", OrderUpdateGoods)
	r.DELETE("/", OrderDelete)
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

// OrderUpdate 修改订单基本信息
func OrderUpdate(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}

	var order dao.Order
	err = c.ShouldBindJSON(&order)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	order.ID = id

	err = service.OrderUpdate(&order)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	c.String(200, "")
}

func OrderAddGoods(c *gin.Context) {
	// 获取订单id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}

	// 获取修改后的信息
	var dto []pojo.OrderAddUpdateGoodsDTO
	err = c.ShouldBindJSON(&dto)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}

	// 执行service
	err = service.OrderAddGoods(id, dto)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	c.String(200, "")
}

func OrderUpdateGoods(c *gin.Context) {
	// 获取订单id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}

	// 获取修改后的信息
	var dto []pojo.OrderAddUpdateGoodsDTO
	err = c.ShouldBindJSON(&dto)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}

	// 执行service
	err = service.OrderUpdateGoods(id, dto)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	c.String(200, "")
}

func OrderDelete(c *gin.Context) {
	type DTO struct {
		ID          int   `form:"id"`
		ChangeStock *bool `form:"change_stock"`
	}
	dto := DTO{}
	err := c.ShouldBindQuery(&dto)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	if dto.ChangeStock == nil {
		dto.ChangeStock = new(bool)
		*dto.ChangeStock = true
	}

	err = service.OrderDelete(dto.ID, *dto.ChangeStock)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	c.String(200, "")
}
