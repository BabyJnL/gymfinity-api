package Library 

import "github.com/gin-gonic/gin"

type ResponseFormat struct {
	StatusCode	int			`json:"statusCode"`
	Error		string		`json:"error"`
	Data		interface{}	`json:"data"`
}

func ApiResponseSuccess(ctx *gin.Context, httpStatusCode int, data interface{}) {
	ctx.JSON(httpStatusCode, ResponseFormat{StatusCode: httpStatusCode, Data: data});
}

func ApiResponseError(ctx *gin.Context, httpStatusCode int, err string) {
	ctx.JSON(httpStatusCode, ResponseFormat{StatusCode: httpStatusCode, Error: err});
}