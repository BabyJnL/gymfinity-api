package UserStatusController

import (
	"database/sql"
	"net/http"

	"gymfinity-backend-api/Library"
	"gymfinity-backend-api/Models/UserStatusModel"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	userStatuses, err := UserStatusModel.GetAll();
	if err != nil {
		if err == sql.ErrNoRows {
			Library.ApiResponseSuccess(c, http.StatusOK, "No user status datas found", nil);
		} else {
			Library.ApiResponseError(c, http.StatusOK, err.Error());
		}
		return;
	}

	Library.ApiResponseSuccess(c, http.StatusOK, "Successfully get user statuses", userStatuses);
}