package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// inisiasi framework fiber
	app := fiber.New()

	// Routing
	// URL = http://localhost:3000/api
	// Return Ahlan wa sahlan
	app.Get("/api/", func(c *fiber.Ctx) error {
		log.Print("Routing ke URL/api")
		return c.SendString("Ahlan wa Sahlan")
	})

	// 1 parameter URL
	app.Get("/api/:p1", func(c *fiber.Ctx) error {
		getParam := c.Params("p1")
		return c.SendString("Parameter yang dikirimkan adalah: " + getParam)
	})

	// 2 parameters URL
	app.Get("/api/:categoryId/:productId", func(c *fiber.Ctx) error {
		param1 := c.Params("categoryId")
		param2 := c.Params("productId")
		return c.SendString("Category ID adalah: " + param1 + ", Product ID adalah: " + param2)
	})

	// optional parameter URL
	app.Get("/api/opsional/:opsiData?", func(c *fiber.Ctx) error {
		opsiData := c.Params("opsiData", "Default Data")
		return c.SendString("Parameter yang dikirim adalah: " + opsiData)
	})

	// POST handler --> first create body
	type BodySample struct {
		Name    string `json:"name"`
		Address string `json:"address"`
	}

	app.Post("/api/info", func(c *fiber.Ctx) error {
		// initialize new object
		bodySample := new(BodySample)
		// error handling
		if err := c.BodyParser(bodySample); err != nil {
			// return error message if there is error
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		// if no error or successful parsing, send the following info
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "body sudah diterima",
			"name":    bodySample.Name,
			"address": bodySample.Address,
		})
	})

	// Reserved PORT
	app.Listen(":3000")
}
