package routes

import (
	"database/sql"
	authRouter "insurance/routes/auth"

	"github.com/gin-gonic/gin"
)

type RouterInsurance struct {
	Router *gin.Engine
	DB *sql.DB
}

func NewRouterInsurance(router *gin.Engine, db *sql.DB) *RouterInsurance {
	return &RouterInsurance{
		Router: router,
		DB: db,
	}
}

func (ri *RouterInsurance) Setup() {
	   // Define additional routes here

	authRouterInstance := authRouter.NewRouterAuthInsurance(ri.Router, ri.DB)
  authRouterInstance.SetupAuthRouter()
}