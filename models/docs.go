package models

// Swagger API 文档

// ResponseCommon 注册的响应
type ResponseCommon struct {
	Code    int    `json:"code"`    // 内部状态码
	Message string `json:"message"` // 消息提示
}

// ResponseSignIn 登录的响应
type ResponseSignIn struct {
	Code    int        `json:"code"`    // 内部状态码
	Message string     `json:"message"` // 消息提示
	Data    DataSignIn `json:"data"`    // 数据
}

// QueryUserInfoList 用户列表的查询参数
type QueryUserInfoList struct {
	Limit  int `json:"limit" default:"10"` // 每页条目数
	Offset int `json:"offset" default:"0"` // 偏移量
}

// ResponseUserInfoList 用户列表的响应
type ResponseUserInfoList struct {
	Code    int              `json:"code"`    // 内部状态码
	Message string           `json:"message"` // 消息提示
	Data    DataUserInfoList `json:"data"`    // 数据
}

// ResponseUserInfo 获取指定用户的响应
type ResponseUserInfo struct {
	Code    int      `json:"code"`    // 内部状态码
	Message string   `json:"message"` // 消息提示
	Data    UserInfo `json:"data"`    // 数据
}

// BodyUpdateUser 更新用户的请求体参数
type BodyUpdateUser struct {
	Password string `json:"password"` // 密码
	Nickname string `json:"nickname"` // 昵称
	Gender   int    `json:"gender"`   // 性别
}
