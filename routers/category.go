package routers

import (
	"github.com/gin-gonic/gin"
	"minimark/dao"
	"minimark/service"
	"strconv"
)

func CategoryRouters(r *gin.RouterGroup) {
	r.POST("/", CategoryCreate)
	r.GET("/", CategoryList)
	r.GET("/:id", CategoryFind)
	r.PUT("/:id", CategoryUpdate)
	r.DELETE("/", CategoryDelete)
}

// CategoryCreate 创建商品分类
func CategoryCreate(c *gin.Context) {
	var category dao.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	err = service.CategoryCreate(&category)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	c.String(200, "")
}

// CategoryList 商品分类列表查询，查询所有
func CategoryList(c *gin.Context) {
	categories, err := service.CategoryList()
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	c.JSON(200, categories)
}

// CategoryFind 根据id查询商品分类
func CategoryFind(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}

	category, err := service.CategoryFind(id)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	c.JSON(200, category)
}

// CategoryUpdate 更新商品分类
func CategoryUpdate(c *gin.Context) {
	var category dao.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	category.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}

	err = service.CategoryUpdate(&category)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	c.String(200, "")
}

// CategoryDelete 删除商品分类
func CategoryDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}

	err = service.CategoryDelete(id)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	c.String(200, "")
}
