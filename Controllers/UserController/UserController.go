package UserController

import (
	"gymfinity-backend-api/Library"
	"gymfinity-backend-api/Models/UserModel"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	roleParams := c.Query("role")

	users, err := UserModel.GetAll(&roleParams);

	if err != nil {
		Library.ApiResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	Library.ApiResponseSuccess(c, http.StatusOK, users)
}