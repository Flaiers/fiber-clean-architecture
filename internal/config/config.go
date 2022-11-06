package config

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/flaiers/fiber-clean-architecture/internal/usecase/utils"

	"github.com/caarlos0/env/v6"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/swaggo/swag"
)

const (
	API = "/api"
)

type Config struct {
	Port string `env:"BACKEND_PORT"`
	Host string `env:"BACKEND_HOST"`
	Addr string

	CorsOrigins string `env:"BACKEND_CORS_ORIGINS"`

	Database struct {
		Host     string `env:"DB_HOST"`
		Port     string `env:"DB_PORT"`
		User     string `env:"DB_USER"`
		Pass     string `env:"DB_PASS"`
		Name     string `env:"DB_NAME"`
		TimeZone string `env:"DB_TIMEZONE" envDefault:"UTC"`
		SSLMode  string `env:"DB_SSL_MODE" envDefault:"disable"`
		Dsn      string
	}

	S3 struct {
		Region          string        `env:"S3_REGION"`
		Bucket          string        `env:"S3_BUCKET"`
		Endpoint        string        `env:"S3_ENDPOINT"`
		AccessKey       string        `env:"S3_ACCESS_KEY"`
		SecretAccessKey string        `env:"S3_SECRET_ACCESS_KEY"`
		MaxAttempts     int           `env:"S3_MAX_ATTEMPTS" envDefault:"3"`
		RequestTimeout  time.Duration `env:"S3_REQUEST_TIMEOUT" envDefault:"0"`
	}
}

func NewConfig() (Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return cfg, err
	}

	if cfg.Addr == "" {
		cfg.Addr = fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	}

	if cfg.Database.Dsn == "" {
		cfg.Database.Dsn = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
			cfg.Database.Host,
			cfg.Database.Port,
			cfg.Database.User,
			cfg.Database.Pass,
			cfg.Database.Name,
			cfg.Database.SSLMode,
			cfg.Database.TimeZone,
		)
	}

	return cfg, nil
}

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		ErrorHandler:  utils.JSONErrorHandler,
	}
}

func NewCorsConfig(cfg Config) cors.Config {
	return cors.Config{
		AllowOrigins:     cfg.CorsOrigins,
		AllowMethods:     "*",
		AllowHeaders:     "*",
		AllowCredentials: true,
	}
}

func NewLoggerConfig() logger.Config {
	return logger.ConfigDefault
}

func NewSwaggerConfig(s *swag.Spec) swagger.Config {
	return swagger.Config{Title: s.Title}
}
