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
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/helper.BaseHttpResponse"
                        }
                    },
                    "400": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/helper.BaseHttpResponseWithValidationError"
                        }
                    },
                    "406": {
                        "description": "Failed",
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
