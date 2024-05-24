package ClassController

import (
	"database/sql"
	"fmt"
	"net/http"

	"gymfinity-backend-api/Entities"
	"gymfinity-backend-api/Library"
	"gymfinity-backend-api/Models/ClassModel"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Index(c *gin.Context) {
	classes, err := ClassModel.GetAll();
	if err != nil {
		if err == sql.ErrNoRows {
			Library.ApiResponseSuccess(c, http.StatusOK, "No classes found", nil);
		} else {
			Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		}
		return;
	}

	Library.ApiResponseSuccess(c, http.StatusOK, "Successfully get all class datas", classes);
}

func Show(c *gin.Context) {
	classId := Library.ParseInt(c.Param("id"));
	class, err := ClassModel.GetById(&classId);

	if err != nil {
		if err == sql.ErrNoRows {
			Library.ApiResponseSuccess(c, http.StatusOK, fmt.Sprintf("No class found with id %d", classId), nil);
		} else {
			Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		}
		return;
	}

	Library.ApiResponseSuccess(c, http.StatusOK, "Successfully get class data", class);
} 

func Create(c *gin.Context) {
	var classData Entities.Class;

	if err := c.BindJSON(&classData); err != nil {
		Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		return;
	}

	validate := validator.New();
	if err := validate.Struct(classData); err != nil {
		errors := err.(validator.ValidationErrors);
		Library.ApiResponseError(c, http.StatusBadRequest, fmt.Sprintf("%v", errors));
		return;
	}

	if err := ClassModel.Create(&classData); err != nil {
		Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		return;
	}

	Library.ApiResponseSuccess(c, http.StatusCreated, "A new class has been created", classData);
}

func Update(c *gin.Context) {
	classId := Library.ParseInt(c.Param("id"));
	var updatedData Entities.Class;

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

	if err := ClassModel.Update(&classId, &updatedData); err != nil {
		if err == sql.ErrNoRows {
			Library.ApiResponseSuccess(c, http.StatusOK, fmt.Sprintf("A class with id %d is not found", classId), nil);
		} else {
			Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		}
		return;
	}

	Library.ApiResponseSuccess(c, http.StatusOK, fmt.Sprintf("A class with id %d has been updated", classId), updatedData);
}

func Delete(c *gin.Context) {
	classId := Library.ParseInt(c.Param("id"));
	err := ClassModel.Delete(&classId);
	if err != nil {
		if err == sql.ErrNoRows {
			Library.ApiResponseSuccess(c, http.StatusOK, fmt.Sprintf("No class found with id %d", classId), nil);
		} else {
			Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		}
		return; 
	}

	Library.ApiResponseSuccess(c, http.StatusOK, fmt.Sprintf("A class with id %d has been removed", classId), nil);
}