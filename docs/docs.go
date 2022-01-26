// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/books": {
            "get": {
                "description": "Get all books",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "Get all books",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/schemas.APIResponseBooks"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schemas.APIResponse400"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/schemas.APIResponse404"
                        }
                    }
                }
            },
            "post": {
                "description": "Post a book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "Post a book",
                "parameters": [
                    {
                        "description": "Body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.BookRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/schemas.APIResponseBook"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schemas.APIResponse400"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/schemas.APIResponse404"
                        }
                    }
                }
            }
        },
        "/books/{id}": {
            "get": {
                "description": "Find book by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "Find book by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/schemas.APIResponseBook"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schemas.APIResponse400"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/schemas.APIResponse404"
                        }
                    }
                }
            },
            "put": {
                "description": "Update the book by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "Update the book by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.BookRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/schemas.APIResponseBook"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schemas.APIResponse400"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/schemas.APIResponse404"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete book by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "Delete book by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/schemas.APIResponse200"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schemas.APIResponse400"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/schemas.APIResponse404"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "schemas.APIResponse200": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "message": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "schemas.APIResponse400": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string",
                    "example": "not found"
                }
            }
        },
        "schemas.APIResponse404": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 404
                },
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string",
                    "example": "bad request"
                }
            }
        },
        "schemas.APIResponseBook": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "data": {
                    "$ref": "#/definitions/schemas.BookResponse"
                },
                "message": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "schemas.APIResponseBooks": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schemas.BookResponse"
                    }
                },
                "message": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "schemas.BookRequest": {
            "type": "object",
            "required": [
                "price",
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "rating": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "schemas.BookResponse": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "price": {
                    "type": "integer"
                },
                "rating": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
