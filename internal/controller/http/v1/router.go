package v1

import (
	docs "github.com/flaiers/fiber-clean-architecture/docs/v1"
	"github.com/flaiers/fiber-clean-architecture/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title          Fiber Clean API
// @version        1.0
// @description    Fiber Clean REST API
// @termsOfService https://swagger.io/terms

// @contact.name  Maxim Bigin
// @contact.email i@flaiers.me
// @contact.url   https://flaiers.me

// @license.name Apache 2.0
// @license.url  https://www.apache.org/licenses/LICENSE-2.0

// @basePath /api/v1

// @securityDefinitions.basic  BasicAuth
// @securityDefinitions.apikey ApiKeyAuth
// @in                         header
// @name                       Authorization

func Router(cfg config.Config) func(fiber.Router) {
	return func(router fiber.Router) {
		router.Get("/swagger/*", swagger.New(
			config.NewSwaggerConfig(docs.SwaggerInfo),
		))
	}
}
