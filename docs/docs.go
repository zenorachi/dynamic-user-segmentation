// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Maksim Sonkin",
            "email": "msonkin33@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/operations/add_segments/": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "addition a user to segments by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "operations"
                ],
                "summary": "Add a user to segments by id",
                "parameters": [
                    {
                        "description": "input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.addSegmentsByIdInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/v1.operationsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/operations/add_segments_by_names/": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "addition a user to segments by name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "operations"
                ],
                "summary": "Add a user to segments by name",
                "parameters": [
                    {
                        "description": "input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.addSegmentsByNameInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/v1.operationsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/operations/delete_segments/": {
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "delete user-segments relation by ids",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "operations"
                ],
                "summary": "Delete User From Segments by ids",
                "parameters": [
                    {
                        "description": "input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.deleteSegmentsByIdInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.operationsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/operations/delete_segments_by_names/": {
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "delete user-segments relation by names",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "operations"
                ],
                "summary": "Delete User From Segments By Names",
                "parameters": [
                    {
                        "description": "input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.deleteSegmentsByNameInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.operationsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/operations/history/": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "getting operations history",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "operations"
                ],
                "summary": "Get Operations History",
                "parameters": [
                    {
                        "description": "input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.getOperationsHistoryInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.getOperationsHistoryResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/reports/file/": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "getting history in csv format",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reports"
                ],
                "summary": "Get CSV-File History",
                "parameters": [
                    {
                        "description": "input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.getReportInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "CSV-File",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/reports/link/": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "getting history by link to csv-file",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reports"
                ],
                "summary": "Get CSV-File Link",
                "parameters": [
                    {
                        "description": "input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.getReportInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.getReportLinkResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/segments/": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "getting all segments",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "segments"
                ],
                "summary": "Get all segments",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.getAllSegmentsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/segments/:segment_id": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "getting segment by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "segments"
                ],
                "summary": "Get Segment By ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Segment ID",
                        "name": "segment_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.getSegmentByIdResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/segments/active_users/:segment_id": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "getting active users by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "segment-users"
                ],
                "summary": "Get Active Users By ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Segment ID",
                        "name": "segment_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.getActiveUsersResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/segments/create": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "create new segment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "segments"
                ],
                "summary": "Create segment",
                "parameters": [
                    {
                        "description": "input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.createSegmentInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/v1.createSegmentResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/segments/delete/": {
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "deletion segment by name",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "segments"
                ],
                "summary": "Delete Segment By Name",
                "parameters": [
                    {
                        "description": "input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.deleteByNameInput"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/segments/delete_by_id/": {
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "deletion segment by id",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "segments"
                ],
                "summary": "Delete Segment By ID",
                "parameters": [
                    {
                        "description": "input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.deleteByIdInput"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/users/active_segments/:user_id": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "get active segments for a specific user by user_id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user-segments"
                ],
                "summary": "Get active segments for a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.getActiveSegmentsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/users/refresh": {
            "get": {
                "description": "refresh user's access token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "User Refresh Token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.tokenResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/users/sign-in": {
            "post": {
                "description": "user sign in",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "User SignIn",
                "parameters": [
                    {
                        "description": "input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.signInInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.tokenResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/users/sign-up": {
            "post": {
                "description": "create user account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "User SignUp",
                "parameters": [
                    {
                        "description": "input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.signUpInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/v1.signUpResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Operation": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "segment_name": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "entity.Segment": {
            "type": "object",
            "properties": {
                "assign_percent": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "entity.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "registered_at": {
                    "type": "string"
                }
            }
        },
        "v1.addSegmentsByIdInput": {
            "type": "object",
            "required": [
                "segment_ids",
                "user_id"
            ],
            "properties": {
                "segment_ids": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "ttl": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "v1.addSegmentsByNameInput": {
            "type": "object",
            "required": [
                "segment_names",
                "user_id"
            ],
            "properties": {
                "segment_names": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "ttl": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "v1.createSegmentInput": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "assign_percent": {
                    "type": "number"
                },
                "name": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 2
                }
            }
        },
        "v1.createSegmentResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "v1.deleteByIdInput": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "v1.deleteByNameInput": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "v1.deleteSegmentsByIdInput": {
            "type": "object",
            "required": [
                "segment_ids",
                "user_id"
            ],
            "properties": {
                "segment_ids": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "v1.deleteSegmentsByNameInput": {
            "type": "object",
            "required": [
                "segment_names",
                "user_id"
            ],
            "properties": {
                "segment_names": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "v1.errorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "v1.getActiveSegmentsResponse": {
            "type": "object",
            "properties": {
                "segments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Segment"
                    }
                }
            }
        },
        "v1.getActiveUsersResponse": {
            "type": "object",
            "properties": {
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.User"
                    }
                }
            }
        },
        "v1.getAllSegmentsResponse": {
            "type": "object",
            "properties": {
                "segments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Segment"
                    }
                }
            }
        },
        "v1.getOperationsHistoryInput": {
            "type": "object",
            "required": [
                "month",
                "year"
            ],
            "properties": {
                "month": {
                    "type": "integer"
                },
                "page_size": {
                    "type": "integer"
                },
                "user_ids": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "v1.getOperationsHistoryResponse": {
            "type": "object",
            "properties": {
                "operations": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Operation"
                    }
                }
            }
        },
        "v1.getReportInput": {
            "type": "object",
            "required": [
                "month",
                "year"
            ],
            "properties": {
                "month": {
                    "type": "integer"
                },
                "user_ids": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "v1.getReportLinkResponse": {
            "type": "object",
            "properties": {
                "link": {
                    "type": "string"
                }
            }
        },
        "v1.getSegmentByIdResponse": {
            "type": "object",
            "properties": {
                "segment": {
                    "$ref": "#/definitions/entity.Segment"
                }
            }
        },
        "v1.operationsResponse": {
            "type": "object",
            "properties": {
                "operation_ids": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "v1.signInInput": {
            "type": "object",
            "required": [
                "login",
                "password"
            ],
            "properties": {
                "login": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 2
                },
                "password": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 8
                }
            }
        },
        "v1.signUpInput": {
            "type": "object",
            "required": [
                "email",
                "login",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 64
                },
                "login": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 2
                },
                "password": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 8
                }
            }
        },
        "v1.signUpResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "v1.tokenResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
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

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Dynamic User Segmentation Service",
	Description:      "This is a service for segmenting users with the ability to automatically add and remove users from segments.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
