package controllers

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	CTXUserKey = "userID"
)

// getCurrentUserID 获取当前请求上下文中的用户ID
func getCurrentUserID(ctx *gin.Context) (userID int64, err error) {
	// 从上下文获取
	value, ok := ctx.Get(CTXUserKey)
	if !ok {
		return 0, errors.New(GetMessage(CodeNeedAuth))
	}

	// 验证类型
	userID, ok = value.(int64)
	if !ok {
		return 0, errors.New(GetMessage(CodeNeedAuth))
	}
	return userID, err
}

// getParam 获取参数
func getParam(ctx *gin.Context, key string) string {
	// 优先获取路径参数
	value := ctx.Param(key)
	if value == "" {
		value = ctx.Query(key)
	}
	return strings.TrimSpace(value)
}

// getStrParam 获取字符串参数
func getStrParam(ctx *gin.Context, key string) (string, error) {
	value := getParam(ctx, key)
	if value == "" {
		return value, errors.New(fmt.Sprintf("The [%s] was not found in the parameters", key))
	}
	return value, nil
}

// getInt64Param 获取Int64参数
func getInt64Param(ctx *gin.Context, key string) (int64, error) {
	value, err := getStrParam(ctx, key)
	if err != nil {
		return 0, err
	}

	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, err
	}
	return intValue, nil
}
