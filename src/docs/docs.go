// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "huoyijie",
            "url": "https://huoyijie.cn",
            "email": "yijie.huo@foxmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/apps": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Show apps",
                "tags": [
                    "app"
                ],
                "summary": "Show apps",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.App"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/ext.ClientError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/ext.ClientError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/dashboard": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Show dashboard",
                "tags": [
                    "dashboard"
                ],
                "summary": "Show dashboard",
                "parameters": [
                    {
                        "enum": [
                            "staff",
                            "user"
                        ],
                        "type": "string",
                        "description": "staff",
                        "name": "category",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin.dashboardRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ext.ClientError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/ext.ClientError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/ext.ClientError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/token/grant": {
            "post": {
                "description": "Generate token",
                "tags": [
                    "token"
                ],
                "summary": "Generate token",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "grantReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/token.GrantReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/token.GrantRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ext.ClientError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/token/refresh": {
            "post": {
                "description": "Refresh token",
                "tags": [
                    "token"
                ],
                "summary": "Refresh token",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "refreshReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/token.refreshReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/token.GrantRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ext.ClientError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/token/revoke": {
            "post": {
                "description": "Revoke token",
                "tags": [
                    "token"
                ],
                "summary": "Revoke token",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "refreshReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/token.refreshReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ext.ClientError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "admin.dashboardRes": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                }
            }
        },
        "ext.ClientError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "model.App": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "boolean"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "secret": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "token.GrantReq": {
            "type": "object",
            "required": [
                "appId",
                "appSecret",
                "email",
                "password"
            ],
            "properties": {
                "appId": {
                    "type": "integer",
                    "minimum": 1,
                    "example": 100000
                },
                "appSecret": {
                    "type": "string",
                    "minLength": 1,
                    "example": "123456"
                },
                "email": {
                    "type": "string",
                    "example": "admin@huoyijie.cn"
                },
                "password": {
                    "type": "string",
                    "minLength": 6,
                    "example": "123456"
                }
            }
        },
        "token.GrantRes": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "expires_in": {
                    "type": "integer"
                },
                "refresh_token": {
                    "type": "string"
                },
                "token_type": {
                    "type": "string"
                }
            }
        },
        "token.refreshReq": {
            "type": "object",
            "required": [
                "access_token",
                "refresh_token"
            ],
            "properties": {
                "access_token": {
                    "type": "string",
                    "minLength": 1
                },
                "refresh_token": {
                    "type": "string",
                    "minLength": 1
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Ink API",
	Description:      "This is ink server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
