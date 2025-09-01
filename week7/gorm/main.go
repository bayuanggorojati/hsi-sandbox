package main

import (
	// Import Library ORM GORM
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	// Import Aplikasi Standar

	"log"
)

// Struktur Table -> Table ukur_suhus
type UkurSuhu struct {
	ID             uint `gorm:"primaryKey"`
	SuhuCelcius    float64
	SuhuFahrenheit float64
	SuhuReamur     float64
}

func main() {
	// 1. Mendefinisikan Koneksi Database
	dsn := "root:Active123@tcp(localhost:6603)/hsi_sandbox?charset=utf8mb4&parseTime=True&loc=Local"
	// 2. Membuat Koneksi Database
	koneksiDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Koneksi Gagal, terjadi kesalahan pada saat berkomunikasi dengan DB", err.Error())
	}
	//3. Migrasi Struktur Table dari Struct
	koneksiDb.AutoMigrate(&UkurSuhu{})
	//4. Operasi CRUD -> Create / Insert Data
	fmt.Println("Menambahkan Data Ukur Suhu Baru")
	newData := UkurSuhu{
		SuhuCelcius:    30,
		SuhuFahrenheit: 86,
		SuhuReamur:     24,
	}
	operasiTambahData := koneksiDb.Create(&newData)
	if operasiTambahData.Error != nil {
		log.Fatal("Gagal menambahkan data baru", operasiTambahData.Error.Error())
	}
	fmt.Println("Berhasil menambahkan data baru dengan ID:", newData.ID)

	//5. Operasi CRUD -> Read / Get Data
	fmt.Println("Menampilkan Data Ukur Suhu")
	// Menggunakan Function First untuk mendapatkan 1 data pertama
	var bacaData UkurSuhu
	koneksiDb.First(&bacaData)
	fmt.Println("Data Ukur Suhu Pertama:", bacaData)

	//6. Operasi CRUD -> Update Data
	fmt.Println("Mengubah Data Ukur Suhu")
	// Konsepnya adalah mendapatkan data terlebih Dahulu (berdasarkan ID misalnya),
	// Lalu kita rubah Value/Nilai berdasarkan dari Field Atau Data yNg kita dapatkan
	var ubahDataCelcius UkurSuhu
	koneksiDb.First(&ubahDataCelcius, 1) // Mendapatkan data berdasarkan ID = 1
	ubahDataCelcius.SuhuCelcius = 40
	koneksiDb.Save(&ubahDataCelcius)
	fmt.Println("Data Ukur Suhu Setelah di Update:", ubahDataCelcius)

	//7. Operasi CRUD -> Delete Data
	fmt.Println("Menghapus Data Ukur Suhu")
	// Konsepnya adalaha mendapatkan data terlebih dahulu (berdasarkan ID misalnya),
	// Lalu kita hapus data tersebut
	var hapusDataId3 UkurSuhu
	koneksiDb.First(&hapusDataId3, 3) // Mendapatkan data berdasarkan ID = 3
	koneksiDb.Delete(&hapusDataId3)
	fmt.Println("Data Ukur Suhu ID:3 di Hapus:", hapusDataId3)
}
