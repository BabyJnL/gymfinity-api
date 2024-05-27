package UserController

import (
	"database/sql"
	"fmt"
	"net/http"

	"gymfinity-backend-api/Entities"
	"gymfinity-backend-api/Library"
	"gymfinity-backend-api/Models/UserModel"

	"github.com/gin-gonic/gin"
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

	// Menangkap teks dari formulir
	firstname := c.PostForm("firstname")
	lastname := c.PostForm("lastname")
	gender := c.PostForm("gender");
	address := c.PostForm("address")
	phoneNumber := c.PostForm("phoneNumber");
	email := c.PostForm("email")
	status := c.PostForm("status")
	validUntil := c.PostForm("validUntil")
	role := c.PostForm("role")


	// Menangkap file dari formulir
	file, err := c.FormFile("photoPath")
	if err != nil {
		Library.ApiResponseError(c, http.StatusBadRequest, err.Error());
		return
	}

	filename := Library.GenerateUniqueFileName(file.Filename)

	err = c.SaveUploadedFile(file, "uploads/"+filename)
	if err != nil {
		Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		return
	}	

	userData.Firstname = firstname;
	userData.Lastname = lastname;
	userData.Gender = gender;
	userData.Address = address;
	userData.PhoneNumber = phoneNumber;
	userData.Email = email;
	userData.Status = status;
	userData.ValidUntil = validUntil;
	userData.Role = role;
	userData.PhotoPath = filename

	// validate := validator.New();
	// if err := validate.Struct(userData); err != nil {
	// 	errors := err.(validator.ValidationErrors);
	// 	Library.ApiResponseError(c, http.StatusBadRequest, fmt.Sprintf("%v", errors));
	// 	return;
	// }

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