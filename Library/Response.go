package Library 

import "github.com/gin-gonic/gin"

type SuccessResponseFormat struct {
	StatusCode	int			`json:"statusCode"`
	Message		string		`json:"message"`
	Data		interface{}	`json:"data"`
}

type ErrorResponseFormat struct {
	StatusCode	int			`json:"statusCode"`
	Error		string		`json:"error"`
}

func ApiResponseSuccess(ctx *gin.Context, httpStatusCode int, message string, data interface{}) {
	ctx.JSON(httpStatusCode, SuccessResponseFormat{StatusCode: httpStatusCode, Message: message, Data: data});
}

func ApiResponseError(ctx *gin.Context, httpStatusCode int, err string) {
	ctx.JSON(httpStatusCode, ErrorResponseFormat{StatusCode: httpStatusCode, Error: err});
}