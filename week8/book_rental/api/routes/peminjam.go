package routes

import (
	"book_rental_app/api/handlers"
	"book_rental_app/pkg/peminjam"

	"github.com/gofiber/fiber/v2"
)

func PeminjamRoutes(api fiber.Router, peminjamService *peminjam.PeminjamService) {
	api.Get("/peminjams", handlers.LoadAllPeminjam(peminjamService))
	api.Get("/peminjams/:id", handlers.LoadPeminjamByID(peminjamService))
	api.Post("/peminjams", handlers.BuatPeminjamBaru(peminjamService))
}
