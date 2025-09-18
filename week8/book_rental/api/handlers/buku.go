package handlers

import (
	"book_rental_app/pkg/buku"
	"book_rental_app/pkg/entities"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func LoadAllBuku(service *buku.BukuService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		bukus, err := service.GetAllBuku()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to load books",
			})

		}
		return c.JSON(bukus)
	}
}

func LoadBukuByID(service *buku.BukuService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		getID := c.Params("id")
		ID, errID := strconv.Atoi(getID)
		if errID != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error converting string to int",
			})
		}
		buku, err := service.GetBukuByID(ID)
		if err != nil {
			errorMsg := fmt.Sprint("Failed to load book with ID: ", ID)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": errorMsg,
			})
		}
		return c.JSON(buku)
	}
}

func BuatBukuBaru(service *buku.BukuService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var buku entities.Buku
		if err := c.BodyParser(&buku); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request",
			})
		}
		if err := service.CreateBuku(buku); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create book",
			})
		}
		return c.Status(fiber.StatusCreated).JSON(buku)
	}
}
