package main

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type User struct {
	ID       int    `json:"id" query:"id" params:"id"`
	Username string `json:"username" query:"username" params:"username"`
}

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive:     true,
		StrictRouting:     true,
		EnablePrintRoutes: true,
		JSONEncoder:       json.Marshal,
		JSONDecoder:       json.Unmarshal,
	})
	app.Use(cors.New(), logger.New())

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/users", func(c *fiber.Ctx) error {
		query := new(User)
		if err := c.QueryParser(query); err != nil {
			return err
		}
		return c.JSON(query)
	}).Name("send-query")

	v1.Post("/users", func(c *fiber.Ctx) error {
		body := new(User)
		if err := c.BodyParser(body); err != nil {
			return err
		}
		return c.Status(fiber.StatusCreated).JSON(body)
	}).Name("send-body")

	v1.Get("/users/:id/:username", func(c *fiber.Ctx) error {
		params := new(User)
		if err := c.ParamsParser(params); err != nil {
			return err
		}
		return c.JSON(params)
	}).Name("send-params")

	app.Get("healthcheck", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"successful": true})
	}).Name("healthcheck")
	log.Fatal(app.Listen(":3000"))
}
