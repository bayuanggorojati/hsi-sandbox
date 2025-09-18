package handlers

import (
	"book_rental_app/pkg/entities"
	"book_rental_app/pkg/peminjam"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func LoadAllPeminjam(service *peminjam.PeminjamService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		peminjams, err := service.GetAllPeminjam()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to load borrowers",
			})
		}
		return c.JSON(peminjams)
	}
}

func LoadPeminjamByID(service *peminjam.PeminjamService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		getID := c.Params("id")
		ID, errID := strconv.Atoi(getID)
		if errID != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error converting string to int",
			})
		}
		peminjam, err := service.GetPeminjamByID(ID)
		if err != nil {
			errorMsg := fmt.Sprint("Failed to load borrower with ID: ", ID)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": errorMsg,
			})
		}
		return c.JSON(peminjam)
	}
}

func BuatPeminjamBaru(service *peminjam.PeminjamService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var peminjam entities.Peminjam
		if err := c.BodyParser(&peminjam); err != nil {
			//log.Println("Nama: ", peminjam.Nama)
			//log.Println("Tanggal peminjaman", peminjam.TanggalPeminjaman)
			//log.Println("Tanggal pengembalian", peminjam.TanggalPengembalian)
			//log.Println("Buku pinjaman:", peminjam.BukuPinjaman)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request",
			})
		}
		//log.Println("Nama: ", peminjam.Nama)
		//log.Println("Tanggal peminjaman", peminjam.TanggalPeminjaman)
		log.Println("Tanggal pengembalian", peminjam.TanggalPengembalian)
		log.Println("Buku pinjaman:", peminjam.BukuPinjaman)
		if err := service.CreatePeminjam(peminjam); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create a borrower",
			})
		}
		return c.Status(fiber.StatusCreated).JSON(peminjam)
	}
}
