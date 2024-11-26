package base

import "github.com/gin-gonic/gin"

type baseResponse struct {
	IsSuccess bool        `json:"isSuccess"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Error     interface{} `json:"error,omitempty"`
}

func SuccessResponse(ctx *gin.Context, statusCode int, message string, data interface{}) {
	response := baseResponse{
		IsSuccess: true,
		Message:   message,
		Data:      data,
	}
	ctx.JSON(statusCode, response)
}

func ErrorResponse(ctx *gin.Context, statusCode int, message string, err interface{}) {
	response := baseResponse{
		IsSuccess: false,
		Message:   message,
		Error:     err,
	}
	ctx.JSON(statusCode, response)
}
