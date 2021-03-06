// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/sign_in": {
            "post": {
                "description": "登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "description": "请求体参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.BodySignIn"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseSignIn"
                        }
                    }
                }
            }
        },
        "/sign_up": {
            "post": {
                "description": "注册",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "注册",
                "parameters": [
                    {
                        "description": "请求体参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.BodySignUP"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseCommon"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "用户列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "每页条目数",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "偏移量",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseUserInfoList"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "获取指定用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "获取指定用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseUserInfo"
                        }
                    }
                }
            },
            "post": {
                "description": "更新指定用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "更新指定用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "请求体参数",
                        "name": "object",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/models.BodyUpdateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseUserInfo"
                        }
                    }
                }
            },
            "delete": {
                "description": "删除指定用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "删除指定用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseCommon"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.BodySignIn": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                }
            }
        },
        "models.BodySignUP": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "gender": {
                    "description": "性别",
                    "type": "integer"
                },
                "nickname": {
                    "description": "昵称",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                }
            }
        },
        "models.BodyUpdateUser": {
            "type": "object",
            "properties": {
                "gender": {
                    "description": "性别",
                    "type": "integer"
                },
                "nickname": {
                    "description": "昵称",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                }
            }
        },
        "models.DataSignIn": {
            "type": "object",
            "properties": {
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "gender": {
                    "description": "性别",
                    "type": "integer"
                },
                "id": {
                    "description": "由于前端可能存在数字失真（2^53-1），故转为字符串",
                    "type": "string",
                    "example": "0"
                },
                "nickname": {
                    "description": "昵称",
                    "type": "string"
                },
                "token": {
                    "description": "JSON Web Token",
                    "type": "string"
                }
            }
        },
        "models.DataUserInfoList": {
            "type": "object",
            "properties": {
                "limit": {
                    "description": "每页条目数",
                    "type": "integer"
                },
                "offset": {
                    "description": "偏移量",
                    "type": "integer"
                },
                "total": {
                    "description": "总数",
                    "type": "integer"
                },
                "users": {
                    "description": "用户列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.UserInfo"
                    }
                }
            }
        },
        "models.ResponseCommon": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "内部状态码",
                    "type": "integer"
                },
                "message": {
                    "description": "消息提示",
                    "type": "string"
                }
            }
        },
        "models.ResponseSignIn": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "内部状态码",
                    "type": "integer"
                },
                "data": {
                    "description": "数据",
                    "$ref": "#/definitions/models.DataSignIn"
                },
                "message": {
                    "description": "消息提示",
                    "type": "string"
                }
            }
        },
        "models.ResponseUserInfo": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "内部状态码",
                    "type": "integer"
                },
                "data": {
                    "description": "数据",
                    "$ref": "#/definitions/models.UserInfo"
                },
                "message": {
                    "description": "消息提示",
                    "type": "string"
                }
            }
        },
        "models.ResponseUserInfoList": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "内部状态码",
                    "type": "integer"
                },
                "data": {
                    "description": "数据",
                    "$ref": "#/definitions/models.DataUserInfoList"
                },
                "message": {
                    "description": "消息提示",
                    "type": "string"
                }
            }
        },
        "models.UserInfo": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "创建时间",
                    "type": "string"
                },
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "gender": {
                    "description": "性别",
                    "type": "integer"
                },
                "id": {
                    "description": "用户ID，由于前端JS可能存在数字失真，故序列化时转为字符串",
                    "type": "string",
                    "example": "0"
                },
                "nickname": {
                    "description": "昵称",
                    "type": "string"
                },
                "updated_at": {
                    "description": "更新时间",
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
