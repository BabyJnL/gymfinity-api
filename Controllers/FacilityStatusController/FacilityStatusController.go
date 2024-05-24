package FacilityStatusController

import (
	"database/sql"
	"net/http"

	"gymfinity-backend-api/Library"
	"gymfinity-backend-api/Models/FacilityStatusModel"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	facilityStatuses, err := FacilityStatusModel.GetAll();
	if err != nil {
		if err == sql.ErrNoRows {
			Library.ApiResponseSuccess(c, http.StatusOK, "No facility statuses found", nil);
		} else {
			Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		}
		return;
	}

	Library.ApiResponseSuccess(c, http.StatusOK, "Successfully get facility statuses", facilityStatuses);
}