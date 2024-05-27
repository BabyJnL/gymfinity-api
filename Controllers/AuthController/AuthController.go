package AuthController

import (
	"fmt"
	"net/http"
	"time"

	"gymfinity-backend-api/Entities"
	"gymfinity-backend-api/Library"
	"gymfinity-backend-api/Middleware"
	"gymfinity-backend-api/Models/UserModel"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
)

func Verify(c *gin.Context) {
	var verifyData Entities.UserVerify;

	if err := c.BindJSON(&verifyData); err != nil {
		Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		return;
	}

	fmt.Println(verifyData);

	validate := validator.New();
	if err := validate.Struct(verifyData); err != nil {
		errors := err.(validator.ValidationErrors);
		Library.ApiResponseError(c, http.StatusBadRequest, fmt.Sprintf("%v", errors));
		return;
	}

	userExists := UserModel.Verify(&verifyData);
	fmt.Println(userExists);
	if  userExists == nil || userExists.UserID == 0 {
		Library.ApiResponseSuccess(c, http.StatusUnauthorized, "Invalid credentials", nil);
		return;
	}

	expirationTime := time.Now().Add(1 * time.Hour)
    claims := &jwt.StandardClaims{
        Subject:   verifyData.Email,
        ExpiresAt: expirationTime.Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(Middleware.JwtyKey);
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
        return
    }

	Library.ApiResponseSuccess(c, http.StatusOK, "Authenticated", gin.H{"token": tokenString, "data": userExists});
}