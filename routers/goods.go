package routers

import (
	"github.com/gin-gonic/gin"
	"minimark/dao"
	"minimark/pojo"
	"minimark/service"
	"strconv"
	"strings"
)

func GoodsRouters(r *gin.RouterGroup) {
	r.POST("/", GoodsCreate)
	r.GET("/", GoodsPage)
	r.GET("/:ids", GoodsFind)
	r.PUT("/:id", GoodsUpdate)
	r.DELETE("/", GoodsDelete)
}

// GoodsCreate 新增商品
func GoodsCreate(c *gin.Context) {
	goods := dao.Goods{}
	err := c.ShouldBindJSON(&goods)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	err = service.GoodsCreate(&goods)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	c.String(200, "")
}

// GoodsPage 商品的条件分页查询
func GoodsPage(c *gin.Context) {
	goodsPageDTO := pojo.GoodsPageDTO{}
	err := c.ShouldBindQuery(&goodsPageDTO)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	goods, err := service.GoodsPage(goodsPageDTO)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	c.JSON(200, goods)
}

// GoodsFind 根据id查询商品
func GoodsFind(c *gin.Context) {
	idsString := c.Param("ids")
	idsSlice := strings.Split(idsString, ",")
	ids := make([]int, len(idsSlice))
	for i := range idsSlice {
		id, err := strconv.Atoi(idsSlice[i])
		if err != nil {
			c.String(400, err.Error())
			c.Abort()
			return
		}
		ids[i] = id
	}

	goods, err := service.GoodsFind(ids)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	c.JSON(200, goods)
}

// GoodsUpdate 修改商品信息
func GoodsUpdate(c *gin.Context) {
	goods := dao.Goods{}
	err := c.ShouldBindJSON(&goods)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	goods.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}

	err = service.GoodsUpdate(&goods)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	c.String(200, "")
}

func GoodsDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	err = service.GoodsDelete(id)
	if err != nil {
		c.String(400, err.Error())
		c.Abort()
		return
	}
	c.String(200, "")
}
