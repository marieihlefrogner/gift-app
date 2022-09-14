package main

import (
	"gift-app/config"
	"gift-app/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// https://dev.to/franciscomendes10866/how-to-build-rest-api-using-go-fiber-and-gorm-orm-2jbe

func main() {
	app := fiber.New()

	config.Connect()

	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/gifts/:id", handlers.GetGift)
	app.Post("/gifts/:id/use", handlers.UseGift)

	admin := app.Group("/admin")
	admin.Get("/gifts", handlers.GetGifts)
	admin.Post("/gifts", handlers.CreateGift)
	admin.Get("/gifts/:id", handlers.GetGift)
	admin.Put("/gifts/:id", handlers.UpdateGift)
	admin.Delete("/gifts/:id", handlers.RemoveGift)

	app.Listen(":8080")
}
