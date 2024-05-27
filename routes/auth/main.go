package authRouter

import (
	"database/sql"
	authController "insurance/controller/auth"

	"github.com/gin-gonic/gin"
)

type RouterAuthInsurance struct {
	Router *gin.Engine
	DB *sql.DB
}

func NewRouterAuthInsurance(router *gin.Engine, db *sql.DB) *RouterAuthInsurance {
	return &RouterAuthInsurance{
		Router: router,
		DB: db,
	}
}

func (ri *RouterAuthInsurance) SetupAuthRouter() {

	authControllerInstance := authController.NewAuthController(ri.DB) 

	auth := ri.Router.Group("/assessors")
	{
		auth.POST("/login", authControllerInstance.LoginHandler)
	}
}
