package ClassScheduleController

import (
	"database/sql"
	"fmt"

	//"fmt"
	"net/http"

	"gymfinity-backend-api/Entities"
	"gymfinity-backend-api/Library"
	"gymfinity-backend-api/Models/ClassScheduleModel"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Index(c *gin.Context) {
	classSchedules, err := ClassScheduleModel.GetAll();
	if err != nil {
		if err == sql.ErrNoRows {
			Library.ApiResponseSuccess(c, http.StatusOK, "No class schedules found", nil);
		} else {
			Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		}
		return;
	}

	Library.ApiResponseSuccess(c, http.StatusOK, "Successfully get all class schedules", classSchedules);
}

func Show(c *gin.Context) {
	scheduleId := Library.ParseInt(c.Param("id"));

	classSchedule, err := ClassScheduleModel.GetById(&scheduleId);
	if err != nil {
		if err == sql.ErrNoRows {
			Library.ApiResponseSuccess(c, http.StatusOK, fmt.Sprintf("No class schedule found with id %d", scheduleId), nil);
		} else {
			Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		}
		return;
	}

	Library.ApiResponseSuccess(c, http.StatusOK, "Successfuly get class schedule data", classSchedule);
}

func Create(c *gin.Context) {
	var classScheduleData Entities.ClassSchedule;

	if err := c.BindJSON(&classScheduleData); err != nil {
		Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		return;
	}

	validate := validator.New();
	if err := validate.Struct(classScheduleData); err != nil {
		errors := err.(validator.ValidationErrors);
		Library.ApiResponseError(c, http.StatusBadRequest, fmt.Sprintf("%v", errors));
		return;
	}

	if err := ClassScheduleModel.Create(&classScheduleData); err != nil {
		Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		return;
	}

	Library.ApiResponseSuccess(c, http.StatusCreated, "A new class schedule has been created", classScheduleData);
}

func Update(c *gin.Context) {
	scheduleId := Library.ParseInt(c.Param("id"));
	var updatedData Entities.ClassSchedule;

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

	if err := ClassScheduleModel.Update(&scheduleId, &updatedData); err != nil {
		if err == sql.ErrNoRows {
			Library.ApiResponseSuccess(c, http.StatusOK, fmt.Sprintf("A class schedule with id %d is not found", scheduleId), nil);
		} else {
			Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		}
		return;
	}

	Library.ApiResponseSuccess(c, http.StatusOK, fmt.Sprintf("A class schedule with id %d has been updated", scheduleId), updatedData);
}

func Delete(c *gin.Context) {
	classScheduleId := Library.ParseInt(c.Param("id"));

	err := ClassScheduleModel.Delete(&classScheduleId);
	if err != nil {
		if err == sql.ErrNoRows {
			Library.ApiResponseSuccess(c, http.StatusOK, fmt.Sprintf("A class schedule with id %d is not founded", classScheduleId), nil);
		} else {
			Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		}
		return;
	}

	Library.ApiResponseSuccess(c, http.StatusOK, fmt.Sprintf("A class schedule with id %d has been removed", classScheduleId), nil);
}