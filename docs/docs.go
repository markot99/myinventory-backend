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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/items": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Get a list of items belongig to the user with reduced information",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "items"
                ],
                "summary": "Get a list of items",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/itemcontroller.GetItemsResponseBody"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Add an item to the database",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "items"
                ],
                "summary": "Add an item",
                "parameters": [
                    {
                        "description": "Item Body",
                        "name": "item",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/itemcontroller.AddItemRequestBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/itemcontroller.AddItemResponseBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/items/": {
            "put": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Edit an item in the database",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "items"
                ],
                "summary": "Edit an item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "item id",
                        "name": "itemID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Item Body",
                        "name": "item",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/itemcontroller.AddItemRequestBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/itemcontroller.AddItemResponseBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/items/{id}": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Get detailed information about an item from the database",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "items"
                ],
                "summary": "Get an item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Object id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Item"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/items/{itemID}": {
            "delete": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Delete an item from the database including invoice and all images",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "items"
                ],
                "summary": "Delete an item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "item id",
                        "name": "itemID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/v1/items/{itemID}/images": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Upload one ore more images to the server and get their internal ids",
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "images"
                ],
                "summary": "Upload images",
                "parameters": [
                    {
                        "type": "string",
                        "description": "item id",
                        "name": "itemID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "file"
                        },
                        "collectionFormat": "csv",
                        "description": "images of the item",
                        "name": "images",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/items/{itemID}/images/preview": {
            "put": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Set preview image for an item",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "images"
                ],
                "summary": "Set preview image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "item id",
                        "name": "itemID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Set Preview Image Name",
                        "name": "item",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/imagecontroller.SetPreviewImageRequestBody"
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
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/items/{itemID}/images/{imageID}": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Get an image from the item",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "images"
                ],
                "summary": "Get an image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "item id",
                        "name": "itemID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "image id",
                        "name": "imageID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Image",
                        "schema": {
                            "type": "file"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Delete an image from the item",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "images"
                ],
                "summary": "Delete an image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "item id",
                        "name": "itemID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "image id",
                        "name": "imageID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Image",
                        "schema": {
                            "type": "file"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/items/{itemID}/invoice": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Get the invoice from the item",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "invoice"
                ],
                "summary": "Get invoice",
                "parameters": [
                    {
                        "type": "string",
                        "description": "item id",
                        "name": "itemID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Invoice",
                        "schema": {
                            "type": "file"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Upload invoice to the item",
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "invoice"
                ],
                "summary": "Upload invoice",
                "parameters": [
                    {
                        "type": "string",
                        "description": "item id",
                        "name": "itemID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "invoice of the item",
                        "name": "invoice",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Delete the invoice from the item and from the storage",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "invoice"
                ],
                "summary": "Delete invoice",
                "parameters": [
                    {
                        "type": "string",
                        "description": "item id",
                        "name": "itemID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/users/login": {
            "post": {
                "description": "Login user and retrieve an authorization token",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "Login Body",
                        "name": "item",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usercontroller.LoginUserRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usercontroller.LoginUserResponseBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/users/me": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Get user info about the logged in user",
                "tags": [
                    "users"
                ],
                "summary": "Get user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usercontroller.MeResponseBody"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/users/register": {
            "post": {
                "description": "Register a new user",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Register user",
                "parameters": [
                    {
                        "description": "Registration Body",
                        "name": "item",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usercontroller.RegisterUserRequestBody"
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
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.APIErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.APIErrorResponse": {
            "type": "object",
            "properties": {
                "diagnosisCode": {
                    "description": "diagnosis code of the error",
                    "type": "integer"
                },
                "diagnosisMessage": {
                    "description": "diagnosis message of the error",
                    "type": "string"
                }
            }
        },
        "imagecontroller.SetPreviewImageRequestBody": {
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
        "itemcontroller.AddItemPurchaseInfo": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "place": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                },
                "unitPrice": {
                    "type": "number"
                }
            }
        },
        "itemcontroller.AddItemRequestBody": {
            "type": "object",
            "required": [
                "name",
                "purchaseInfo"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "purchaseInfo": {
                    "$ref": "#/definitions/itemcontroller.AddItemPurchaseInfo"
                }
            }
        },
        "itemcontroller.AddItemResponseBody": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "itemcontroller.GetItemsResponseBody": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/itemcontroller.ResponseItem"
                    }
                }
            }
        },
        "itemcontroller.ResponseItem": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "previewImage": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "models.File": {
            "type": "object",
            "properties": {
                "fileName": {
                    "description": "name of the file",
                    "type": "string"
                },
                "id": {
                    "description": "identification number of the file",
                    "type": "string"
                }
            }
        },
        "models.Images": {
            "type": "object",
            "properties": {
                "images": {
                    "description": "all images",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.File"
                    }
                },
                "previewImage": {
                    "description": "preview image name to be displayed first",
                    "type": "string"
                }
            }
        },
        "models.Item": {
            "type": "object",
            "properties": {
                "description": {
                    "description": "description of the item",
                    "type": "string"
                },
                "id": {
                    "description": "identification number of the item",
                    "type": "string"
                },
                "images": {
                    "description": "images of the item",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Images"
                        }
                    ]
                },
                "name": {
                    "description": "name of the item",
                    "type": "string"
                },
                "ownerID": {
                    "description": "identification number of the owner of the item",
                    "type": "string"
                },
                "purchaseInfo": {
                    "description": "purchase information of the item",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.PurchaseInfo"
                        }
                    ]
                }
            }
        },
        "models.PurchaseInfo": {
            "type": "object",
            "properties": {
                "date": {
                    "description": "date of purchase",
                    "type": "string"
                },
                "invoice": {
                    "description": "invoice of the purchase",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.File"
                        }
                    ]
                },
                "place": {
                    "description": "place of purchase",
                    "type": "string"
                },
                "quantity": {
                    "description": "number of items purchased",
                    "type": "integer"
                },
                "unitPrice": {
                    "description": "price of a single item",
                    "type": "number"
                }
            }
        },
        "usercontroller.LoginUserRequestBody": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "usercontroller.LoginUserResponseBody": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "usercontroller.MeResponseBody": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                }
            }
        },
        "usercontroller.RegisterUserRequestBody": {
            "type": "object",
            "required": [
                "email",
                "firstName",
                "lastName",
                "password"
            ],
            "properties": {
                "email": {
                    "description": "email of the user",
                    "type": "string"
                },
                "firstName": {
                    "description": "first name of the user",
                    "type": "string"
                },
                "lastName": {
                    "description": "last name of the user",
                    "type": "string"
                },
                "password": {
                    "description": "unhashed password of the user",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "JWT": {
            "description": "Description for what is this security definition being used",
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
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "myInventory API",
	Description:      "Swagger documentation to test the myInventory API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
