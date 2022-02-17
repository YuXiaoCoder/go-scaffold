package router

import (
	"go-scaffold/controllers"
	"go-scaffold/middlerwares"
	"go-scaffold/pkg/configs"
	"time"

	"go.uber.org/zap"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"

	_ "go-scaffold/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 设置路由
func Setup() *gin.Engine {
	// 创建引擎
	engine := gin.New()

	// 运行模式
	gin.SetMode(configs.AllConfig.Basic.Mode)

	// 注册公共中间件
	engine.Use(cors.Default())
	engine.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	engine.Use(ginzap.RecoveryWithZap(zap.L(), true))
	engine.Use(middlerwares.RateLimit(1*time.Second, 10))

	// 注册路由
	registerRoutes(engine)

	// 返回引擎
	return engine
}

// 注册路由
func registerRoutes(engine *gin.Engine) {
	// 文档路由
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// V1 API（不需要认证）
	v1Group := engine.Group("/api/v1")
	{
		// 登录
		v1Group.POST("/sign_in", controllers.SignInHandler)
		// 注册
		v1Group.POST("/sign_up", controllers.SignUPHandler)
	}

	// V1 API（需要认证）
	v1GroupAuth := engine.Group("/api/v1")
	// JWT认证中间件
	v1GroupAuth.Use(middlerwares.JWTAuth())
	{
		// 用户
		v1GroupAuth.GET("/users", controllers.ListUserHandler)              // 用户列表
		v1GroupAuth.GET("/users/:id", controllers.GetUserByIDHandler)       // 指定用户
		v1GroupAuth.PUT("/users/:id", controllers.UpdateUserByIDHandler)    // 更新用户
		v1GroupAuth.DELETE("/users/:id", controllers.DeleteUserByIDHandler) // 删除用户
	}

	// 未匹配到路由信息
	engine.NoRoute(func(ctx *gin.Context) {
		controllers.ResponseError(ctx, controllers.CodeInterfaceNotExist)
	})
}
