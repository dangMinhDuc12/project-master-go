package claimImageRouter

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	redis "github.com/go-redis/redis/v8"

	claimImageController "insurance/controller/claimImage"

	middleware "insurance/middleware"
)


type RouterClaimImage struct {
	Router *gin.Engine
	DB *sql.DB
	Redis *redis.Client
}

func NewRouterClaimImage(router *gin.Engine, db *sql.DB, rdb *redis.Client) *RouterClaimImage {
	return &RouterClaimImage{
		Router: router,
		DB: db,
		Redis: rdb,
	}
}


func (rci *RouterClaimImage) SetupClaimImageRouter() {
	claimImageControllerInstance := claimImageController.NewClaimImageController(rci.DB, rci.Redis)

	claimImageRouter := rci.Router.Group("/claimimages")

	claimImageRouter.Use(middleware.AuthMiddleware())

	{
		claimImageRouter.POST("/triton-assessment", claimImageControllerInstance.AddNewClaimImage)
	}
}