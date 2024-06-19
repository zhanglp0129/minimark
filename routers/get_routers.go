package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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
		routers.Use(LoginAuthorization)

		// 用户子路由
		user := routers.Group("/user")
		UserRouters(user)
		// 商品分类子路由
		category := routers.Group("/category")
		CategoryRouters(category)
		// 商品子路由
		goods := routers.Group("/goods")
		GoodsRouters(goods)
		// 支付方式子路由
		payMethod := routers.Group("/pay_method")
		PayMethodRouters(payMethod)
	})
	return routers
}
