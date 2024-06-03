package claimFolderModel

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type ClaimFolderModel struct {
	DB *sql.DB
}

type AddNewClaimFolderRequest struct {
	ClaimName string `json:"claimName" binding:"required"`
	IsClaim bool `json:"isClaim"`
}

type ClaimFolder struct {
	ClaimId int `json:"claimId"`
	ClaimName string `json:"claimName"`
	IsClaim bool `json:"isClaim"`
}

func NewClaimFolderModel(db *sql.DB) *ClaimFolderModel {
    return &ClaimFolderModel{DB: db}
}

func (cfm *ClaimFolderModel) AddNewClaimFolder(claimFolderInfo AddNewClaimFolderRequest) (ClaimFolder, error) {

	var claimFolder ClaimFolder

	sqlQuery := `
		INSERT INTO public.insurance_claims (claim_name, is_claim)
		VALUES ($1, $2)
		RETURNING claim_id, claim_name, is_claim
	`

	errQuery := cfm.DB.QueryRow(sqlQuery, claimFolderInfo.ClaimName, claimFolderInfo.IsClaim).Scan(&claimFolder.ClaimId, &claimFolder.ClaimName, &claimFolder.IsClaim)

	if errQuery != nil {
		return ClaimFolder{}, errQuery
	}

	return claimFolder, nil;
}