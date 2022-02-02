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
    "description": "Open Camera APIs to provide a secure / open security camera solution.",
    "title": "OpenCamera",
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/device": {
      "get": {
        "security": [
          {
            "key": []
          }
        ],
        "description": "Get device information",
        "produces": [
          "application/json"
        ],
        "tags": [
          "device"
        ],
        "summary": "Get device information",
        "operationId": "getDeviceInfo",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/DeviceInfo"
            }
          },
          "401": {
            "description": "Access token is missing or invalid"
          }
        }
      }
    },
    "/media/live/session": {
      "post": {
        "security": [
          {
            "key": []
          }
        ],
        "description": "Start media session to view camera live status",
        "tags": [
          "media"
        ],
        "summary": "Start live media session",
        "operationId": "startLiveSession",
        "parameters": [
          {
            "description": "sdp data of peer connect data",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/SDP"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation"
          }
        }
      }
    },
    "/media/sdp": {
      "get": {
        "security": [
          {
            "key": []
          }
        ],
        "description": "Get camera's SDP data",
        "tags": [
          "media"
        ],
        "summary": "Get camera's SDP data",
        "operationId": "getCameraSDP",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/SDP"
            }
          }
        }
      }
    },
    "/media/stubservers": {
      "put": {
        "security": [
          {
            "key": []
          }
        ],
        "description": "Update backend stun servers address",
        "tags": [
          "media"
        ],
        "summary": "Update backend stun servers address",
        "operationId": "updateStunServers",
        "parameters": [
          {
            "description": "Created user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "array",
              "items": {
                "type": "string"
              },
              "example": [
                "stun.l.google.com:19302"
              ]
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation"
          }
        }
      }
    },
    "/media/vod/{start}/{end}": {
      "get": {
        "security": [
          {
            "key": []
          }
        ],
        "description": "Download record videos",
        "produces": [
          "application/json"
        ],
        "tags": [
          "media"
        ],
        "summary": "Download recorded videos",
        "operationId": "downloadVideos",
        "parameters": [
          {
            "type": "string",
            "description": "recording start time",
            "name": "start",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "recording end time",
            "name": "end",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation"
          }
        }
      }
    },
    "/system": {
      "get": {
        "security": [
          {
            "key": []
          }
        ],
        "description": "Get system information",
        "produces": [
          "application/json"
        ],
        "tags": [
          "system"
        ],
        "summary": "Get system information",
        "operationId": "getSystemInfo",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/SystemInfo"
            }
          },
          "401": {
            "description": "Access token is missing or invalid"
          }
        }
      }
    },
    "/system/upgrade": {
      "post": {
        "security": [
          {
            "key": []
          }
        ],
        "description": "Upgrade system firmware",
        "consumes": [
          "multipart/form-data"
        ],
        "tags": [
          "system"
        ],
        "summary": "Upgrade system firmware with uploaded file",
        "operationId": "upgradeSystem",
        "parameters": [
          {
            "type": "file",
            "description": "The firmware image to upload.",
            "name": "img_file",
            "in": "formData"
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "401": {
            "description": "Access token is missing or invalid"
          }
        }
      }
    },
    "/user": {
      "post": {
        "security": [
          {
            "key": []
          }
        ],
        "description": "Only one user is supported.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Create user",
        "operationId": "createUser",
        "parameters": [
          {
            "description": "Created user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation"
          }
        }
      }
    },
    "/user/logout": {
      "get": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "user"
        ],
        "summary": "Logs out current logged in user session",
        "operationId": "logoutUser",
        "responses": {
          "200": {
            "description": "successful operation"
          },
          "401": {
            "description": "Access token is missing or invalid"
          }
        }
      }
    },
    "/user/reset": {
      "put": {
        "security": [
          {
            "key": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Reset user password",
        "operationId": "resetPassword",
        "parameters": [
          {
            "description": "Created user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "string",
              "example": 123456
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation"
          }
        }
      }
    },
    "/user/token/obtain": {
      "post": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "tags": [
          "user"
        ],
        "summary": "Logs user into the system and obtain token",
        "operationId": "tokenObtain",
        "responses": {
          "200": {
            "description": "Encoded token",
            "schema": {
              "properties": {
                "access_token": {
                  "type": "string",
                  "format": "string"
                },
                "valid_until": {
                  "description": "token is valid until this specified date in UTC.",
                  "type": "string",
                  "format": "date",
                  "example": "2022-02-02T00:00:00Z"
                }
              }
            }
          },
          "400": {
            "description": "Invalid username/password supplied"
          }
        }
      }
    },
    "/user/token/refresh": {
      "post": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "user"
        ],
        "summary": "Refresh token",
        "operationId": "tokenRefresh",
        "responses": {
          "200": {
            "description": "New extend date",
            "schema": {
              "properties": {
                "valid_until": {
                  "description": "token is valid until this specified date in UTC.",
                  "type": "string",
                  "format": "date",
                  "example": "2022-02-02T00:00:00Z"
                }
              }
            }
          },
          "400": {
            "description": "Invalid username/password supplied"
          }
        }
      }
    }
  },
  "definitions": {
    "DeviceInfo": {
      "type": "object",
      "properties": {
        "audio_codec": {
          "description": "audio codec",
          "type": "string",
          "enum": [
            "G.711"
          ],
          "example": "G.711"
        },
        "fps": {
          "description": "frame per second",
          "type": "integer",
          "example": 30
        },
        "resolution": {
          "description": "camera resolution. Its format is \"widthxheight\"",
          "type": "string",
          "example": "1920×1080"
        },
        "uptime": {
          "description": "device uptime in unix time",
          "type": "integer",
          "format": "int64"
        },
        "uuid": {
          "type": "string",
          "example": "123e4567-e89b-12d3-a456-426614174000"
        },
        "video_codec": {
          "description": "video codec",
          "type": "string",
          "enum": [
            "H.264",
            "H.265"
          ],
          "example": "H.265"
        }
      }
    },
    "SDP": {
      "type": "object",
      "properties": {
        "data": {
          "type": "string",
          "example": ""
        },
        "valid_until": {
          "description": "data is valid until this specified date in UTC.",
          "type": "string",
          "format": "date",
          "example": "2022-02-02T00:00:00Z"
        }
      }
    },
    "SystemInfo": {
      "type": "object",
      "properties": {
        "firmwareVersion": {
          "type": "string",
          "example": "1.0.1"
        },
        "sdcard_free": {
          "description": "free sd card disk in GB",
          "type": "integer",
          "example": 10
        },
        "sdcard_total": {
          "description": "total sd card disk in GB",
          "type": "integer",
          "example": 10
        },
        "stun_servers": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "example": [
            "stun.l.google.com:19302"
          ]
        }
      }
    },
    "User": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string",
          "example": "john@email.com"
        },
        "firstName": {
          "type": "string",
          "example": "John"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "example": 10
        },
        "lastName": {
          "type": "string",
          "example": "James"
        },
        "password": {
          "type": "string",
          "example": 12345
        },
        "phone": {
          "type": "string",
          "example": "12345"
        }
      }
    }
  },
  "securityDefinitions": {
    "basicAuth": {
      "type": "basic"
    },
    "key": {
      "type": "apiKey",
      "name": "x-token",
      "in": "header"
    }
  },
  "tags": [
    {
      "description": "Everything about device itself, such as firmware version, health status",
      "name": "device"
    },
    {
      "description": "User to manage the camera",
      "name": "user"
    },
    {
      "description": "Camera audio and video operation",
      "name": "media"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "description": "Open Camera APIs to provide a secure / open security camera solution.",
    "title": "OpenCamera",
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/device": {
      "get": {
        "security": [
          {
            "key": []
          }
        ],
        "description": "Get device information",
        "produces": [
          "application/json"
        ],
        "tags": [
          "device"
        ],
        "summary": "Get device information",
        "operationId": "getDeviceInfo",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/DeviceInfo"
            }
          },
          "401": {
            "description": "Access token is missing or invalid"
          }
        }
      }
    },
    "/media/live/session": {
      "post": {
        "security": [
          {
            "key": []
          }
        ],
        "description": "Start media session to view camera live status",
        "tags": [
          "media"
        ],
        "summary": "Start live media session",
        "operationId": "startLiveSession",
        "parameters": [
          {
            "description": "sdp data of peer connect data",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/SDP"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation"
          }
        }
      }
    },
    "/media/sdp": {
      "get": {
        "security": [
          {
            "key": []
          }
        ],
        "description": "Get camera's SDP data",
        "tags": [
          "media"
        ],
        "summary": "Get camera's SDP data",
        "operationId": "getCameraSDP",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/SDP"
            }
          }
        }
      }
    },
    "/media/stubservers": {
      "put": {
        "security": [
          {
            "key": []
          }
        ],
        "description": "Update backend stun servers address",
        "tags": [
          "media"
        ],
        "summary": "Update backend stun servers address",
        "operationId": "updateStunServers",
        "parameters": [
          {
            "description": "Created user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "array",
              "items": {
                "type": "string"
              },
              "example": [
                "stun.l.google.com:19302"
              ]
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation"
          }
        }
      }
    },
    "/media/vod/{start}/{end}": {
      "get": {
        "security": [
          {
            "key": []
          }
        ],
        "description": "Download record videos",
        "produces": [
          "application/json"
        ],
        "tags": [
          "media"
        ],
        "summary": "Download recorded videos",
        "operationId": "downloadVideos",
        "parameters": [
          {
            "type": "string",
            "description": "recording start time",
            "name": "start",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "recording end time",
            "name": "end",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation"
          }
        }
      }
    },
    "/system": {
      "get": {
        "security": [
          {
            "key": []
          }
        ],
        "description": "Get system information",
        "produces": [
          "application/json"
        ],
        "tags": [
          "system"
        ],
        "summary": "Get system information",
        "operationId": "getSystemInfo",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/SystemInfo"
            }
          },
          "401": {
            "description": "Access token is missing or invalid"
          }
        }
      }
    },
    "/system/upgrade": {
      "post": {
        "security": [
          {
            "key": []
          }
        ],
        "description": "Upgrade system firmware",
        "consumes": [
          "multipart/form-data"
        ],
        "tags": [
          "system"
        ],
        "summary": "Upgrade system firmware with uploaded file",
        "operationId": "upgradeSystem",
        "parameters": [
          {
            "type": "file",
            "description": "The firmware image to upload.",
            "name": "img_file",
            "in": "formData"
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "401": {
            "description": "Access token is missing or invalid"
          }
        }
      }
    },
    "/user": {
      "post": {
        "security": [
          {
            "key": []
          }
        ],
        "description": "Only one user is supported.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Create user",
        "operationId": "createUser",
        "parameters": [
          {
            "description": "Created user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation"
          }
        }
      }
    },
    "/user/logout": {
      "get": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "user"
        ],
        "summary": "Logs out current logged in user session",
        "operationId": "logoutUser",
        "responses": {
          "200": {
            "description": "successful operation"
          },
          "401": {
            "description": "Access token is missing or invalid"
          }
        }
      }
    },
    "/user/reset": {
      "put": {
        "security": [
          {
            "key": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Reset user password",
        "operationId": "resetPassword",
        "parameters": [
          {
            "description": "Created user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "string",
              "example": 123456
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation"
          }
        }
      }
    },
    "/user/token/obtain": {
      "post": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "tags": [
          "user"
        ],
        "summary": "Logs user into the system and obtain token",
        "operationId": "tokenObtain",
        "responses": {
          "200": {
            "description": "Encoded token",
            "schema": {
              "properties": {
                "access_token": {
                  "type": "string",
                  "format": "string"
                },
                "valid_until": {
                  "description": "token is valid until this specified date in UTC.",
                  "type": "string",
                  "format": "date",
                  "example": "2022-02-02T00:00:00Z"
                }
              }
            }
          },
          "400": {
            "description": "Invalid username/password supplied"
          }
        }
      }
    },
    "/user/token/refresh": {
      "post": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "user"
        ],
        "summary": "Refresh token",
        "operationId": "tokenRefresh",
        "responses": {
          "200": {
            "description": "New extend date",
            "schema": {
              "properties": {
                "valid_until": {
                  "description": "token is valid until this specified date in UTC.",
                  "type": "string",
                  "format": "date",
                  "example": "2022-02-02T00:00:00Z"
                }
              }
            }
          },
          "400": {
            "description": "Invalid username/password supplied"
          }
        }
      }
    }
  },
  "definitions": {
    "DeviceInfo": {
      "type": "object",
      "properties": {
        "audio_codec": {
          "description": "audio codec",
          "type": "string",
          "enum": [
            "G.711"
          ],
          "example": "G.711"
        },
        "fps": {
          "description": "frame per second",
          "type": "integer",
          "example": 30
        },
        "resolution": {
          "description": "camera resolution. Its format is \"widthxheight\"",
          "type": "string",
          "example": "1920×1080"
        },
        "uptime": {
          "description": "device uptime in unix time",
          "type": "integer",
          "format": "int64"
        },
        "uuid": {
          "type": "string",
          "example": "123e4567-e89b-12d3-a456-426614174000"
        },
        "video_codec": {
          "description": "video codec",
          "type": "string",
          "enum": [
            "H.264",
            "H.265"
          ],
          "example": "H.265"
        }
      }
    },
    "SDP": {
      "type": "object",
      "properties": {
        "data": {
          "type": "string",
          "example": ""
        },
        "valid_until": {
          "description": "data is valid until this specified date in UTC.",
          "type": "string",
          "format": "date",
          "example": "2022-02-02T00:00:00Z"
        }
      }
    },
    "SystemInfo": {
      "type": "object",
      "properties": {
        "firmwareVersion": {
          "type": "string",
          "example": "1.0.1"
        },
        "sdcard_free": {
          "description": "free sd card disk in GB",
          "type": "integer",
          "example": 10
        },
        "sdcard_total": {
          "description": "total sd card disk in GB",
          "type": "integer",
          "example": 10
        },
        "stun_servers": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "example": [
            "stun.l.google.com:19302"
          ]
        }
      }
    },
    "User": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string",
          "example": "john@email.com"
        },
        "firstName": {
          "type": "string",
          "example": "John"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "example": 10
        },
        "lastName": {
          "type": "string",
          "example": "James"
        },
        "password": {
          "type": "string",
          "example": 12345
        },
        "phone": {
          "type": "string",
          "example": "12345"
        }
      }
    }
  },
  "securityDefinitions": {
    "basicAuth": {
      "type": "basic"
    },
    "key": {
      "type": "apiKey",
      "name": "x-token",
      "in": "header"
    }
  },
  "tags": [
    {
      "description": "Everything about device itself, such as firmware version, health status",
      "name": "device"
    },
    {
      "description": "User to manage the camera",
      "name": "user"
    },
    {
      "description": "Camera audio and video operation",
      "name": "media"
    }
  ]
}`))
}
