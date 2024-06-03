package claimFoldeRouter

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	redis "github.com/go-redis/redis/v8"

	claimFolderController "insurance/controller/claimFolderController"

	middleware "insurance/middleware"
)


type RouterClaimFolder struct {
	Router *gin.Engine
	DB *sql.DB
	Redis *redis.Client
}

func NewRouterAuthInsurance(router *gin.Engine, db *sql.DB, rdb *redis.Client) *RouterClaimFolder {
	return &RouterClaimFolder{
		Router: router,
		DB: db,
		Redis: rdb,
	}
}

func (rcf *RouterClaimFolder) SetupClaimFolderRouter() {
	claimFolderControllerInstance := claimFolderController.NewClaimFolderController(rcf.DB, rcf.Redis) 


	claimFolderRouter := rcf.Router.Group("/claimfolders")

	claimFolderRouter.Use(middleware.AuthMiddleware())

	{
		claimFolderRouter.POST("/", claimFolderControllerInstance.AddNewClaimFolder)
	}
}