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
  "swagger": "2.0",
  "info": {
    "description": "This is a sample notes-rew server.",
    "title": "Notes-Service API",
    "contact": {},
    "version": "1.0"
  },
  "host": "localhost:8081",
  "basePath": "/",
  "paths": {
    "/auth/login": {
      "post": {
        "description": "login user",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "auth"
        ],
        "summary": "SignIn",
        "parameters": [
          {
            "description": "User info",
            "name": "user",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/controller.SignInRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Bad Request"
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      }
    },
    "/auth/register": {
      "post": {
        "description": "create user",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "auth"
        ],
        "summary": "SignUp",
        "parameters": [
          {
            "description": "User info",
            "name": "user",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/controller.SignUpRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created"
          },
          "400": {
            "description": "Bad Request"
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      }
    },
    "/notes": {
      "get": {
        "security": [
          {
            "JWTAuth": []
          }
        ],
        "description": "get all notes",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "notes"
        ],
        "summary": "GetAllNotes",
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Bad Request"
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      },
      "post": {
        "security": [
          {
            "JWTAuth": []
          }
        ],
        "description": "create note",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "notes"
        ],
        "summary": "CreateNote",
        "parameters": [
          {
            "description": "Note info",
            "name": "note",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/controller.CreateNoteRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created"
          },
          "400": {
            "description": "Bad Request"
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      }
    },
    "/notes/{id}": {
      "get": {
        "security": [
          {
            "JWTAuth": []
          }
        ],
        "description": "get note",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "notes"
        ],
        "summary": "GetNote",
        "parameters": [
          {
            "type": "string",
            "description": "Note ID",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Bad Request"
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      },
      "delete": {
        "security": [
          {
            "JWTAuth": []
          }
        ],
        "description": "delete note",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "notes"
        ],
        "summary": "DeleteNote",
        "parameters": [
          {
            "type": "string",
            "description": "Note ID",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Bad Request"
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      },
      "patch": {
        "security": [
          {
            "JWTAuth": []
          }
        ],
        "description": "update note",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "notes"
        ],
        "summary": "UpdateNote",
        "parameters": [
          {
            "type": "string",
            "description": "Note ID",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "Note info",
            "name": "note",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/controller.UpdateNoteRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Bad Request"
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      }
    },
    "/users": {
      "get": {
        "security": [
          {
            "JWTAuth": []
          }
        ],
        "description": "get user",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "users"
        ],
        "summary": "GetUser",
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Bad Request"
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      },
      "delete": {
        "security": [
          {
            "JWTAuth": []
          }
        ],
        "description": "delete user",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "users"
        ],
        "summary": "DeleteUser",
        "responses": {
          "204": {
            "description": "No Content"
          },
          "400": {
            "description": "Bad Request"
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      },
      "patch": {
        "security": [
          {
            "JWTAuth": []
          }
        ],
        "description": "update user",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "users"
        ],
        "summary": "UpdateUser",
        "parameters": [
          {
            "description": "User info",
            "name": "user",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/controller.UpdateUserRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Bad Request"
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      }
    }
  },
  "definitions": {
    "controller.CreateNoteRequest": {
      "type": "object",
      "required": [
        "body",
        "title"
      ],
      "properties": {
        "body": {
          "type": "string"
        },
        "tags": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "title": {
          "type": "string",
          "maxLength": 50,
          "minLength": 1
        }
      }
    },
    "controller.SignInRequest": {
      "type": "object",
      "required": [
        "email",
        "password"
      ],
      "properties": {
        "email": {
          "type": "string",
          "maxLength": 254,
          "minLength": 5
        },
        "password": {
          "type": "string",
          "maxLength": 30,
          "minLength": 6
        }
      }
    },
    "controller.SignUpRequest": {
      "type": "object",
      "required": [
        "email",
        "password",
        "username"
      ],
      "properties": {
        "email": {
          "type": "string",
          "maxLength": 254,
          "minLength": 5
        },
        "password": {
          "type": "string"
        },
        "username": {
          "type": "string",
          "maxLength": 20,
          "minLength": 3
        }
      }
    },
    "controller.UpdateNoteRequest": {
      "type": "object",
      "required": [
        "body",
        "title"
      ],
      "properties": {
        "body": {
          "type": "string"
        },
        "tags": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "title": {
          "type": "string",
          "maxLength": 50,
          "minLength": 1
        },
        "updated_at": {
          "type": "string"
        }
      }
    },
    "controller.UpdateUserRequest": {
      "type": "object",
      "required": [
        "email",
        "password"
      ],
      "properties": {
        "email": {
          "type": "string",
          "maxLength": 254,
          "minLength": 5
        },
        "password": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "JWTAuth": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "description": "This is a sample notes-rew server.",
    "title": "Notes-Service API",
    "contact": {},
    "version": "1.0"
  },
  "host": "localhost:8081",
  "basePath": "/",
  "paths": {
    "/auth/login": {
      "post": {
        "description": "login user",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "auth"
        ],
        "summary": "SignIn",
        "parameters": [
          {
            "description": "User info",
            "name": "user",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/controller.SignInRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Bad Request"
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      }
    },
    "/auth/register": {
      "post": {
        "description": "create user",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "auth"
        ],
        "summary": "SignUp",
        "parameters": [
          {
            "description": "User info",
            "name": "user",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/controller.SignUpRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created"
          },
          "400": {
            "description": "Bad Request"
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      }
    },
    "/notes": {
      "get": {
        "security": [
          {
            "JWTAuth": []
          }
        ],
        "description": "get all notes",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "notes"
        ],
        "summary": "GetAllNotes",
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Bad Request"
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      },
      "post": {
        "security": [
          {
            "JWTAuth": []
          }
        ],
        "description": "create note",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "notes"
        ],
        "summary": "CreateNote",
        "parameters": [
          {
            "description": "Note info",
            "name": "note",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/controller.CreateNoteRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created"
          },
          "400": {
            "description": "Bad Request"
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      }
    },
    "/notes/{id}": {
      "get": {
        "security": [
          {
            "JWTAuth": []
          }
        ],
        "description": "get note",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "notes"
        ],
        "summary": "GetNote",
        "parameters": [
          {
            "type": "string",
            "description": "Note ID",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Bad Request"
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      },
      "delete": {
        "security": [
          {
            "JWTAuth": []
          }
        ],
        "description": "delete note",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "notes"
        ],
        "summary": "DeleteNote",
        "parameters": [
          {
            "type": "string",
            "description": "Note ID",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Bad Request"
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      },
      "patch": {
        "security": [
          {
            "JWTAuth": []
          }
        ],
        "description": "update note",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "notes"
        ],
        "summary": "UpdateNote",
        "parameters": [
          {
            "type": "string",
            "description": "Note ID",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "Note info",
            "name": "note",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/controller.UpdateNoteRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Bad Request"
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      }
    },
    "/users": {
      "get": {
        "security": [
          {
            "JWTAuth": []
          }
        ],
        "description": "get user",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "users"
        ],
        "summary": "GetUser",
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Bad Request"
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      },
      "delete": {
        "security": [
          {
            "JWTAuth": []
          }
        ],
        "description": "delete user",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "users"
        ],
        "summary": "DeleteUser",
        "responses": {
          "204": {
            "description": "No Content"
          },
          "400": {
            "description": "Bad Request"
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      },
      "patch": {
        "security": [
          {
            "JWTAuth": []
          }
        ],
        "description": "update user",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "users"
        ],
        "summary": "UpdateUser",
        "parameters": [
          {
            "description": "User info",
            "name": "user",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/controller.UpdateUserRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Bad Request"
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      }
    }
  },
  "definitions": {
    "controller.CreateNoteRequest": {
      "type": "object",
      "required": [
        "body",
        "title"
      ],
      "properties": {
        "body": {
          "type": "string"
        },
        "tags": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "title": {
          "type": "string",
          "maxLength": 50,
          "minLength": 1
        }
      }
    },
    "controller.SignInRequest": {
      "type": "object",
      "required": [
        "email",
        "password"
      ],
      "properties": {
        "email": {
          "type": "string",
          "maxLength": 254,
          "minLength": 5
        },
        "password": {
          "type": "string",
          "maxLength": 30,
          "minLength": 6
        }
      }
    },
    "controller.SignUpRequest": {
      "type": "object",
      "required": [
        "email",
        "password",
        "username"
      ],
      "properties": {
        "email": {
          "type": "string",
          "maxLength": 254,
          "minLength": 5
        },
        "password": {
          "type": "string"
        },
        "username": {
          "type": "string",
          "maxLength": 20,
          "minLength": 3
        }
      }
    },
    "controller.UpdateNoteRequest": {
      "type": "object",
      "required": [
        "body",
        "title"
      ],
      "properties": {
        "body": {
          "type": "string"
        },
        "tags": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "title": {
          "type": "string",
          "maxLength": 50,
          "minLength": 1
        },
        "updated_at": {
          "type": "string"
        }
      }
    },
    "controller.UpdateUserRequest": {
      "type": "object",
      "required": [
        "email",
        "password"
      ],
      "properties": {
        "email": {
          "type": "string",
          "maxLength": 254,
          "minLength": 5
        },
        "password": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "JWTAuth": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
}
