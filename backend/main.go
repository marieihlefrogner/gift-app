package main

import (
	"gift-app/config"
	"gift-app/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"os"
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

	api := app.Group("/api")

	api.Get("/gifts/:id", handlers.GetGift)
	api.Post("/gifts/:id/use", handlers.UseGift)

	admin := api.Group("/admin")
	admin.Get("/gifts", handlers.GetGifts)
	admin.Post("/gifts", handlers.CreateGift)
	admin.Get("/gifts/:id", handlers.GetGift)
	admin.Put("/gifts/:id", handlers.UpdateGift)
	admin.Delete("/gifts/:id", handlers.RemoveGift)

	app.Static("/", "./public")
	app.Static("*", "./public/index.html")

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on port %s\n\n", port)

	app.Listen(":" + port)
}
