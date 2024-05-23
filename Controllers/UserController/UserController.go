package UserController

import (
	"database/sql"
	"fmt"
	"net/http"

	"gymfinity-backend-api/Entities"
	"gymfinity-backend-api/Library"
	"gymfinity-backend-api/Models/UserModel"
	"gymfinity-backend-api/Validator"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Index(c *gin.Context) {
	roleParams := c.Query("role");

	users, err := UserModel.GetAll(&roleParams);

	if err != nil {
		if err == sql.ErrNoRows {
			Library.ApiResponseSuccess(c, http.StatusOK, fmt.Sprintf("No users found with role %v", roleParams), nil);
		} else {
			Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		}
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

func Create(c *gin.Context) {
	var userData Entities.User;

	if err := c.BindJSON(&userData); err != nil {
		Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		return;
	}

	validate := validator.New();
	Validator.RegisterCustomValidators(validate);
	
	if err := validate.Struct(userData); err != nil {
		errors := err.(validator.ValidationErrors);
		Library.ApiResponseError(c, http.StatusBadRequest, fmt.Sprintf("%v", errors));
		return;
	}

	if err := UserModel.Create(&userData); err != nil {
		Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		return;
	}

	Library.ApiResponseSuccess(c, http.StatusCreated, "A new user has been created", userData);
}

func Update(c *gin.Context) {
	userId := Library.ParseInt(c.Param("id"));
	var updatedData Entities.User;
	
	if err := c.BindJSON(&updatedData); err != nil {
		Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		return;
	}

	validate := validator.New();
	Validator.RegisterCustomValidators(validate);

	if err := validate.Struct(updatedData); err != nil {
		errors := err.(validator.ValidationErrors);
		Library.ApiResponseError(c, http.StatusBadRequest, fmt.Sprintf("%v", errors));
		return;
	}

	if err := UserModel.Update(&userId, &updatedData); err != nil {
		if err == sql.ErrNoRows {
			Library.ApiResponseSuccess(c, http.StatusOK, fmt.Sprintf("An user with id %d is not founded", userId), nil);
		} else {
			Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		}
		return;
	}

	Library.ApiResponseSuccess(c, http.StatusOK, fmt.Sprintf("An user with id %d has been updated", userId), updatedData);
}

func Delete(c *gin.Context) {
	userId := Library.ParseInt(c.Param("id"));

	err := UserModel.Delete(&userId);
	if err != nil {
		if err == sql.ErrNoRows {
			Library.ApiResponseSuccess(c, http.StatusOK, fmt.Sprintf("An user with id %d is not founded", userId), nil);
		} else {
			Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		}
		return;
	}

	Library.ApiResponseSuccess(c, http.StatusOK, fmt.Sprintf("An user with id %d has been removed", userId), nil);
}