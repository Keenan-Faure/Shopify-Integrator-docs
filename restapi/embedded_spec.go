// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Shopify Integrator Docs",
    "version": "0.1.0"
  },
  "paths": {
    "/products": {
      "get": {
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "type": "integer",
            "format": "int32",
            "default": 1,
            "name": "page",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "lists products",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/product"
              }
            }
          }
        }
      }
    }
  },
  "definitions": {
    "product": {
      "type": "object",
      "required": [
        "product_code",
        "title",
        "variants"
      ],
      "properties": {
        "body_html": {
          "type": "string"
        },
        "category": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "product_code": {
          "type": "string",
          "minLength": 1
        },
        "product_type": {
          "type": "string"
        },
        "title": {
          "type": "string",
          "minLength": 1
        },
        "variants": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/variants"
          }
        },
        "vendor": {
          "type": "string"
        }
      }
    },
    "variants": {
      "type": "object",
      "required": [
        "sku"
      ],
      "properties": {
        "barcode": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "option1": {
          "type": "string"
        },
        "option2": {
          "type": "string"
        },
        "option3": {
          "type": "string"
        },
        "product_id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "sku": {
          "type": "string",
          "minLength": 1
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Shopify Integrator Docs",
    "version": "0.1.0"
  },
  "paths": {
    "/products": {
      "get": {
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "type": "integer",
            "format": "int32",
            "default": 1,
            "name": "page",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "lists products",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/product"
              }
            }
          }
        }
      }
    }
  },
  "definitions": {
    "product": {
      "type": "object",
      "required": [
        "product_code",
        "title",
        "variants"
      ],
      "properties": {
        "body_html": {
          "type": "string"
        },
        "category": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "product_code": {
          "type": "string",
          "minLength": 1
        },
        "product_type": {
          "type": "string"
        },
        "title": {
          "type": "string",
          "minLength": 1
        },
        "variants": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/variants"
          }
        },
        "vendor": {
          "type": "string"
        }
      }
    },
    "variants": {
      "type": "object",
      "required": [
        "sku"
      ],
      "properties": {
        "barcode": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "option1": {
          "type": "string"
        },
        "option2": {
          "type": "string"
        },
        "option3": {
          "type": "string"
        },
        "product_id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "sku": {
          "type": "string",
          "minLength": 1
        }
      }
    }
  }
}`))
}
