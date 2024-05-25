package MembershipTypeController 

import (
	"database/sql"
	"net/http"

	"gymfinity-backend-api/Library"
	"gymfinity-backend-api/Models/MembershipTypeModel"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	membershipTypes, err := MembershipTypeModel.GetAll();
	if err != nil {
		if err == sql.ErrNoRows {
			Library.ApiResponseSuccess(c, http.StatusOK, "No membership type datas found", nil);
		} else {
			Library.ApiResponseError(c, http.StatusOK, err.Error());
		}
		return;
	}

	Library.ApiResponseSuccess(c, http.StatusOK, "Successfully get membership type datas", membershipTypes);
}