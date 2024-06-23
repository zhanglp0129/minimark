package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"minimark"
	"minimark/utils"
	"net/http"
	"sync"
)

var (
	routers     *gin.Engine
	onceRouters sync.Once
)

// LoginAuthorization 登录验证
func LoginAuthorization(c *gin.Context) {
	if c.Request.URL.Path == "/api/user/login" {
		return
	}
	token := c.GetHeader("Authorization")
	if token == "" {
		c.String(401, "登录信息错误或已过期，请重新登录")
		c.Abort()
		return
	}
	t, err := utils.ParseJwtToken(token)
	if err != nil {
		c.String(401, "登录信息错误或已过期，请重新登录")
		c.Abort()
		return
	}

	claims := t.Claims.(jwt.MapClaims)
	c.Set("username", claims["username"])
}

func GetRouters() *gin.Engine {
	onceRouters.Do(func() {
		routers = gin.Default()
		// 提供嵌入的前端资源，并将访问前端资源的请求重定向到正确的未知
		routers.StaticFS("/static", http.FS(minimark.WebFS))
		routers.GET("/assets/:filename", func(c *gin.Context) {
			fmt.Println(c.Request.URL.Path)
			c.Redirect(301, "/static/web/dist/assets/"+c.Param("filename"))
		})
		// 将所有请求重定向到index.html，以便Vue处理路由
		routers.NoRoute(func(c *gin.Context) {
			file, err := minimark.WebFS.ReadFile("web/dist/index.html")
			if err != nil {
				c.String(http.StatusInternalServerError, "Internal Server Error")
				return
			}
			c.Data(200, "text/html; charset=utf-8", file)
		})

		// api子路由
		api := routers.Group("/api")
		api.Use(LoginAuthorization)
		// 用户子路由
		user := api.Group("/user")
		UserRouters(user)
		// 商品分类子路由
		category := api.Group("/category")
		CategoryRouters(category)
		// 商品子路由
		goods := api.Group("/goods")
		GoodsRouters(goods)
		// 支付方式子路由
		payMethod := api.Group("/pay_method")
		PayMethodRouters(payMethod)
		// 订单子路由
		order := api.Group("/order")
		OrderRouters(order)
		// 进货子路由
		procurement := api.Group("/procurement")
		ProcurementRouters(procurement)
	})
	return routers
}
