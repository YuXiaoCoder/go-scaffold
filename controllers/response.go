package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    StatusCode  `json:"code"`           // 内部状态码
	Message interface{} `json:"message"`        // 消息提示
	Data    interface{} `json:"data,omitempty"` // 数据
}

// 请求成功的响应
func ResponseSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, &Response{
		Code:    CodeSuccess,
		Message: GetMessage(CodeSuccess),
		Data:    data,
	})
}

// 请求失败的响应
func ResponseError(ctx *gin.Context, code StatusCode) {
	ctx.JSON(http.StatusOK, &Response{
		Code:    code,
		Message: GetMessage(code),
		Data:    nil,
	})
}

// 请求失败的响应（自定义消息）
func ResponseErrorWithMessage(ctx *gin.Context, code StatusCode, message interface{}) {
	ctx.JSON(http.StatusOK, &Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
