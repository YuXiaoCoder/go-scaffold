package controllers

import "go.uber.org/zap"

type StatusCode = int

const (
	CodeSuccess           StatusCode = 1000 // 成功
	CodeServerBusy        StatusCode = 2000 // 服务繁忙
	CodeInvalidParam      StatusCode = 2001 // 无效的参数
	CodeInterfaceNotExist StatusCode = 2002 // 不存在的接口
	CodeUserExist         StatusCode = 2100 // 用户已存在
	CodeInvalidUser       StatusCode = 2101 // 无效的用户信息
	CodeNeedAuth          StatusCode = 2102 // 需要认证信息
	CodeInvalidAuth       StatusCode = 2103 // 无效的认证信息
	CodeRateLimit         StatusCode = 2014 // 触发限速
	CodeHostExist         StatusCode = 2200 // 主机已存在
)

var messages = map[StatusCode]string{
	CodeSuccess:           "请求成功",
	CodeServerBusy:        "服务繁忙",
	CodeInvalidParam:      "无效的参数",
	CodeInterfaceNotExist: "不存在的接口",
	CodeUserExist:         "用户已存在",
	CodeInvalidUser:       "无效的用户信息",
	CodeNeedAuth:          "需要认证信息",
	CodeInvalidAuth:       "无效的认证信息",
	CodeRateLimit:         "触发限速",
	CodeHostExist:         "主机已存在",
}

// 获取信息提示
func GetMessage(code StatusCode) string {
	message, ok := messages[code]
	if !ok {
		zap.L().Error("invalid status code", zap.Int("code", code))
		message = messages[CodeServerBusy]
	}
	return message
}
