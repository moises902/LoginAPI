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
    "description": "API used to get medicare data",
    "title": "Medicare API",
    "version": "0.0.1"
  },
  "basePath": "/go/api",
  "paths": {
    "/login": {
      "post": {
        "description": "Authenticates a user's email and password. Returns a cookie and token if successful",
        "tags": [
          "login"
        ],
        "summary": "Log ing with email and password",
        "parameters": [
          {
            "name": "login",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Login"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "User info",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "default": {
            "description": "Subscribe service unexpected error response",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          }
        }
      }
    },
    "/record": {
      "get": {
        "description": "Returns the details of a single record with a matching id",
        "tags": [
          "record"
        ],
        "summary": "Single record details",
        "parameters": [
          {
            "type": "integer",
            "description": "the service to be added",
            "name": "id",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Data based on parameters",
            "schema": {
              "$ref": "#/definitions/Record"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "default": {
            "description": "Subscribe service unexpected error response",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          }
        }
      }
    },
    "/records": {
      "get": {
        "description": "Returns a list of records based on some parameters",
        "tags": [
          "records"
        ],
        "summary": "List of records",
        "parameters": [
          {
            "type": "string",
            "name": "city",
            "in": "query"
          },
          {
            "type": "string",
            "name": "state",
            "in": "query"
          },
          {
            "type": "string",
            "name": "specialty",
            "in": "query"
          },
          {
            "type": "string",
            "name": "drugName",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Returned data based on parameters",
            "schema": {
              "$ref": "#/definitions/Records"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "default": {
            "description": "Subscribe service unexpected error response",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          }
        }
      }
    },
    "/register": {
      "post": {
        "description": "Adds a new user to the application DB",
        "tags": [
          "register"
        ],
        "summary": "Register a new user",
        "parameters": [
          {
            "name": "register",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Register"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Return success",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "default": {
            "description": "Subscribe service unexpected error response",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Login": {
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
    "MinRecord": {
      "type": "object",
      "properties": {
        "city": {
          "type": "string"
        },
        "drugName": {
          "type": "string"
        },
        "firstName": {
          "type": "string"
        },
        "genericName": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "lastOrgName": {
          "type": "string"
        },
        "specialty": {
          "type": "string"
        },
        "state": {
          "type": "string"
        }
      }
    },
    "Record": {
      "type": "object",
      "properties": {
        "beneCount": {
          "type": "string"
        },
        "beneCountGe65": {
          "type": "string"
        },
        "beneCountGe65Flag": {
          "type": "string"
        },
        "beneCountNum": {
          "type": "integer"
        },
        "city": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "drugName": {
          "type": "string"
        },
        "firstName": {
          "type": "string"
        },
        "ge65SuppressFlag": {
          "type": "string"
        },
        "genericName": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        },
        "lastOrgName": {
          "type": "string"
        },
        "npi": {
          "type": "string"
        },
        "specialty": {
          "type": "string"
        },
        "state": {
          "type": "string"
        },
        "total30DayFillCount": {
          "type": "string"
        },
        "total30DayFillCountGe65": {
          "type": "string"
        },
        "totalClaimCount": {
          "type": "string"
        },
        "totalClaimCountGe65": {
          "type": "string"
        },
        "totalClaimCountNum": {
          "type": "string"
        },
        "totalDaySupply": {
          "type": "string"
        },
        "totalDaySupplyGe65": {
          "type": "string"
        },
        "totalDrugCost": {
          "type": "string"
        },
        "totalDrugCostNum": {
          "type": "number"
        },
        "totalDrugcostGe65": {
          "type": "string"
        }
      }
    },
    "Records": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/MinRecord"
      }
    },
    "Register": {
      "type": "object",
      "required": [
        "email",
        "password1",
        "password2"
      ],
      "properties": {
        "email": {
          "type": "string",
          "format": "email"
        },
        "password1": {
          "type": "string"
        },
        "password2": {
          "type": "string"
        }
      }
    },
    "ReturnCode": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "readOnly": true
        },
        "message": {
          "type": "string",
          "readOnly": true
        }
      }
    },
    "User": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "token": {
          "type": "string"
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
    "description": "API used to get medicare data",
    "title": "Medicare API",
    "version": "0.0.1"
  },
  "basePath": "/go/api",
  "paths": {
    "/login": {
      "post": {
        "description": "Authenticates a user's email and password. Returns a cookie and token if successful",
        "tags": [
          "login"
        ],
        "summary": "Log ing with email and password",
        "parameters": [
          {
            "name": "login",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Login"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "User info",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "default": {
            "description": "Subscribe service unexpected error response",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          }
        }
      }
    },
    "/record": {
      "get": {
        "description": "Returns the details of a single record with a matching id",
        "tags": [
          "record"
        ],
        "summary": "Single record details",
        "parameters": [
          {
            "type": "integer",
            "description": "the service to be added",
            "name": "id",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Data based on parameters",
            "schema": {
              "$ref": "#/definitions/Record"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "default": {
            "description": "Subscribe service unexpected error response",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          }
        }
      }
    },
    "/records": {
      "get": {
        "description": "Returns a list of records based on some parameters",
        "tags": [
          "records"
        ],
        "summary": "List of records",
        "parameters": [
          {
            "type": "string",
            "name": "city",
            "in": "query"
          },
          {
            "type": "string",
            "name": "state",
            "in": "query"
          },
          {
            "type": "string",
            "name": "specialty",
            "in": "query"
          },
          {
            "type": "string",
            "name": "drugName",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Returned data based on parameters",
            "schema": {
              "$ref": "#/definitions/Records"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "default": {
            "description": "Subscribe service unexpected error response",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          }
        }
      }
    },
    "/register": {
      "post": {
        "description": "Adds a new user to the application DB",
        "tags": [
          "register"
        ],
        "summary": "Register a new user",
        "parameters": [
          {
            "name": "register",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Register"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Return success",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "default": {
            "description": "Subscribe service unexpected error response",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Login": {
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
    "MinRecord": {
      "type": "object",
      "properties": {
        "city": {
          "type": "string"
        },
        "drugName": {
          "type": "string"
        },
        "firstName": {
          "type": "string"
        },
        "genericName": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "lastOrgName": {
          "type": "string"
        },
        "specialty": {
          "type": "string"
        },
        "state": {
          "type": "string"
        }
      }
    },
    "Record": {
      "type": "object",
      "properties": {
        "beneCount": {
          "type": "string"
        },
        "beneCountGe65": {
          "type": "string"
        },
        "beneCountGe65Flag": {
          "type": "string"
        },
        "beneCountNum": {
          "type": "integer"
        },
        "city": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "drugName": {
          "type": "string"
        },
        "firstName": {
          "type": "string"
        },
        "ge65SuppressFlag": {
          "type": "string"
        },
        "genericName": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        },
        "lastOrgName": {
          "type": "string"
        },
        "npi": {
          "type": "string"
        },
        "specialty": {
          "type": "string"
        },
        "state": {
          "type": "string"
        },
        "total30DayFillCount": {
          "type": "string"
        },
        "total30DayFillCountGe65": {
          "type": "string"
        },
        "totalClaimCount": {
          "type": "string"
        },
        "totalClaimCountGe65": {
          "type": "string"
        },
        "totalClaimCountNum": {
          "type": "string"
        },
        "totalDaySupply": {
          "type": "string"
        },
        "totalDaySupplyGe65": {
          "type": "string"
        },
        "totalDrugCost": {
          "type": "string"
        },
        "totalDrugCostNum": {
          "type": "number"
        },
        "totalDrugcostGe65": {
          "type": "string"
        }
      }
    },
    "Records": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/MinRecord"
      }
    },
    "Register": {
      "type": "object",
      "required": [
        "email",
        "password1",
        "password2"
      ],
      "properties": {
        "email": {
          "type": "string",
          "format": "email"
        },
        "password1": {
          "type": "string"
        },
        "password2": {
          "type": "string"
        }
      }
    },
    "ReturnCode": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "readOnly": true
        },
        "message": {
          "type": "string",
          "readOnly": true
        }
      }
    },
    "User": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "token": {
          "type": "string"
        }
      }
    }
  }
}`))
}
