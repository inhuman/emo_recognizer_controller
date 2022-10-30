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
    "description": "Сервис контроллер для распознавателя эмоций",
    "title": "Emotions recognizer",
    "version": "{{.version}}"
  },
  "paths": {
    "/api/v1/jobs": {
      "get": {
        "description": "Эндпоинт для получения списка задач на обработку",
        "tags": [
          "Job"
        ],
        "operationId": "getJobs",
        "parameters": [
          {
            "type": "string",
            "x-go-name": "Status",
            "description": "Status",
            "name": "status",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int64",
            "x-go-name": "Limit",
            "description": "Limit",
            "name": "limit",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int64",
            "x-go-name": "Offset",
            "description": "Offset",
            "name": "offset",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/getJobsResponse"
          },
          "400": {
            "$ref": "#/responses/badDataResponse"
          },
          "500": {
            "$ref": "#/responses/internalErrorResponse"
          }
        }
      },
      "post": {
        "description": "Эндпоинт для загрузки звукового файла (.wav)",
        "consumes": [
          "multipart/form-data"
        ],
        "tags": [
          "Job"
        ],
        "operationId": "createJob",
        "parameters": [
          {
            "type": "file",
            "x-go-name": "File",
            "description": "Звуковой файл в формате wav",
            "name": "file",
            "in": "formData"
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/createJobResponse"
          },
          "400": {
            "$ref": "#/responses/badDataResponse"
          },
          "500": {
            "$ref": "#/responses/internalErrorResponse"
          }
        }
      }
    },
    "/api/v1/jobs/{Uuid}": {
      "get": {
        "description": "Эндпоинт для получения задачи по UUID",
        "tags": [
          "Job"
        ],
        "operationId": "getJob",
        "parameters": [
          {
            "type": "string",
            "description": "Uuid задания",
            "name": "Uuid",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/getJobResponse"
          },
          "400": {
            "$ref": "#/responses/badDataResponse"
          },
          "404": {
            "$ref": "#/responses/notFoundResponse"
          },
          "500": {
            "$ref": "#/responses/internalErrorResponse"
          }
        }
      }
    },
    "/api/v1/jobs/{Uuid}/file/original": {
      "get": {
        "description": "Эндпоинт для получения файла задачи по UUID",
        "produces": [
          "application/octet-stream"
        ],
        "tags": [
          "Job"
        ],
        "operationId": "getJobOriginalFile",
        "parameters": [
          {
            "type": "string",
            "description": "Uuid задания",
            "name": "Uuid",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/getJobOriginalFileResponse"
          },
          "400": {
            "$ref": "#/responses/badDataResponse"
          },
          "404": {
            "$ref": "#/responses/notFoundResponse"
          },
          "500": {
            "$ref": "#/responses/internalErrorResponse"
          }
        }
      }
    }
  },
  "definitions": {
    "Job": {
      "type": "object",
      "properties": {
        "CreatedAt": {
          "type": "string",
          "format": "date-time"
        },
        "Filename": {
          "type": "string"
        },
        "Status": {
          "$ref": "#/definitions/JobStatus"
        },
        "UUID": {
          "type": "string"
        },
        "UpdatedAt": {
          "type": "string",
          "format": "date-time"
        }
      },
      "x-go-package": "github.com/inhuman/emo_recognizer_common/jobs"
    },
    "JobStatus": {
      "type": "string",
      "x-go-package": "github.com/inhuman/emo_recognizer_common/jobs"
    },
    "Reader": {
      "description": "Read reads up to len(p) bytes into p. It returns the number of bytes\nread (0 \u003c= n \u003c= len(p)) and any error encountered. Even if Read\nreturns n \u003c len(p), it may use all of p as scratch space during the call.\nIf some data is available but not len(p) bytes, Read conventionally\nreturns what is available instead of waiting for more.\n\nWhen Read encounters an error or end-of-file condition after\nsuccessfully reading n \u003e 0 bytes, it returns the number of\nbytes read. It may return the (non-nil) error from the same call\nor return the error (and n == 0) from a subsequent call.\nAn instance of this general case is that a Reader returning\na non-zero number of bytes at the end of the input stream may\nreturn either err == EOF or err == nil. The next Read should\nreturn 0, EOF.\n\nCallers should always process the n \u003e 0 bytes returned before\nconsidering the error err. Doing so correctly handles I/O errors\nthat happen after reading some bytes and also both of the\nallowed EOF behaviors.\n\nImplementations of Read are discouraged from returning a\nzero byte count with a nil error, except when len(p) == 0.\nCallers should treat a return of 0 and nil as indicating that\nnothing happened; in particular it does not indicate EOF.\n\nImplementations must not retain p.",
      "type": "object",
      "title": "Reader is the interface that wraps the basic Read method.",
      "x-go-package": "io"
    },
    "commonErrorResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64",
          "x-go-name": "Code"
        },
        "details": {
          "type": "array",
          "items": {},
          "x-go-name": "Details"
        },
        "error": {
          "type": "string",
          "x-go-name": "Error"
        }
      },
      "x-go-package": "github.com/inhuman/emo_recognizer_controller/internal/docs"
    }
  },
  "responses": {
    "OK": {
      "description": "Success (200)"
    },
    "badDataResponse": {
      "description": "Bad data (400)",
      "schema": {
        "$ref": "#/definitions/commonErrorResponse"
      }
    },
    "createJobResponse": {
      "description": "",
      "schema": {
        "type": "object",
        "properties": {
          "UUID": {
            "type": "string"
          }
        }
      }
    },
    "getJobOriginalFileResponse": {
      "description": "",
      "schema": {
        "$ref": "#/definitions/Reader"
      }
    },
    "getJobResponse": {
      "description": "",
      "schema": {
        "$ref": "#/definitions/Job"
      }
    },
    "getJobsResponse": {
      "description": "",
      "schema": {
        "type": "array",
        "items": {
          "$ref": "#/definitions/Job"
        }
      }
    },
    "internalErrorResponse": {
      "description": "Internal server error (500)",
      "schema": {
        "$ref": "#/definitions/commonErrorResponse"
      }
    },
    "notFoundResponse": {
      "description": "Not found (404)",
      "schema": {
        "$ref": "#/definitions/commonErrorResponse"
      }
    },
    "tooManyRequestsResponse": {
      "description": "Too many requests (429)",
      "schema": {
        "$ref": "#/definitions/commonErrorResponse"
      }
    },
    "unauthorizedResponse": {
      "description": "Not authorized (401)",
      "schema": {
        "$ref": "#/definitions/commonErrorResponse"
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "description": "Сервис контроллер для распознавателя эмоций",
    "title": "Emotions recognizer",
    "version": "{{.version}}"
  },
  "paths": {
    "/api/v1/jobs": {
      "get": {
        "description": "Эндпоинт для получения списка задач на обработку",
        "tags": [
          "Job"
        ],
        "operationId": "getJobs",
        "parameters": [
          {
            "type": "string",
            "x-go-name": "Status",
            "description": "Status",
            "name": "status",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int64",
            "x-go-name": "Limit",
            "description": "Limit",
            "name": "limit",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int64",
            "x-go-name": "Offset",
            "description": "Offset",
            "name": "offset",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Job"
              }
            }
          },
          "400": {
            "description": "Bad data (400)",
            "schema": {
              "$ref": "#/definitions/commonErrorResponse"
            }
          },
          "500": {
            "description": "Internal server error (500)",
            "schema": {
              "$ref": "#/definitions/commonErrorResponse"
            }
          }
        }
      },
      "post": {
        "description": "Эндпоинт для загрузки звукового файла (.wav)",
        "consumes": [
          "multipart/form-data"
        ],
        "tags": [
          "Job"
        ],
        "operationId": "createJob",
        "parameters": [
          {
            "type": "file",
            "x-go-name": "File",
            "description": "Звуковой файл в формате wav",
            "name": "file",
            "in": "formData"
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "type": "object",
              "properties": {
                "UUID": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "description": "Bad data (400)",
            "schema": {
              "$ref": "#/definitions/commonErrorResponse"
            }
          },
          "500": {
            "description": "Internal server error (500)",
            "schema": {
              "$ref": "#/definitions/commonErrorResponse"
            }
          }
        }
      }
    },
    "/api/v1/jobs/{Uuid}": {
      "get": {
        "description": "Эндпоинт для получения задачи по UUID",
        "tags": [
          "Job"
        ],
        "operationId": "getJob",
        "parameters": [
          {
            "type": "string",
            "description": "Uuid задания",
            "name": "Uuid",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/Job"
            }
          },
          "400": {
            "description": "Bad data (400)",
            "schema": {
              "$ref": "#/definitions/commonErrorResponse"
            }
          },
          "404": {
            "description": "Not found (404)",
            "schema": {
              "$ref": "#/definitions/commonErrorResponse"
            }
          },
          "500": {
            "description": "Internal server error (500)",
            "schema": {
              "$ref": "#/definitions/commonErrorResponse"
            }
          }
        }
      }
    },
    "/api/v1/jobs/{Uuid}/file/original": {
      "get": {
        "description": "Эндпоинт для получения файла задачи по UUID",
        "produces": [
          "application/octet-stream"
        ],
        "tags": [
          "Job"
        ],
        "operationId": "getJobOriginalFile",
        "parameters": [
          {
            "type": "string",
            "description": "Uuid задания",
            "name": "Uuid",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/Reader"
            }
          },
          "400": {
            "description": "Bad data (400)",
            "schema": {
              "$ref": "#/definitions/commonErrorResponse"
            }
          },
          "404": {
            "description": "Not found (404)",
            "schema": {
              "$ref": "#/definitions/commonErrorResponse"
            }
          },
          "500": {
            "description": "Internal server error (500)",
            "schema": {
              "$ref": "#/definitions/commonErrorResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Job": {
      "type": "object",
      "properties": {
        "CreatedAt": {
          "type": "string",
          "format": "date-time"
        },
        "Filename": {
          "type": "string"
        },
        "Status": {
          "$ref": "#/definitions/JobStatus"
        },
        "UUID": {
          "type": "string"
        },
        "UpdatedAt": {
          "type": "string",
          "format": "date-time"
        }
      },
      "x-go-package": "github.com/inhuman/emo_recognizer_common/jobs"
    },
    "JobStatus": {
      "type": "string",
      "x-go-package": "github.com/inhuman/emo_recognizer_common/jobs"
    },
    "Reader": {
      "description": "Read reads up to len(p) bytes into p. It returns the number of bytes\nread (0 \u003c= n \u003c= len(p)) and any error encountered. Even if Read\nreturns n \u003c len(p), it may use all of p as scratch space during the call.\nIf some data is available but not len(p) bytes, Read conventionally\nreturns what is available instead of waiting for more.\n\nWhen Read encounters an error or end-of-file condition after\nsuccessfully reading n \u003e 0 bytes, it returns the number of\nbytes read. It may return the (non-nil) error from the same call\nor return the error (and n == 0) from a subsequent call.\nAn instance of this general case is that a Reader returning\na non-zero number of bytes at the end of the input stream may\nreturn either err == EOF or err == nil. The next Read should\nreturn 0, EOF.\n\nCallers should always process the n \u003e 0 bytes returned before\nconsidering the error err. Doing so correctly handles I/O errors\nthat happen after reading some bytes and also both of the\nallowed EOF behaviors.\n\nImplementations of Read are discouraged from returning a\nzero byte count with a nil error, except when len(p) == 0.\nCallers should treat a return of 0 and nil as indicating that\nnothing happened; in particular it does not indicate EOF.\n\nImplementations must not retain p.",
      "type": "object",
      "title": "Reader is the interface that wraps the basic Read method.",
      "x-go-package": "io"
    },
    "commonErrorResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64",
          "x-go-name": "Code"
        },
        "details": {
          "type": "array",
          "items": {},
          "x-go-name": "Details"
        },
        "error": {
          "type": "string",
          "x-go-name": "Error"
        }
      },
      "x-go-package": "github.com/inhuman/emo_recognizer_controller/internal/docs"
    }
  },
  "responses": {
    "OK": {
      "description": "Success (200)"
    },
    "badDataResponse": {
      "description": "Bad data (400)",
      "schema": {
        "$ref": "#/definitions/commonErrorResponse"
      }
    },
    "createJobResponse": {
      "description": "",
      "schema": {
        "type": "object",
        "properties": {
          "UUID": {
            "type": "string"
          }
        }
      }
    },
    "getJobOriginalFileResponse": {
      "description": "",
      "schema": {
        "$ref": "#/definitions/Reader"
      }
    },
    "getJobResponse": {
      "description": "",
      "schema": {
        "$ref": "#/definitions/Job"
      }
    },
    "getJobsResponse": {
      "description": "",
      "schema": {
        "type": "array",
        "items": {
          "$ref": "#/definitions/Job"
        }
      }
    },
    "internalErrorResponse": {
      "description": "Internal server error (500)",
      "schema": {
        "$ref": "#/definitions/commonErrorResponse"
      }
    },
    "notFoundResponse": {
      "description": "Not found (404)",
      "schema": {
        "$ref": "#/definitions/commonErrorResponse"
      }
    },
    "tooManyRequestsResponse": {
      "description": "Too many requests (429)",
      "schema": {
        "$ref": "#/definitions/commonErrorResponse"
      }
    },
    "unauthorizedResponse": {
      "description": "Not authorized (401)",
      "schema": {
        "$ref": "#/definitions/commonErrorResponse"
      }
    }
  }
}`))
}
