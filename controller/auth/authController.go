package authController

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	authModel "insurance/model/auth"

	redis "github.com/go-redis/redis/v8"

	"golang.org/x/crypto/bcrypt"

	jwt "github.com/golang-jwt/jwt/v4"
)


type AuthController struct {
    AuthModel *authModel.AuthModel
		Redis *redis.Client
}

type UserRequest struct {
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	PhoneNumber string `json:"phoneNumber"`
	Token string `json:"token"`
}


type JWTClaims struct {
	PhoneNumber string `json:"phoneNumber"`
	jwt.RegisteredClaims
}

var jwtKey = []byte("MY SECRET KEY")


func NewAuthController(db *sql.DB, rdb *redis.Client) *AuthController {
		authModelInstance := authModel.NewAuthModel(db) 

    return &AuthController{
			AuthModel: authModelInstance, 
			Redis: rdb,
		}
}


func (ac *AuthController) LoginHandler(c *gin.Context)  {
		//get body request and validate
		var userRequest UserRequest

		if errBindRequest := c.BindJSON(&userRequest); errBindRequest != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errBindRequest.Error(),
			})
			return
		}

		// get user with current phoneNumber
		userWithPhoneNumber, errGetUserByPhoneNumber := ac.AuthModel.GetUserByPhoneNumber(userRequest.PhoneNumber);
		
		if errGetUserByPhoneNumber != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": errGetUserByPhoneNumber,
			})

			return
		}

		// Compare the password with the hashed password

		if errComparePassword := bcrypt.CompareHashAndPassword(
			[]byte(userWithPhoneNumber.Password),
			[]byte(userRequest.Password),
		); errComparePassword != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "password incorrect",
			})

			return
		}


		//Generate token if user login success
		 expirationTime := time.Now().Add(24 * time.Hour)
		 jwtClaims := &JWTClaims{
			PhoneNumber: userRequest.PhoneNumber,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expirationTime),
			},
		 }

		 token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
		 tokenString, errGenToken := token.SignedString(jwtKey)

		 if errGenToken != nil {
			 c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			 return
		 }

  

    c.JSON(http.StatusOK, gin.H{"result": UserResponse{
			PhoneNumber: userRequest.PhoneNumber,
			Token: tokenString,
		}})
}

func (ac* AuthController) TestMiddleware(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}