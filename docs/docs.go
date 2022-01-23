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
        "termsOfService": "/terms",
        "contact": {},
        "license": {
            "name": "MIT",
            "url": "/license"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/getmoves": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GameAnalysis"
                ],
                "summary": "Get moves for a given board layout and dice roll",
                "parameters": [
                    {
                        "description": "Move arguments",
                        "name": "args",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/api.MoveArgs"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.Move"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.Board": {
            "type": "object",
            "properties": {
                "o": {
                    "$ref": "#/definitions/api.CheckerLayout"
                },
                "x": {
                    "$ref": "#/definitions/api.CheckerLayout"
                }
            }
        },
        "api.CheckerLayout": {
            "type": "object",
            "properties": {
                "1": {
                    "type": "integer"
                },
                "10": {
                    "type": "integer"
                },
                "11": {
                    "type": "integer"
                },
                "12": {
                    "type": "integer"
                },
                "13": {
                    "type": "integer",
                    "example": 5
                },
                "14": {
                    "type": "integer"
                },
                "15": {
                    "type": "integer"
                },
                "16": {
                    "type": "integer"
                },
                "17": {
                    "type": "integer"
                },
                "18": {
                    "type": "integer"
                },
                "19": {
                    "type": "integer"
                },
                "2": {
                    "type": "integer"
                },
                "20": {
                    "type": "integer"
                },
                "21": {
                    "type": "integer"
                },
                "22": {
                    "type": "integer"
                },
                "23": {
                    "type": "integer"
                },
                "24": {
                    "type": "integer",
                    "example": 2
                },
                "3": {
                    "type": "integer"
                },
                "4": {
                    "type": "integer"
                },
                "5": {
                    "type": "integer"
                },
                "6": {
                    "type": "integer",
                    "example": 5
                },
                "7": {
                    "type": "integer"
                },
                "8": {
                    "type": "integer",
                    "example": 3
                },
                "9": {
                    "type": "integer"
                },
                "bar": {
                    "type": "integer"
                }
            }
        },
        "api.CheckerPlay": {
            "type": "object",
            "properties": {
                "from": {
                    "type": "string",
                    "enum": [
                        "1",
                        "2",
                        "3",
                        "4",
                        "5",
                        "6",
                        "7",
                        "8",
                        "9",
                        "10",
                        "11",
                        "12",
                        "13",
                        "14",
                        "15",
                        "16",
                        "17",
                        "18",
                        "19",
                        "20",
                        "21",
                        "22",
                        "23",
                        "24",
                        "bar"
                    ]
                },
                "to": {
                    "type": "string",
                    "enum": [
                        "1",
                        "2",
                        "3",
                        "4",
                        "5",
                        "6",
                        "7",
                        "8",
                        "9",
                        "10",
                        "11",
                        "12",
                        "13",
                        "14",
                        "15",
                        "16",
                        "17",
                        "18",
                        "19",
                        "20",
                        "21",
                        "22",
                        "23",
                        "24",
                        "off"
                    ]
                }
            }
        },
        "api.EvalInfo": {
            "type": "object",
            "properties": {
                "cubeful": {
                    "type": "boolean"
                },
                "plies": {
                    "type": "integer"
                }
            }
        },
        "api.Evaluation": {
            "type": "object",
            "properties": {
                "diff": {
                    "type": "number"
                },
                "eq": {
                    "type": "number"
                },
                "info": {
                    "$ref": "#/definitions/api.EvalInfo"
                },
                "probability": {
                    "$ref": "#/definitions/api.Probablity"
                }
            }
        },
        "api.Move": {
            "type": "object",
            "properties": {
                "evaluation": {
                    "$ref": "#/definitions/api.Evaluation"
                },
                "play": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.CheckerPlay"
                    }
                }
            }
        },
        "api.MoveArgs": {
            "type": "object",
            "properties": {
                "board": {
                    "$ref": "#/definitions/api.Board"
                },
                "cubeful": {
                    "type": "boolean"
                },
                "dice": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    },
                    "example": [
                        3,
                        1
                    ]
                },
                "max-moves": {
                    "type": "integer",
                    "default": 0,
                    "minimum": 0
                },
                "player": {
                    "type": "string",
                    "default": "x",
                    "enum": [
                        "x",
                        " o"
                    ],
                    "example": "x"
                },
                "score-moves": {
                    "type": "boolean",
                    "default": true,
                    "example": true
                }
            }
        },
        "api.Probablity": {
            "type": "object",
            "properties": {
                "lose": {
                    "type": "number"
                },
                "loseBG": {
                    "type": "number"
                },
                "loseG": {
                    "type": "number"
                },
                "win": {
                    "type": "number"
                },
                "winBG": {
                    "type": "number"
                },
                "winG": {
                    "type": "number"
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
	Version:     "1.0",
	Host:        "",
	BasePath:    "/api/v1",
	Schemes:     []string{"http"},
	Title:       "BGWeb API",
	Description: "BGWeb API",
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