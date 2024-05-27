package authController

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	authModel "insurance/model/auth"
)


type AuthController struct {
    AuthModel *authModel.AuthModel
}

func NewAuthController(db *sql.DB) *AuthController {
		authModelInstance := authModel.NewAuthModel(db) 

    return &AuthController{AuthModel: authModelInstance}
}


func (ac *AuthController) LoginHandler(c *gin.Context)  {
    // Handle login logic
		user, errGetAllUser := ac.AuthModel.GetAllUser()

	

		if errGetAllUser != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server errorr"})
			return
		}

    c.JSON(http.StatusOK, gin.H{"user": user})
}