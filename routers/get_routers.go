package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"minimark"
	"minimark/utils"
	"sync"
)

var (
	routers     *gin.Engine
	onceRouters sync.Once
)

// LoginAuthorization 登录验证
func LoginAuthorization(c *gin.Context) {
	if c.Request.URL.Path == "/user/login" {
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
		// 加载前端资源
		routers.GET("/", func(c *gin.Context) {
			c.Writer.WriteHeader(200)
			c.Header("Content-Type", "text/html; charset=utf-8")
			b, _ := minimark.WebFS.ReadFile("index.html")
			_, _ = c.Writer.Write(b)
			c.Writer.Flush()
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
