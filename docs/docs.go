// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "FIO API Support",
            "email": "danilakalash60@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/people": {
            "get": {
                "description": "Get people",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "People"
                ],
                "summary": "Get people",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "surname",
                        "name": "surname",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "patronymic",
                        "name": "patronymic",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "start_age",
                        "name": "start_age",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "end_age",
                        "name": "end_age",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "gender",
                        "name": "gender",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "nationality",
                        "name": "nationality",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.People"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "put": {
                "description": "Update people",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "People"
                ],
                "summary": "Update people",
                "parameters": [
                    {
                        "type": "string",
                        "description": "peopleID",
                        "name": "peopleID",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "New people info",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AbbreviatedPeopleRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.People"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "post": {
                "description": "Create people",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "People"
                ],
                "summary": "Create people",
                "parameters": [
                    {
                        "description": "Create people",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AbbreviatedPeopleRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.People"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "delete": {
                "description": "Delete people",
                "tags": [
                    "People"
                ],
                "summary": "Delete people",
                "parameters": [
                    {
                        "type": "string",
                        "description": "peopleID",
                        "name": "peopleID",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/models.NoContentResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "models.AbbreviatedPeopleRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "patronymic": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        },
        "models.NoContentResponse": {
            "type": "object"
        },
        "models.People": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "nationality": {
                    "type": "string"
                },
                "patronymic": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.1",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "FIO API",
	Description:      "Server API for FIO Service Application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
