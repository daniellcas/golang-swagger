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
            "name": "Daniel lucas",
            "email": "danielsantos120615@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/finances": {
            "get": {
                "description": "Get finances",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "finances"
                ],
                "summary": "Get finances",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handlers.OutputGetFinances"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.OutputGetFinancesFailure"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.OutputGetFinances": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "date": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "handlers.OutputGetFinancesFailure": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v0.0.1",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Finances example api",
	Description:      "This is an example of an api using swagger",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
