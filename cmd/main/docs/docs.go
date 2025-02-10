// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import (
    "github.com/swaggo/swag"
    "os"
)

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/getsong": {
            "post": {
                "description": "Возвращает информацию о музыкальном треке по переданной ссылке.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Получение информации о треке",
                "parameters": [
                    {
                        "description": "JSON URL трека",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/app.GetSongRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Информация о треке",
                        "schema": {
                            "$ref": "#/definitions/track_info.TrackInfo"
                        }
                    },
                    "400": {
                        "description": "Некорректный запрос",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Трек не найден",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app.GetSongRequest": {
            "type": "object",
            "properties": {
                "url": {
                    "type": "string",
                    "example": "https://music.apple.com/song/697195787"
                }
            }
        },
        "gin.H": {
            "type": "object",
            "additionalProperties": {}
        },
        "track_info.TrackInfo": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "properties": {
                        "artist": {
                            "type": "string"
                        },
                        "cover_url": {
                            "type": "string"
                        },
                        "id": {
                            "type": "integer"
                        },
                        "links": {
                            "type": "object",
                            "properties": {
                                "apple_music_url": {
                                    "type": "string"
                                },
                                "deezer_url": {
                                    "type": "string"
                                },
                                "spotify_url": {
                                    "type": "string"
                                },
                                "vk_music_url": {
                                    "type": "string"
                                },
                                "yandex_music_url": {
                                    "type": "string"
                                }
                            }
                        },
                        "preview": {
                            "type": "string"
                        },
                        "title": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          getEnv("SWAGGER_VERSION", "2.0"),
	Host:             getEnv("SWAGGER_HOST", "localhost:8088"),
	BasePath:         getEnv("SWAGGER_BASE_PATH", "/api"),
	Schemes:          []string{},
	Title:            getEnv("SWAGGER_TITLE", "Swagger API"),
	Description:      getEnv("SWAGGER_DESCRIPTION", "This is a sample server celler server."),
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}