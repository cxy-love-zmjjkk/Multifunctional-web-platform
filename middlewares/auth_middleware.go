package middlewares

import (
	"exchange_backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

//中间件是为了过滤路由而发明的一种机制
//http先经过中间件再到具体的处理函数（路由）
//这个文件是用来验证JWT的

func AuthMidddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization Header"})
			ctx.Abort()
			return
		}

		username, err := utils.ParseJWT(token)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization Header"})
			ctx.Abort()
			return
		}
		//验证中间件之后，将用户信息存储到context中
		//后续处理函数（Handler）可以通过ctx.Get("username")访问这个值
		ctx.Set("username", username)

		ctx.Next()
	}
}
