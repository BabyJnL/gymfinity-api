package PaymentController

import (
	"database/sql"
	"fmt"
	"net/http"

	"gymfinity-backend-api/Entities"
	"gymfinity-backend-api/Library"
	"gymfinity-backend-api/Models/PaymentModel"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Index(c *gin.Context) {
	payments, err := PaymentModel.GetAll();
	if err != nil {
		if err == sql.ErrNoRows {
			Library.ApiResponseSuccess(c, http.StatusOK, "No payments found", nil);
		} else {
			Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		}
		return;
	}

	Library.ApiResponseSuccess(c, http.StatusOK, "Successfully get payment datas", payments);
}

func Show(c *gin.Context) {
	paymentId := Library.ParseInt(c.Param("id"));
	payment, err := PaymentModel.GetById(&paymentId);
	if err != nil {
		if err == sql.ErrNoRows {
			Library.ApiResponseSuccess(c, http.StatusOK, fmt.Sprintf("No payment found with id %d", paymentId), nil);
		} else {
			Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		}
		return;
	}

	Library.ApiResponseSuccess(c, http.StatusOK, "Successfully get payment data", payment);
}

func Create(c *gin.Context) {
	var paymentData Entities.CreatePayment;

	if err := c.BindJSON(&paymentData); err != nil {
		Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		return;
	}

	validate := validator.New();
	if err := validate.Struct(paymentData); err != nil {
		errors := err.(validator.ValidationErrors);
		Library.ApiResponseError(c, http.StatusBadRequest, fmt.Sprintf("%v", errors));
		return;
	}

	if err := PaymentModel.Create(&paymentData); err != nil {
		Library.ApiResponseError(c, http.StatusInternalServerError, err.Error());
		return;
	}

	Library.ApiResponseSuccess(c, http.StatusCreated, "A new payment has been created", paymentData);
}