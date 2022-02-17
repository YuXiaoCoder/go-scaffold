package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

// User 用户
type User struct {
	BasicMode
	Email    string `gorm:"column:email;"`    // 邮箱
	Password string `gorm:"column:password;"` // 密码
	Nickname string `gorm:"column:nickname;"` // 昵称
	Gender   int8   `gorm:"column:gender;"`   // 性别
}

// TableName 表名
func (User) TableName() string {
	return "user"
}

// UserInfo 用户信息
type UserInfo struct {
	ID        int64     `json:"id,string"`  // 用户ID，由于前端JS可能存在数字失真，故序列化时转为字符串
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
	Email     string    `json:"email"`      // 邮箱
	Nickname  string    `json:"nickname"`   // 昵称
	Gender    int8      `json:"gender" `    // 性别
}

// BodySignUP 注册的请求体参数
type BodySignUP struct {
	Email    string `json:"email" binding:"required,email"` // 邮箱
	Password string `json:"password" binding:"required"`    // 密码
	Nickname string `json:"nickname"`                       // 昵称
	Gender   int8   `json:"gender" binding:"oneof=0 1 2 3"` // 性别
}

// BodySignIn 登录的请求体参数
type BodySignIn struct {
	Email    string `json:"email" binding:"required,email"` // 邮箱
	Password string `json:"password" binding:"required"`    // 密码
}

// DataSignIn 登录的数据参数
type DataSignIn struct {
	ID       int64  `json:"id,string"` // 由于前端可能存在数字失真（2^53-1），故转为字符串
	Email    string `json:"email"`     // 邮箱
	Nickname string `json:"nickname"`  // 昵称
	Token    string `json:"token"`     // JSON Web Token
	Gender   int8   `json:"gender"`    // 性别
}

// QueryUserInfoListNull 用户列表的查询参数
type QueryUserInfoListNull struct {
	Limit  null.Int `form:"limit"`  // 每页条目数
	Offset null.Int `form:"offset"` // 偏移量
}

// DataUserInfoList 用户列表的数据
type DataUserInfoList struct {
	Users  []*UserInfo `json:"users"`  // 用户列表
	Total  int64       `json:"total"`  // 总数
	Limit  int64       `json:"limit"`  // 每页条目数
	Offset int64       `json:"offset"` // 偏移量
}

// BodyUpdateUserNull 更新用户的请求体参数
type BodyUpdateUserNull struct {
	Password null.String `json:"password"`                       // 密码
	Nickname null.String `json:"nickname"`                       // 昵称
	Gender   null.Int    `json:"gender" binding:"oneof=0 1 2 3"` // 性别
}
