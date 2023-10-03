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
        "license": {
            "name": "Apache 2.0"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "description": "checking auth",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authorization"
                ],
                "summary": "checking auth",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Checking user authentication",
                        "name": "cookie",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success authenticate return id",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "401": {
                        "description": "unauthorized",
                        "schema": {}
                    }
                }
            }
        },
        "/logout": {
            "post": {
                "description": "Log out user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authorization"
                ],
                "summary": "Log out user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Log out user",
                        "name": "cookie",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "success log out"
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {}
                    },
                    "401": {
                        "description": "unauthorized",
                        "schema": {}
                    }
                }
            }
        },
        "/users": {
            "post": {
                "description": "Signing up a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Signing up a user",
                "parameters": [
                    {
                        "description": "User object for signing up",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/store.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success create User return id",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {}
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "store.User": {
            "type": "object",
            "properties": {
                "birthday": {
                    "type": "string"
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
                "phoneNumber": {
                    "type": "string"
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
	Host:             "http://84.23.53.216:8001/",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Prinesi-Poday API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
