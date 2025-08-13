package handlers

import (
	"fiber_go/database"
	"fiber_go/models"

	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT id, name, description, price, stock, image_url FROM products")

	if err != nil {
		return c.Status(500).SendString("Ошибка выполнения запроса к базе данных")
	}

	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var product models.Product

		err := rows.Scan(&product.Description, &product.ID, &product.ImageURL, &product.Name, &product.Price, &product.Stock)

		if err != nil {
			return c.Status(500).SendString("Ошибка сканирования данных")
		}

		products = append(products, product)

	}

	return c.JSON(products)
}

func CreateProduct(c *fiber.Ctx) error {
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	_, err := database.DB.Exec("INSERT INTO products (name, description, price, stock, image_url) VALUES ($1, $2, $3, $4, $5)",
		product.Name, product.Description, product.Price, product.Stock, product.ImageURL)

	if err != nil {
		return c.Status(500).SendString("Ошибка вставки данных в базу")
	}
	return c.Status(201).SendString("Продукт успешно создан")
}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	row := database.DB.QueryRow("SELECT id, name, description, price, stock, image_url FROM products WHERE id = $1", id)

	var product models.Product
	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.ImageURL)
	if err != nil {
		return c.Status(404).SendString("Продукт не найден")
	}
	return c.JSON(product)

}

func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	product := new(models.Product)

	if err := c.BodyParser(product); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}

	_, err := database.DB.Exec("UPDATE products SET name = $1, description = $2, price = $3, stock = $4, image_url = $5 WHERE id = $6",
		product.Name, product.Description, product.Price, product.Stock, product.ImageURL, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления данных")
	}

	return c.SendString("Продукт успешно обновлён")
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := database.DB.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return c.Status(500).SendString("Ошибка удаления продукта")
	}

	return c.SendString("Продукт успешно удалён")
}
