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
        "/admin": {
            "get": {
                "description": "Get a list of all Admins",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "List all Admins",
                "operationId": "list-all-Admins",
                "responses": {
                    "200": {
                        "description": "Successfully retrieved Admins",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/request.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/dto.AdminResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new Admin with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Create a Admin",
                "operationId": "create-Admin",
                "parameters": [
                    {
                        "description": "Admin data",
                        "name": "Admin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AdminRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully created Admin",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/request.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.AdminResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/admin/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get a Admin by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Get a Admin",
                "operationId": "get-Admin-by-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Admin ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/request.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.AdminResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a Admin by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Delete a Admin",
                "operationId": "delete-Admin-by-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Admin ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully deleted Admin",
                        "schema": {
                            "$ref": "#/definitions/request.Response"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update a Admin by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Update a Admin",
                "operationId": "update-Admin-by-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Admin ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Admin data",
                        "name": "Admin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AdminRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully updated Admin",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/request.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.AdminResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/permission": {
            "get": {
                "description": "Get a list of all Permissions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Permission"
                ],
                "summary": "List all Permissions",
                "operationId": "list-all-Permissions",
                "responses": {
                    "200": {
                        "description": "Successfully retrieved Permissions",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/request.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/dto.PermissionResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new Permission with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Permission"
                ],
                "summary": "Create a Permission",
                "operationId": "create-Permission",
                "parameters": [
                    {
                        "description": "Permission data",
                        "name": "Permission",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.PermissionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully created Permission",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/request.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.PermissionResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/permission/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get a Permission by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Permission"
                ],
                "summary": "Get a Permission",
                "operationId": "get-Permission-by-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Permission ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/request.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.PermissionResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a Permission by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Permission"
                ],
                "summary": "Delete a Permission",
                "operationId": "delete-Permission-by-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Permission ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully deleted Permission",
                        "schema": {
                            "$ref": "#/definitions/request.Response"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update a Permission by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Permission"
                ],
                "summary": "Update a Permission",
                "operationId": "update-Permission-by-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Permission ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Permission data",
                        "name": "Permission",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.PermissionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully updated Permission",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/request.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.PermissionResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/role": {
            "get": {
                "description": "Get a list of all Roles",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role"
                ],
                "summary": "List all Roles",
                "operationId": "list-all-Roles",
                "responses": {
                    "200": {
                        "description": "Successfully retrieved Roles",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/request.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/dto.RoleResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new Role with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role"
                ],
                "summary": "Create a Role",
                "operationId": "create-Role",
                "parameters": [
                    {
                        "description": "Role data",
                        "name": "Role",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RoleRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully created Role",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/request.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.RoleResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/role/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get a Role by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role"
                ],
                "summary": "Get a Role",
                "operationId": "get-Role-by-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Role ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/request.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.RoleResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a Role by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role"
                ],
                "summary": "Delete a Role",
                "operationId": "delete-Role-by-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Role ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully deleted Role",
                        "schema": {
                            "$ref": "#/definitions/request.Response"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update a Role by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role"
                ],
                "summary": "Update a Role",
                "operationId": "update-Role-by-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Role ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Role data",
                        "name": "Role",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RoleRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully updated Role",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/request.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.RoleResponse"
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
        "dto.AdminRequest": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "CreatedAt holds the value of the \"created_at\" field.",
                    "type": "string"
                },
                "displayName": {
                    "description": "DisplayName holds the value of the \"display_name\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the AdminQuery when eager-loading is set.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.AdminEdges"
                        }
                    ]
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "updatedAt": {
                    "description": "UpdatedAt holds the value of the \"updated_at\" field.",
                    "type": "string"
                },
                "username": {
                    "description": "Username holds the value of the \"username\" field.",
                    "type": "string"
                }
            }
        },
        "dto.AdminResponse": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "CreatedAt holds the value of the \"created_at\" field.",
                    "type": "string"
                },
                "displayName": {
                    "description": "DisplayName holds the value of the \"display_name\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the AdminQuery when eager-loading is set.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.AdminEdges"
                        }
                    ]
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "updatedAt": {
                    "description": "UpdatedAt holds the value of the \"updated_at\" field.",
                    "type": "string"
                },
                "username": {
                    "description": "Username holds the value of the \"username\" field.",
                    "type": "string"
                }
            }
        },
        "dto.PermissionRequest": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the PermissionQuery when eager-loading is set.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.PermissionEdges"
                        }
                    ]
                },
                "group": {
                    "description": "Group holds the value of the \"group\" field.",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "key": {
                    "description": "Key holds the value of the \"key\" field.",
                    "type": "string"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                }
            }
        },
        "dto.PermissionResponse": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the PermissionQuery when eager-loading is set.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.PermissionEdges"
                        }
                    ]
                },
                "group": {
                    "description": "Group holds the value of the \"group\" field.",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "key": {
                    "description": "Key holds the value of the \"key\" field.",
                    "type": "string"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                }
            }
        },
        "dto.RoleRequest": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "CreatedAt holds the value of the \"created_at\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the RoleQuery when eager-loading is set.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.RoleEdges"
                        }
                    ]
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                },
                "updatedAt": {
                    "description": "UpdatedAt holds the value of the \"updated_at\" field.",
                    "type": "string"
                }
            }
        },
        "dto.RoleResponse": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "CreatedAt holds the value of the \"created_at\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the RoleQuery when eager-loading is set.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.RoleEdges"
                        }
                    ]
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                },
                "updatedAt": {
                    "description": "UpdatedAt holds the value of the \"updated_at\" field.",
                    "type": "string"
                }
            }
        },
        "ent.Admin": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "CreatedAt holds the value of the \"created_at\" field.",
                    "type": "string"
                },
                "displayName": {
                    "description": "DisplayName holds the value of the \"display_name\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the AdminQuery when eager-loading is set.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.AdminEdges"
                        }
                    ]
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "updatedAt": {
                    "description": "UpdatedAt holds the value of the \"updated_at\" field.",
                    "type": "string"
                },
                "username": {
                    "description": "Username holds the value of the \"username\" field.",
                    "type": "string"
                }
            }
        },
        "ent.AdminEdges": {
            "type": "object",
            "properties": {
                "roles": {
                    "description": "Roles holds the value of the roles edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Role"
                    }
                }
            }
        },
        "ent.Permission": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the PermissionQuery when eager-loading is set.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.PermissionEdges"
                        }
                    ]
                },
                "group": {
                    "description": "Group holds the value of the \"group\" field.",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "key": {
                    "description": "Key holds the value of the \"key\" field.",
                    "type": "string"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                }
            }
        },
        "ent.PermissionEdges": {
            "type": "object",
            "properties": {
                "roles": {
                    "description": "Roles holds the value of the roles edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Role"
                    }
                }
            }
        },
        "ent.Role": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "CreatedAt holds the value of the \"created_at\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the RoleQuery when eager-loading is set.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.RoleEdges"
                        }
                    ]
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                },
                "updatedAt": {
                    "description": "UpdatedAt holds the value of the \"updated_at\" field.",
                    "type": "string"
                }
            }
        },
        "ent.RoleEdges": {
            "type": "object",
            "properties": {
                "admins": {
                    "description": "Admins holds the value of the admins edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Admin"
                    }
                },
                "permissions": {
                    "description": "Permissions holds the value of the permissions edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Permission"
                    }
                }
            }
        },
        "request.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "errors": {
                    "type": "array",
                    "items": {}
                },
                "message": {
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
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "GVA API",
	Description:      "GO VUE ADMIN Boilerplate",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
