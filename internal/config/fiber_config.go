package config

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewFiberConfig() fiber.Config {
	conf := fiber.Config{
		ServerHeader:       "Api-Project",
		AppName:            "Api-Project",
		DisableDefaultDate: true,
		RequestMethods: []string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPatch,
			fiber.MethodDelete,
			fiber.MethodOptions,
		},
		Concurrency: 256 * 1024 * 4,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		BodyLimit:   2 * 1024 * 1024 * 1024,
	}

	return conf
}
func CompressResponseConfig() compress.Config {
	return compress.Config{
		Level: compress.LevelBestCompression,
	}
}

func LoggerConfig() logger.Config {
	return logger.Config{
		Format:     "[${ip}]:${port} ${time} ${status} - ${method} ${path} ${time}\n",
		TimeFormat: "02-12-2022",
		TimeZone:   "Asia/Kathmandu",
	}
}
func CorsConfig() cors.Config {
	return cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: false,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}
}
