// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/addapi": {
            "post": {
                "description": "参数校验慢慢完善",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api记录操作"
                ],
                "summary": "创建api",
                "parameters": [
                    {
                        "description": "api info",
                        "name": "UserLogin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structtypes.Addapireq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/structtypes.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/chacklowversion": {
            "get": {
                "description": "比较客户端与服务器的版本号",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "客户端信息"
                ],
                "summary": "获取客户端最低版本",
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/structtypes.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/createproject": {
            "post": {
                "description": "目前暂时没有做权限",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "项目操作"
                ],
                "summary": "创建项目",
                "parameters": [
                    {
                        "description": "project info",
                        "name": "UserLogin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structtypes.Createproject"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/structtypes.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/delapi": {
            "post": {
                "description": "没有说明",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api记录操作"
                ],
                "summary": "删除api",
                "parameters": [
                    {
                        "description": "api info",
                        "name": "UserLogin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structtypes.Delapi"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/structtypes.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/delproject": {
            "post": {
                "description": "只要创建者才能删除",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "项目操作"
                ],
                "summary": "删除项目",
                "parameters": [
                    {
                        "description": "project info",
                        "name": "UserLogin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structtypes.Delproject"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/structtypes.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/getapibyid": {
            "post": {
                "description": "目前没有需要注意的",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api记录操作"
                ],
                "summary": "通过id获取api",
                "parameters": [
                    {
                        "description": "api info",
                        "name": "UserLogin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structtypes.Getapi"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/structtypes.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/structtypes.Getapires"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/getapilist": {
            "post": {
                "description": "目前没有需要注意的",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api记录操作"
                ],
                "summary": "获取api列表",
                "parameters": [
                    {
                        "description": "api info",
                        "name": "UserLogin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structtypes.Getapilist"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/structtypes.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/structtypes.Getapilistres"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/getpermissionlist": {
            "post": {
                "description": "需要管理员权限",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "权限操作"
                ],
                "summary": "获取项目权限表",
                "parameters": [
                    {
                        "description": "project info",
                        "name": "UserLogin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structtypes.Getpermissionlist"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/structtypes.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/structtypes.Getpermissionlistres"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/getprojectlist": {
            "post": {
                "description": "项目选择页面用的",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "项目操作"
                ],
                "summary": "查询项目列表",
                "parameters": [
                    {
                        "description": "api info",
                        "name": "UserLogin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structtypes.GetProjectlist"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/structtypes.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/structtypes.Getprojectlistres"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/setpermission": {
            "post": {
                "description": "需要管理员权限",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "权限操作"
                ],
                "summary": "设置权限",
                "parameters": [
                    {
                        "description": "project info",
                        "name": "UserLogin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structtypes.Setpermission"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/structtypes.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/userlogin": {
            "post": {
                "description": "用户登陆api",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户操作"
                ],
                "summary": "用户登陆",
                "parameters": [
                    {
                        "description": "user info",
                        "name": "UserLogin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structtypes.UserLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/structtypes.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/userreg": {
            "post": {
                "description": "用户注册api",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户操作"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "user info",
                        "name": "UserRegReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structtypes.UserRegReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/structtypes.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/usersearch": {
            "get": {
                "description": "无需鉴权",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户操作"
                ],
                "summary": "模糊查询用户",
                "parameters": [
                    {
                        "type": "string",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/structtypes.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/structtypes.Usersearchres"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "structtypes.Addapireq": {
            "type": "object",
            "required": [
                "bodyrecodes",
                "bodytype",
                "headcode",
                "isDeprecated",
                "path",
                "pathcode",
                "projectname",
                "returncodes",
                "returntype",
                "tokens",
                "types",
                "username",
                "words"
            ],
            "properties": {
                "bodyrecodes": {
                    "description": "body结构",
                    "type": "string"
                },
                "bodytype": {
                    "description": "body的类型",
                    "type": "string"
                },
                "headcode": {
                    "description": "请求头部结构",
                    "type": "string"
                },
                "isDeprecated": {
                    "description": "弃用状态",
                    "type": "integer",
                    "enum": [
                        1,
                        2
                    ]
                },
                "path": {
                    "description": "路径",
                    "type": "string"
                },
                "pathcode": {
                    "description": "path结构",
                    "type": "string"
                },
                "projectname": {
                    "description": "项目名",
                    "type": "string"
                },
                "returncodes": {
                    "description": "返回结构",
                    "type": "string"
                },
                "returntype": {
                    "description": "返回类型",
                    "type": "string"
                },
                "tokens": {
                    "description": "用户token",
                    "type": "string"
                },
                "types": {
                    "description": "请求类型",
                    "type": "string",
                    "enum": [
                        "get",
                        "post"
                    ]
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                },
                "words": {
                    "description": "注释",
                    "type": "string"
                }
            }
        },
        "structtypes.Createproject": {
            "type": "object",
            "required": [
                "name",
                "tokens",
                "username"
            ],
            "properties": {
                "name": {
                    "description": "名字",
                    "type": "string"
                },
                "tokens": {
                    "description": "用户token",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "structtypes.Delapi": {
            "type": "object",
            "required": [
                "apiid",
                "projectname",
                "tokens",
                "username"
            ],
            "properties": {
                "apiid": {
                    "type": "integer"
                },
                "projectname": {
                    "description": "项目名",
                    "type": "string"
                },
                "tokens": {
                    "description": "用户token",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "structtypes.Delproject": {
            "type": "object",
            "required": [
                "name",
                "tokens",
                "username"
            ],
            "properties": {
                "name": {
                    "description": "名字",
                    "type": "string"
                },
                "tokens": {
                    "description": "用户token",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "structtypes.GetProjectlist": {
            "type": "object",
            "required": [
                "tokens",
                "username"
            ],
            "properties": {
                "tokens": {
                    "description": "用户token",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "structtypes.Getapi": {
            "type": "object",
            "required": [
                "apiid",
                "projectname",
                "tokens",
                "username"
            ],
            "properties": {
                "apiid": {
                    "type": "integer"
                },
                "projectname": {
                    "description": "项目名",
                    "type": "string"
                },
                "tokens": {
                    "description": "用户token",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "structtypes.Getapilist": {
            "type": "object",
            "required": [
                "projectname",
                "tokens",
                "username"
            ],
            "properties": {
                "projectname": {
                    "description": "项目名",
                    "type": "string"
                },
                "tokens": {
                    "description": "用户token",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "structtypes.Getapilistres": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "id",
                    "type": "integer"
                },
                "isDeprecated": {
                    "description": "弃用状态 1表示正常 2表示弃用",
                    "type": "integer"
                },
                "path": {
                    "description": "路径",
                    "type": "string"
                },
                "types": {
                    "description": "请求类型",
                    "type": "string"
                }
            }
        },
        "structtypes.Getapires": {
            "type": "object",
            "properties": {
                "bodyrecodes": {
                    "description": "body结构",
                    "type": "string"
                },
                "bodytype": {
                    "description": "body的类型",
                    "type": "string"
                },
                "headcode": {
                    "description": "请求头部结构",
                    "type": "string"
                },
                "isDeprecated": {
                    "description": "弃用状态 1表示正常 2表示弃用",
                    "type": "integer"
                },
                "path": {
                    "description": "路径",
                    "type": "string"
                },
                "pathcode": {
                    "description": "path结构",
                    "type": "string"
                },
                "returncodes": {
                    "description": "返回结构",
                    "type": "string"
                },
                "returntype": {
                    "description": "返回类型",
                    "type": "string"
                },
                "types": {
                    "description": "请求类型",
                    "type": "string"
                },
                "words": {
                    "description": "注释",
                    "type": "string"
                }
            }
        },
        "structtypes.Getpermissionlist": {
            "type": "object",
            "required": [
                "projectname",
                "tokens",
                "username"
            ],
            "properties": {
                "projectname": {
                    "description": "项目名",
                    "type": "string"
                },
                "tokens": {
                    "description": "用户token",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "structtypes.Getpermissionlistres": {
            "type": "object",
            "properties": {
                "levels": {
                    "type": "integer"
                },
                "userID": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "structtypes.Getprojectlistres": {
            "type": "object",
            "properties": {
                "apinum": {
                    "type": "integer"
                },
                "creater": {
                    "type": "string"
                },
                "projectid": {
                    "type": "integer"
                },
                "projectname": {
                    "type": "string"
                }
            }
        },
        "structtypes.JSONResult": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "structtypes.Setpermission": {
            "type": "object",
            "required": [
                "projectname",
                "tokens",
                "username"
            ],
            "properties": {
                "adminnames": {
                    "description": "管理员",
                    "type": "string"
                },
                "editnames": {
                    "description": "编辑者",
                    "type": "string"
                },
                "guestnames": {
                    "description": "访客",
                    "type": "string"
                },
                "projectname": {
                    "description": "项目名",
                    "type": "string"
                },
                "tokens": {
                    "description": "用户token",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "structtypes.UserLogin": {
            "type": "object",
            "required": [
                "language",
                "password",
                "username"
            ],
            "properties": {
                "language": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "structtypes.UserRegReq": {
            "type": "object",
            "required": [
                "email",
                "language",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "language": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "structtypes.Usersearchres": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "getneko api doc",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
