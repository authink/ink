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
                    "admin_app"
                ],
                "summary": "Show apps",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/admin.appRes"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/inkstone.ClientError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/inkstone.ClientError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Add a app",
                "tags": [
                    "admin_app"
                ],
                "summary": "Add a app",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "addAppReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/admin.addAppReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin.appRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/inkstone.ClientError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/inkstone.ClientError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/inkstone.ClientError"
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
        "/admin/apps/{id}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update a app",
                "tags": [
                    "admin_app"
                ],
                "summary": "Update a app",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "app id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "request body",
                        "name": "updateAppReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/admin.updateAppReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin.appRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/inkstone.ClientError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/inkstone.ClientError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/inkstone.ClientError"
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
                            "$ref": "#/definitions/inkstone.ClientError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/inkstone.ClientError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/inkstone.ClientError"
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
        "/admin/groups": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Show groups",
                "tags": [
                    "admin_group"
                ],
                "summary": "Show groups",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "type",
                        "name": "type",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "appId",
                        "name": "appId",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/inkstone.PagingResponse-admin_groupRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/inkstone.ClientError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/inkstone.ClientError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/inkstone.ClientError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Add a group",
                "tags": [
                    "admin_group"
                ],
                "summary": "Add a group",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "addGroupReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/admin.addGroupReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin.groupRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/inkstone.ClientError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/inkstone.ClientError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/inkstone.ClientError"
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
        "/admin/staffs": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Show staffs",
                "tags": [
                    "admin_staff"
                ],
                "summary": "Show staffs",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/inkstone.PagingResponse-admin_staffRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/inkstone.ClientError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/inkstone.ClientError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/inkstone.ClientError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Add a staff",
                "tags": [
                    "admin_staff"
                ],
                "summary": "Add a staff",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "addStaffReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/admin.addStaffReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin.staffRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/inkstone.ClientError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/inkstone.ClientError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/inkstone.ClientError"
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
        "/admin/staffs/{id}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update a staff",
                "tags": [
                    "admin_staff"
                ],
                "summary": "Update a staff",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "staff id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "request body",
                        "name": "updateStaffReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/admin.updateStaffReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin.staffRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/inkstone.ClientError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/inkstone.ClientError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/inkstone.ClientError"
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
        "/admin/tokens": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Show tokens",
                "tags": [
                    "admin_token"
                ],
                "summary": "Show tokens",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/inkstone.PagingResponse-admin_tokenRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/inkstone.ClientError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/inkstone.ClientError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/inkstone.ClientError"
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
        "/admin/tokens/{id}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete a token",
                "tags": [
                    "admin_token"
                ],
                "summary": "Delete a token",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "token id",
                        "name": "id",
                        "in": "path",
                        "required": true
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
                            "$ref": "#/definitions/inkstone.ClientError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/inkstone.ClientError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/inkstone.ClientError"
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
                        "type": "string",
                        "description": "language",
                        "name": "lang",
                        "in": "query"
                    },
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
                            "$ref": "#/definitions/inkstone.ClientError"
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
                            "$ref": "#/definitions/inkstone.ClientError"
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
                            "$ref": "#/definitions/inkstone.ClientError"
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
        "admin.addAppReq": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "minLength": 6,
                    "example": "appmock"
                }
            }
        },
        "admin.addGroupReq": {
            "type": "object",
            "required": [
                "appId",
                "name",
                "type"
            ],
            "properties": {
                "appId": {
                    "type": "integer",
                    "minimum": 100000,
                    "example": 100000
                },
                "name": {
                    "type": "string",
                    "minLength": 2,
                    "example": "developer"
                },
                "type": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "admin.addStaffReq": {
            "type": "object",
            "required": [
                "email",
                "phone"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "example@huoyijie.cn"
                },
                "phone": {
                    "type": "string",
                    "maxLength": 11,
                    "minLength": 11,
                    "example": "18555201314"
                }
            }
        },
        "admin.appRes": {
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
        "admin.dashboardRes": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                }
            }
        },
        "admin.groupRes": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "boolean"
                },
                "appId": {
                    "type": "integer"
                },
                "appName": {
                    "type": "string"
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
                "type": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "admin.staffRes": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "boolean"
                },
                "createdAt": {
                    "type": "string"
                },
                "departure": {
                    "type": "boolean"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "super": {
                    "type": "boolean"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "admin.tokenRes": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "accountId": {
                    "type": "integer"
                },
                "appId": {
                    "type": "integer"
                },
                "appName": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "refresh_token": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "admin.updateAppReq": {
            "type": "object",
            "properties": {
                "activeToggle": {
                    "type": "boolean",
                    "example": true
                },
                "resetSecret": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "admin.updateStaffReq": {
            "type": "object",
            "properties": {
                "activeToggle": {
                    "type": "boolean",
                    "example": true
                },
                "departureToggle": {
                    "type": "boolean",
                    "example": false
                },
                "phone": {
                    "type": "string",
                    "maxLength": 11,
                    "minLength": 11,
                    "example": "18555201314"
                },
                "resetPassword": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "inkstone.ClientError": {
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
        "inkstone.PagingResponse-admin_groupRes": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/admin.groupRes"
                    }
                },
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "inkstone.PagingResponse-admin_staffRes": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/admin.staffRes"
                    }
                },
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "inkstone.PagingResponse-admin_tokenRes": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/admin.tokenRes"
                    }
                },
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
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
