package authRouter

import (
	"database/sql"
	authController "insurance/controller/auth"

	"github.com/gin-gonic/gin"

	middleware "insurance/middleware"

	redis "github.com/go-redis/redis/v8"
)

type RouterAuthInsurance struct {
	Router *gin.Engine
	DB *sql.DB
	Redis *redis.Client
}

func NewRouterAuthInsurance(router *gin.Engine, db *sql.DB, rdb *redis.Client) *RouterAuthInsurance {
	return &RouterAuthInsurance{
		Router: router,
		DB: db,
		Redis: rdb,
	}
}

func (ri *RouterAuthInsurance) SetupAuthRouter() {

	authControllerInstance := authController.NewAuthController(ri.DB, ri.Redis) 

	auth := ri.Router.Group("/assessors")
	auth.Use(middleware.AuthMiddleware())

	{
		auth.POST("/login", authControllerInstance.LoginHandler)
		auth.GET("/test-middleware", authControllerInstance.TestMiddleware)
	}
}
