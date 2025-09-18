package routes

import (
	"book_rental_app/api/handlers"
	"book_rental_app/pkg/buku"

	"github.com/gofiber/fiber/v2"
)

func BukuRoutes(api fiber.Router, bukuService *buku.BukuService) {
	api.Get("/bukus", handlers.LoadAllBuku(bukuService))
	api.Get("/bukus/:id", handlers.LoadBukuByID(bukuService))
	api.Post("/bukus", handlers.BuatBukuBaru(bukuService))
}
