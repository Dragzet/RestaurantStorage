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
        "/product": {
            "get": {
                "description": "Retrieves the list of all products in the storage.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Get list of products",
                "responses": {
                    "200": {
                        "description": "List of products",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.ProductModel"
                            }
                        }
                    },
                    "500": {
                        "description": "error: Failed to get products list",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Updates the details of an existing product.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Update a product",
                "parameters": [
                    {
                        "description": "Product to update",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ProductModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "status: ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error: Failed to bind product",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error: Failed to change product",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Adds a new product to the storage.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Add a new product",
                "parameters": [
                    {
                        "description": "Product to add",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ProductModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "status: ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error: Failed to bind product",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error: Failed to add product",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a product from the storage by name.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Delete a product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "status: ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error: Failed to get name",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error: Failed to delete product",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ProductModel": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "name": {
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
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Restaurant Storage API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}