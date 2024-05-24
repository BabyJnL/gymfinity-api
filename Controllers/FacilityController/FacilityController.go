package FacilityController

import (
	"database/sql"
	"fmt"

	"net/http"

	"gymfinity-backend-api/Entities"
	"gymfinity-backend-api/Library"
	"gymfinity-backend-api/Models/FacilityModel"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Index(c *gin.Context) {
	facilities, err := FacilityModel.GetAll();
	if err != nil {
		if err == sql.ErrNoRows {
			Library.ApiResponseSuccess(c, http.StatusOK, "No facilities found", nil);
		} else {
			Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		}
		return;
	}

	Library.ApiResponseSuccess(c, http.StatusOK, "Successfully get facility datas", facilities);
}

func Show(c *gin.Context) {
	facilityId := Library.ParseInt(c.Param("id"));

	facility, err := FacilityModel.GetById(&facilityId);
	if err != nil {
		if err == sql.ErrNoRows {
			Library.ApiResponseSuccess(c, http.StatusOK, fmt.Sprintf("No facility found with id %d", facilityId), nil);
		} else {
			Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		}
		return;
	}

	Library.ApiResponseSuccess(c, http.StatusOK, "Successfuly get facility data", facility);
}

func Create(c *gin.Context) {
	var facilityData Entities.Facility;

	if err := c.BindJSON(&facilityData); err != nil {
		Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		return;
	}

	validate := validator.New();
	if err := validate.Struct(facilityData); err != nil {
		errors := err.(validator.ValidationErrors);
		Library.ApiResponseError(c, http.StatusBadRequest, fmt.Sprintf("%v", errors));
		return;
	}

	if err := FacilityModel.Create(&facilityData); err != nil {
		Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		return;
	}

	Library.ApiResponseSuccess(c, http.StatusCreated, "A new facility has been created", facilityData);
}

func Update(c *gin.Context) {
	facilityId := Library.ParseInt(c.Param("id"));
	var updatedData Entities.Facility;

	if err := c.BindJSON(&updatedData); err != nil {
		Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		return;
	}

	validate := validator.New();
	if err := validate.Struct(updatedData); err != nil {
		errors := err.(validator.ValidationErrors);
		Library.ApiResponseError(c, http.StatusBadRequest, fmt.Sprintf("%v", errors));
		return;
	}

	if err := FacilityModel.Update(&facilityId, &updatedData); err != nil {
		if err == sql.ErrNoRows {
			Library.ApiResponseSuccess(c, http.StatusOK, fmt.Sprintf("A facility with id %d is not found", facilityId), nil);
		} else {
			Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		}
		return;
	}

	Library.ApiResponseSuccess(c, http.StatusOK, fmt.Sprintf("A class schedule with id %d has been updated", facilityId), updatedData);
}

func Delete(c *gin.Context) {
	facilityId := Library.ParseInt(c.Param("id"));

	err := FacilityModel.Delete(&facilityId);
	if err != nil {
		if err == sql.ErrNoRows {
			Library.ApiResponseSuccess(c, http.StatusOK, fmt.Sprintf("A facility with id %d is not founded", facilityId), nil);
		} else {
			Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		}
		return;
	}

	Library.ApiResponseSuccess(c, http.StatusOK, fmt.Sprintf("A facility with id %d has been removed", facilityId), nil);
}