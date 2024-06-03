package routes

import (
	"database/sql"
	authRouter "insurance/routes/auth"
	claimFolderRouter "insurance/routes/claimFolder"

	"github.com/gin-gonic/gin"

	redis "github.com/go-redis/redis/v8"
)

type RouterInsurance struct {
	Router *gin.Engine
	DB *sql.DB
	Redis *redis.Client
}

func NewRouterInsurance(router *gin.Engine, db *sql.DB, rdb *redis.Client) *RouterInsurance {
	return &RouterInsurance{
		Router: router,
		DB: db,
		Redis: rdb,
	}
}

func (ri *RouterInsurance) Setup() {
	   // Define additional routes here

	authRouterInstance := authRouter.NewRouterAuthInsurance(ri.Router, ri.DB, ri.Redis)
  authRouterInstance.SetupAuthRouter()

	claimFolderRouterInstance := claimFolderRouter.NewRouterAuthInsurance(ri.Router, ri.DB, ri.Redis)
	claimFolderRouterInstance.SetupClaimFolderRouter()

}