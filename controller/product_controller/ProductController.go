package product_controller

import (
	"yahallo/db"
	"yahallo/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func IndexProduct(c *fiber.Ctx) error {
	// Declare variable
	var products []model.Product

	// GET all data from DB
	result := db.DB.Find(&products)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	// Return data to API
	return c.Status(fiber.StatusOK).JSON(products)
}

func CreateProduct(c *fiber.Ctx) error {
	// Declare variable
	var product model.Product

	// Save the input value
	err := c.BodyParser(&product)
	if err != nil {
		return err
	}

	// Send input data to DB
	db.DB.Create(&product)

	// Return status & Value
	return c.Status(fiber.StatusCreated).JSON(product)
}

func ShowProduct(c *fiber.Ctx) error {

	// GET data from DATABASE based on id
	var product model.Product
	id := c.Params("id")
	result := db.DB.First(&product, id)

	// If data not found
	if result.Error == gorm.ErrRecordNotFound {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "product not found"})
	}

	// If has another error
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	// Return data
	return c.Status(fiber.StatusOK).JSON(&product)
}

func UpdateProduct(c *fiber.Ctx) error {
	// GET data from DATABASE based on id
	var product model.Product
	var update_product model.Product
	id := c.Params("id")

	result := db.DB.First(&product, id)
	if result.Error == gorm.ErrRecordNotFound {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "book not found"})
	}

	// Update the input value
	err := c.BodyParser(&update_product)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Update data in database
	db.DB.Model(&product).Updates(update_product)

	// Return status & Value
	return c.Status(fiber.StatusCreated).JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {

	// GET data from DATABASE based on id
	var product model.Product
	id := c.Params("id")
	result := db.DB.First(&product, id)

	// If data not found
	if result.Error == gorm.ErrRecordNotFound {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "product not found"})
	}

	// If has another error
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	// Delete product from database
	db.DB.Delete(&product)

	// Return data
	return c.Status(fiber.StatusOK).JSON(&product)
}
