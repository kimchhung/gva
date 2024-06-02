// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplateweb = `{
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
        "/now": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Time"
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
        },
        "/sse/now": {
            "get": {
                "consumes": [
                    "text/event-stream"
                ],
                "tags": [
                    "Time"
                ],
                "summary": "SSE Current Server Time",
                "operationId": "sse-now",
                "responses": {
                    "200": {
                        "description": "format time.RFC3339",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/ws/now": {
            "get": {
                "consumes": [
                    "text/event-stream"
                ],
                "tags": [
                    "Time"
                ],
                "summary": "WS Current Server Time",
                "operationId": "ws-now",
                "responses": {
                    "200": {
                        "description": "format time.RFC3339",
                        "schema": {
                            "type": "string"
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

// SwaggerInfoweb holds exported Swagger Info so clients can modify it
var SwaggerInfoweb = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/web/v1",
	Schemes:          []string{},
	Title:            "GVA Web API",
	Description:      "GO VUE ADMIN Boilerplate",
	InfoInstanceName: "web",
	SwaggerTemplate:  docTemplateweb,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfoweb.InstanceName(), SwaggerInfoweb)
}
