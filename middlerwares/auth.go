package middlerwares

import (
	"go-scaffold/controllers"
	"go-scaffold/dao/rds"
	"go-scaffold/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuth() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// 从请求头中获取Token
		tokenStr := ctx.Request.Header.Get("Authorization")
		if tokenStr == "" {
			controllers.ResponseError(ctx, controllers.CodeNeedAuth)
			ctx.Abort()
			return
		}

		// Bearer认证的格式：Bearer TOKEN
		parts := strings.SplitN(tokenStr, " ", 2)
		if len(parts) != 2 && parts[0] == "Bearer" {
			controllers.ResponseError(ctx, controllers.CodeInvalidAuth)
			ctx.Abort()
			return
		}

		// 解析Token
		token, err := jwt.ParseToken(parts[1])
		if err != nil {
			controllers.ResponseError(ctx, controllers.CodeInvalidAuth)
			ctx.Abort()
			return
		}

		// 检测用户权限
		_, err = rds.GetUserDB().GetByID(token.UserID)
		if err != nil {
			controllers.ResponseError(ctx, controllers.CodeInvalidAuth)
			ctx.Abort()
			return
		}

		// 将用户信息保存到当前请求的上下文中
		ctx.Set(controllers.CTXUserKey, token.UserID)
		ctx.Next()
	}
}
