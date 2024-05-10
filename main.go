package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func HelloWorld(ctx *fiber.Ctx) error {

	return ctx.SendString("Hello World!")

}

func main() {

	app := fiber.New()

	rateLimiter := limiter.New(limiter.Config{
		Max:        5,
		Expiration: time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
	})

	app.Use(rateLimiter)
	app.Get("/", HelloWorld)

	app.Listen("localhost:3000")

}
