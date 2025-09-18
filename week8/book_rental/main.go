package main

import (
	"book_rental_app/api/routes"
	"book_rental_app/pkg/buku"
	"book_rental_app/pkg/entities"
	"book_rental_app/pkg/peminjam"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 1. konfigurasi database dan inisialisasi repository serta service
	dsn := "coba:Password123!@tcp(172.16.202.130:3306)/hsisandbox?charset=utf8mb4&parseTime=True&loc=Local"
	koneksiDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Koneksi Gagal, terjadi kesalahan pada saat berkomunikasi dengan DB")
	}
	koneksiDb.AutoMigrate(&entities.Buku{})
	koneksiDb.AutoMigrate(&entities.Peminjam{})

	// 2. inisialisasi repository
	bukuRepository := buku.NewBukuRepository(koneksiDb)
	peminjamRepository := peminjam.NewPeminjamRepository(koneksiDb)

	// 3. inisialisasi service
	bukuService := buku.NewBukuService(bukuRepository)
	peminjamService := peminjam.NewPeminjamService(peminjamRepository)

	// 4. inisialisasi fiber dan setup routes
	app := fiber.New()

	api := app.Group("/api")
	routes.BukuRoutes(api, bukuService)
	routes.PeminjamRoutes(api, peminjamService)

	app.Listen(":3000")
}
