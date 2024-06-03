package claimFolderController

import (
	"database/sql"
	claimFolderModel "insurance/model/claimFolder"
	"net/http"

	redis "github.com/go-redis/redis/v8"

	"github.com/gin-gonic/gin"
)

type ClaimFolderController struct {
		Redis *redis.Client
		ClaimFolderModel *claimFolderModel.ClaimFolderModel
}

type AddNewClaimFolderResponse struct {
	ClaimId int `json:"claimId"`
}


func NewClaimFolderController(db *sql.DB, rdb *redis.Client) *ClaimFolderController {
		claimFolderModelInstance := claimFolderModel.NewClaimFolderModel(db) 

    return &ClaimFolderController{
			ClaimFolderModel: claimFolderModelInstance, 
			Redis: rdb,
		}
}

func (cfc *ClaimFolderController) AddNewClaimFolder(c *gin.Context) {
	//get body request and validate
	var addNewClaimFolderReq claimFolderModel.AddNewClaimFolderRequest


	if errBindRequest := c.ShouldBindJSON(&addNewClaimFolderReq); errBindRequest != nil {
		c.JSON(http.StatusBadRequest, gin.H{
				"error": errBindRequest.Error(),
			})
			return
	}

	
	addedClaimFolder, errAddClaimFolder := cfc.ClaimFolderModel.AddNewClaimFolder(addNewClaimFolderReq)

	if errAddClaimFolder != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
		"error": errAddClaimFolder,
		})
	}


	c.JSON(http.StatusOK, gin.H{
		"result": &AddNewClaimFolderResponse{
			ClaimId: addedClaimFolder.ClaimId,
		},
	})
}