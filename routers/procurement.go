package routers

import (
	"github.com/gin-gonic/gin"
	"minimark/dao"
	"minimark/pojo"
	"minimark/service"
	"strconv"
)

func ProcurementRouters(r *gin.RouterGroup) {
	r.POST("", ProcurementCreate)
	r.GET("", ProcurementPage)
	r.GET("/:id", ProcurementFind)
	r.PUT("/:id", ProcurementUpdate)
	r.POST("/add-goods/:id", ProcurementAddGoods)
	r.PATCH("/update-goods/:id", ProcurementUpdateGoods)
	r.DELETE("", ProcurementDelete)
}

func ProcurementCreate(c *gin.Context) {
	dto := pojo.ProcurementCreateDTO{}
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}

	err = service.ProcurementCreate(&dto)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	c.String(200, "")
}

func ProcurementPage(c *gin.Context) {
	var dto pojo.ProcurementPageDTO
	err := c.ShouldBindQuery(&dto)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}

	procurements, total, err := service.ProcurementPage(&dto)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"items": procurements,
		"total": total,
	})
}

func ProcurementFind(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}

	procurement, err := service.ProcurementFind(id)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	c.JSON(200, procurement)
}

// ProcurementUpdate 修改订单基本信息
func ProcurementUpdate(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}

	var procurement dao.Procurement
	err = c.ShouldBindJSON(&procurement)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	procurement.ID = id

	err = service.ProcurementUpdate(&procurement)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	c.String(200, "")
}

func ProcurementAddGoods(c *gin.Context) {
	// 获取订单id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}

	// 获取添加的商品
	var dto []pojo.ProcurementAddUpdateGoodsDTO
	err = c.ShouldBindJSON(&dto)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}

	// 执行service
	err = service.ProcurementAddGoods(id, dto)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	c.String(200, "")
}

func ProcurementUpdateGoods(c *gin.Context) {
	// 获取订单id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}

	// 获取修改后的信息
	var dto []pojo.ProcurementAddUpdateGoodsDTO
	err = c.ShouldBindJSON(&dto)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}

	// 执行service
	err = service.ProcurementUpdateGoods(id, dto)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	c.String(200, "")
}

func ProcurementDelete(c *gin.Context) {
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

	err = service.ProcurementDelete(dto.ID, *dto.ChangeStock)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	c.String(200, "")
}
