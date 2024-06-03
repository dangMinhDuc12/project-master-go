package claimImageModel

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)


type ClaimImageModel struct {
	DB *sql.DB
}

func NewClaimImageModel (db *sql.DB) *ClaimImageModel {
	return &ClaimImageModel{
		DB: db,
	}
}

type AddNewClaimImageRequest struct {
	ClaimId int `json:"claimId" binding:"required"`
	FilePath string `json:"filePath" binding:"required"`
}

func (cim *ClaimImageModel) AddNewClaimImage(claimImageInfo *AddNewClaimImageRequest) {
	fmt.Println("claimImageInfo", claimImageInfo)
}