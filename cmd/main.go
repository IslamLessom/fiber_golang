package main

import (
	"fiber_go/database"
	route "fiber_go/routes"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatal("Подключение к базе данных")
	}

	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Use(logger.New())
	app.Use(compress.New())
	app.Use(recover.New())
	app.Use(limiter.New())

	route.RegisterProductRoutes(app)

	logrus.Fatal(app.Listen(":80"))
}
