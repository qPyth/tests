// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/ent-test": {
            "get": {
                "description": "get test by language and profiles subjects` + "`" + `s ids",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ENT test"
                ],
                "summary": "getting ENT test",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Language of test: ru or kz",
                        "name": "lang",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "id of first profile subject",
                        "name": "profile1",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "id of second profile subject",
                        "name": "profile2",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.EntTestOutput"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/transport.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Question": {
            "type": "object",
            "properties": {
                "countAnswers": {
                    "type": "integer"
                },
                "countVariants": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "level": {
                    "type": "integer"
                },
                "partID": {
                    "type": "integer"
                },
                "questionDetails": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.QuestionDetails"
                    }
                },
                "subjectID": {
                    "type": "integer"
                },
                "targetID": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                },
                "topicID": {
                    "type": "integer"
                },
                "type": {
                    "type": "integer"
                }
            }
        },
        "models.QuestionDetails": {
            "type": "object",
            "properties": {
                "answers": {
                    "type": "string"
                },
                "explanation": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "question": {
                    "type": "string"
                },
                "questionID": {
                    "type": "integer"
                },
                "source": {
                    "type": "string"
                },
                "var1": {
                    "type": "string"
                },
                "var10": {
                    "type": "string"
                },
                "var2": {
                    "type": "string"
                },
                "var3": {
                    "type": "string"
                },
                "var4": {
                    "type": "string"
                },
                "var5": {
                    "type": "string"
                },
                "var6": {
                    "type": "string"
                },
                "var7": {
                    "type": "string"
                },
                "var8": {
                    "type": "string"
                },
                "var9": {
                    "type": "string"
                }
            }
        },
        "services.EntTestOutput": {
            "type": "object",
            "properties": {
                "tests": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "array",
                        "items": {
                            "$ref": "#/definitions/models.Question"
                        }
                    }
                }
            }
        },
        "transport.ErrorResponse": {
            "type": "object",
            "properties": {
                "error_code": {
                    "type": "integer"
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
	Version:          "2.0",
	Host:             "localhost:8080",
	BasePath:         "/api/",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a server for getting tests",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
