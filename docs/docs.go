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
        "/api/token/refresh/": {
            "post": {
                "description": "generate and get a new access token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "refresh access token",
                "parameters": [
                    {
                        "description": "user profile update",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RefreshToken"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "user profile updated",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.BaseHttpResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/dto.TokenDetail"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/helper.BaseHttpResponseWithValidationError"
                        }
                    },
                    "406": {
                        "description": "not acceptable",
                        "schema": {
                            "$ref": "#/definitions/helper.BaseHttpResponseWithError"
                        }
                    }
                }
            }
        },
        "/api/users/login/": {
            "post": {
                "description": "user login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "user login",
                "parameters": [
                    {
                        "description": "user login",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "user login",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.BaseHttpResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/dto.TokenDetail"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/helper.BaseHttpResponseWithValidationError"
                        }
                    },
                    "406": {
                        "description": "error while login user",
                        "schema": {
                            "$ref": "#/definitions/helper.BaseHttpResponseWithError"
                        }
                    }
                }
            }
        },
        "/api/users/profile/": {
            "get": {
                "security": [
                    {
                        "AuthBearer": []
                    }
                ],
                "description": "user profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "user profile",
                "responses": {
                    "200": {
                        "description": "user profile info",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.BaseHttpResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/dto.Profile"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "UnAuthorization",
                        "schema": {
                            "$ref": "#/definitions/helper.BaseHttpResponseWithError"
                        }
                    }
                }
            }
        },
        "/api/users/profile/update/": {
            "patch": {
                "security": [
                    {
                        "AuthBearer": []
                    }
                ],
                "description": "user profile update",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "user profile update",
                "parameters": [
                    {
                        "description": "user profile update",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ProfileUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "user profile updated",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.BaseHttpResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/dto.Profile"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/helper.BaseHttpResponseWithValidationError"
                        }
                    },
                    "406": {
                        "description": "not acceptable",
                        "schema": {
                            "$ref": "#/definitions/helper.BaseHttpResponseWithError"
                        }
                    }
                }
            }
        },
        "/api/users/register/": {
            "post": {
                "description": "register user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "register user",
                "parameters": [
                    {
                        "description": "ًRegister User",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateUser"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "register user",
                        "schema": {
                            "$ref": "#/definitions/helper.BaseHttpResponse"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/helper.BaseHttpResponseWithValidationError"
                        }
                    },
                    "406": {
                        "description": "error while register user",
                        "schema": {
                            "$ref": "#/definitions/helper.BaseHttpResponseWithError"
                        }
                    }
                }
            }
        },
        "/api/users/verify/": {
            "post": {
                "description": "user verify account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "user verify account",
                "parameters": [
                    {
                        "description": "User verify",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserVerify"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "user verified",
                        "schema": {
                            "$ref": "#/definitions/helper.BaseHttpResponse"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/helper.BaseHttpResponseWithValidationError"
                        }
                    },
                    "406": {
                        "description": "error while verifying user",
                        "schema": {
                            "$ref": "#/definitions/helper.BaseHttpResponseWithError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreateUser": {
            "type": "object",
            "required": [
                "email",
                "password",
                "userName"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 10
                },
                "fullName": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 5
                },
                "password": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 8
                },
                "userName": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 5
                }
            }
        },
        "dto.Profile": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "fullName": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "dto.ProfileUpdate": {
            "type": "object",
            "properties": {
                "fullName": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 5
                }
            }
        },
        "dto.RefreshToken": {
            "type": "object",
            "required": [
                "refreshToken"
            ],
            "properties": {
                "refreshToken": {
                    "type": "string"
                }
            }
        },
        "dto.TokenDetail": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "accessTokenExpireTime": {
                    "type": "integer"
                },
                "refreshToken": {
                    "type": "string"
                },
                "refreshTokenExpireTime": {
                    "type": "integer"
                }
            }
        },
        "dto.UserLogin": {
            "type": "object",
            "required": [
                "password",
                "userName"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 8
                },
                "userName": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 5
                }
            }
        },
        "dto.UserVerify": {
            "type": "object",
            "required": [
                "code",
                "userName"
            ],
            "properties": {
                "code": {
                    "type": "integer",
                    "maximum": 9999,
                    "minimum": 1000
                },
                "userName": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 5
                }
            }
        },
        "helper.BaseHttpResponse": {
            "type": "object",
            "properties": {
                "result": {},
                "statusCode": {
                    "type": "integer"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "helper.BaseHttpResponseWithError": {
            "type": "object",
            "properties": {
                "error": {},
                "result": {},
                "statusCode": {
                    "type": "integer"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "helper.BaseHttpResponseWithValidationError": {
            "type": "object",
            "properties": {
                "result": {},
                "statusCode": {
                    "type": "integer"
                },
                "success": {
                    "type": "boolean"
                },
                "validationErrors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/helper.ValidationError"
                    }
                }
            }
        },
        "helper.ValidationError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "key": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "AuthBearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
