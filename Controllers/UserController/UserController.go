package UserController

import (
	"database/sql"
	"fmt"
	"net/http"

	"gymfinity-backend-api/Library"
	"gymfinity-backend-api/Models/UserModel"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	roleParams := c.Query("role");

	users, err := UserModel.GetAll(&roleParams);

	if err != nil {
		Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		return;
	}

	Library.ApiResponseSuccess(c, http.StatusOK, "Successfully fetch all user's datas", users);
}

func Show(c *gin.Context) {
	userId := Library.ParseInt(c.Param("id"));

	user, err := UserModel.GetById(&userId);
	if err != nil {
		if err == sql.ErrNoRows {
			Library.ApiResponseSuccess(c, http.StatusOK, fmt.Sprintf("No user found with id %d", userId), nil);
		} else {
			Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		}

		return;
	}

	Library.ApiResponseSuccess(c, http.StatusOK, "Successfully fetch user's data", user);
}