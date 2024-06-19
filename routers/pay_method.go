package routers

import (
	"github.com/gin-gonic/gin"
	"minimark/dao"
	"minimark/service"
	"strconv"
)

func PayMethodRouters(r *gin.RouterGroup) {
	r.POST("/", PayMethodCreate)
	r.GET("/", PayMethodList)
	r.GET("/:id", PayMethodFind)
	r.PUT("/:id", PayMethodUpdate)
	r.DELETE("/", PayMethodDelete)
}

// PayMethodCreate 创建商品分类
func PayMethodCreate(c *gin.Context) {
	var payMethod dao.PayMethod
	err := c.ShouldBindJSON(&payMethod)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	err = service.PayMethodCreate(&payMethod)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	c.String(200, "")
}

// PayMethodList 商品分类列表查询，查询所有
func PayMethodList(c *gin.Context) {
	payMethod, err := service.PayMethodList()
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	c.JSON(200, payMethod)
}

// PayMethodFind 根据id查询商品分类
func PayMethodFind(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}

	payMethod, err := service.PayMethodFind(id)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	c.JSON(200, payMethod)
}

// PayMethodUpdate 更新商品分类
func PayMethodUpdate(c *gin.Context) {
	var payMethod dao.PayMethod
	err := c.ShouldBindJSON(&payMethod)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	payMethod.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}

	err = service.PayMethodUpdate(&payMethod)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	c.String(200, "")
}

// PayMethodDelete 删除商品分类
func PayMethodDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}

	err = service.PayMethodDelete(id)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	c.String(200, "")
}
