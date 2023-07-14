// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/api/auth/google-auth": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for user to login with google",
                "tags": [
                    "User Authentication"
                ],
                "summary": "Login with google (User)",
                "operationId": "UserGoogleLogin",
                "parameters": [
                    {
                        "description": "Google Token Input",
                        "name": "inputs",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.GoogleLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully logged in with google",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.TokenResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid inputs",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Failed to login",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/auth/login": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for user to login with email | phone | user_name with password",
                "tags": [
                    "User Authentication"
                ],
                "summary": "Login with password (User)",
                "operationId": "UserLogin",
                "parameters": [
                    {
                        "description": "Login Details",
                        "name": "inputs",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully logged in",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.TokenResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid inputs",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "401": {
                        "description": "User not exist with given login credentials",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "403": {
                        "description": "User blocked by admin",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Failed to login",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/auth/renew-access-token": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for user to renew access token using refresh token",
                "tags": [
                    "User Authentication"
                ],
                "summary": "Renew Access Token (User)",
                "operationId": "UserRenewAccessToken",
                "parameters": [
                    {
                        "description": "Refresh token",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RefreshToken"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully generated access token using refresh token",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.TokenResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "401": {
                        "description": "Invalid refresh token",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "403": {
                        "description": "Refresh token blocked",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "404": {
                        "description": "No session found for the given refresh token",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "410": {
                        "description": "Refresh token expired",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Failed generate access token",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/auth/signup": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for user to register a new account",
                "tags": [
                    "User Authentication"
                ],
                "summary": "Signup (User)",
                "operationId": "UserSignUp",
                "parameters": [
                    {
                        "description": "Input Fields",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserSignUp"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Successfully Account Created",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "409": {
                        "description": "A verified user already exist with given user credentials",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Failed to signup",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/users/all": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for user to list all users in the application",
                "tags": [
                    "Users"
                ],
                "summary": "List all users (User)",
                "operationId": "ListUsers",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page Number",
                        "name": "page_number",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Count",
                        "name": "count",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully retrieved all users",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/response.User"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "204": {
                        "description": "There is no users",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Failed retrieved all users",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.GoogleLogin": {
            "type": "object",
            "required": [
                "token"
            ],
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "request.Login": {
            "type": "object",
            "required": [
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 5
                },
                "user_name": {
                    "type": "string",
                    "maxLength": 15,
                    "minLength": 3
                }
            }
        },
        "request.RefreshToken": {
            "type": "object",
            "properties": {
                "refresh_token": {
                    "type": "string",
                    "minLength": 10
                }
            }
        },
        "request.UserSignUp": {
            "type": "object",
            "required": [
                "age",
                "email",
                "first_name",
                "last_name",
                "password",
                "user_name"
            ],
            "properties": {
                "age": {
                    "type": "integer",
                    "minimum": 13
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                },
                "last_name": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 1
                },
                "password": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 5
                },
                "user_name": {
                    "type": "string",
                    "maxLength": 15,
                    "minLength": 3
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {},
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "response.TokenResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "response.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "profile_image": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                },
                "user_name": {
                    "type": "string"
                }
            }
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
