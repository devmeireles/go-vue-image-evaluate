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
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/report": {
            "get": {
                "description": "List all reports",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "report"
                ],
                "summary": "Get a list of reports id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Report"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Saves a request for avatar evaluate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "report"
                ],
                "summary": "Endpoint to save avatar reports",
                "parameters": [
                    {
                        "description": "Create a report",
                        "name": "evaluate",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateReportDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Report"
                            }
                        }
                    }
                }
            }
        },
        "/report/{id}": {
            "get": {
                "description": "Get details of report corresponding to the input id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "report"
                ],
                "summary": "Get details for a given report id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Report ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Report"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreateReportDTO": {
            "type": "object",
            "required": [
                "report_id",
                "url"
            ],
            "properties": {
                "report_id": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "models.Report": {
            "type": "object",
            "required": [
                "report_id",
                "url"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "report_id": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "url": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Image Evaluate API",
	Description:      "A quickly API setup to evaulate avatar.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}