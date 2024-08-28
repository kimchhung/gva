// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplatebot = `{
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
        "/file/static/img/{name}": {
            "get": {
                "description": "Serves files from ` + "`" + `storage/static` + "`" + ` directory",
                "tags": [
                    "File"
                ],
                "summary": "Serve static files",
                "operationId": "serve-static-files",
                "parameters": [
                    {
                        "type": "string",
                        "description": "filename",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/file/upload-img": {
            "post": {
                "description": "Upload a file",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "File"
                ],
                "summary": "Upload a file",
                "operationId": "upload-file",
                "parameters": [
                    {
                        "type": "file",
                        "description": "File",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/health-check": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Index"
                ],
                "summary": "Health Check",
                "operationId": "health-check",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/now": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Index"
                ],
                "summary": "Current Server Time",
                "operationId": "now",
                "responses": {
                    "200": {
                        "description": "format time.RFC3339",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
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
        }
    },
    "definitions": {
        "response.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "description": "The Data field contains the actual response data"
                },
                "message": {
                    "type": "string"
                },
                "meta": {
                    "description": "Meta provides additional information about the data, such as its type or kind.y."
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfobot holds exported Swagger Info so clients can modify it
var SwaggerInfobot = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:5500",
	BasePath:         "/bot/v1",
	Schemes:          []string{},
	Title:            "GVA bot API",
	Description:      "GO VUE ADMIN Boilerplate",
	InfoInstanceName: "bot",
	SwaggerTemplate:  docTemplatebot,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfobot.InstanceName(), SwaggerInfobot)
}
