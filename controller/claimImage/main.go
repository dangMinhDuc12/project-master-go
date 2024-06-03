package claimImageController

import (
	"database/sql"
	"fmt"

	redis "github.com/go-redis/redis/v8"

	claimImageModel "insurance/model/claimImage"

	"github.com/gin-gonic/gin"
)

type ClaimImageController struct {
	Redis *redis.Client
	ClaimImageModel * claimImageModel.ClaimImageModel
}

func NewClaimImageController(db *sql.DB, rdb *redis.Client) *ClaimImageController {
	claimImageModelInstance := claimImageModel.NewClaimImageModel(db)

	return &ClaimImageController{
		Redis: rdb,
		ClaimImageModel: claimImageModelInstance,
	}
}

func (cic *ClaimImageController) AddNewClaimImage(c *gin.Context) {
	fmt.Println("hello")
}