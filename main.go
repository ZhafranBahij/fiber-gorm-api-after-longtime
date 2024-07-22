package main

import (
	"yahallo/controller/hello_world_controller"
	"yahallo/controller/product_controller"
	"yahallo/db"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// Setup DB
	db.ConnectDB()

	app := fiber.New()

	app.Get("/", hello_world_controller.HelloWorld)
	app.Get("/product", product_controller.IndexProduct)
	app.Post("/product", product_controller.CreateProduct)
	app.Get("/product/:id", product_controller.ShowProduct)
	app.Put("/product/:id", product_controller.UpdateProduct)
	app.Delete("/product/:id", product_controller.DeleteProduct)

	app.Listen(":3000")
}
