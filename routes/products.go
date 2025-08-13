package routes

import (
	"fiber_go/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterProductRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/products", handlers.GetProducts)          // Получить все продукты
	api.Post("/products", handlers.CreateProduct)       // Создать новый продукт
	api.Get("/products/:id", handlers.GetProduct)       // Получить продукт по ID
	api.Put("/products/:id", handlers.UpdateProduct)    // Обновить продукт
	api.Delete("/products/:id", handlers.DeleteProduct) // Удалить продукт
}
