package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/datatypes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserDate struct {
	ID          uint `gorm:"primarykey"`
	Name        string
	BirthOfDate datatypes.Date
	Address     string
}

func main() {
	// 1. Mendefinisikan Koneksi Database
	dsn := "coba:Password123!@tcp(172.16.202.130:3306)/hsisandbox?charset=utf8mb4&parseTime=True&loc=Local"
	// 2. Membuat Koneksi Database
	koneksiDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Koneksi Gagal, terjadi kesalahan pada saat berkomunikasi dengan DB", err.Error())
	}

	koneksiDb.AutoMigrate(&UserDate{})

	fmt.Println("Adding a new user")
	newUser := UserDate{
		Name:        "Bayu Anggorojati",
		BirthOfDate: datatypes.Date(time.Date(1984, 4, 10, 2, 0, 0, 0, time.Now().Local().Location())),
		Address:     "Tangerang Selatan",
	}

	insertUserDb := koneksiDb.Create(&newUser)
	if insertUserDb.Error != nil {
		log.Fatal("Gagal menambahkan user baru ke database")
	}
	fmt.Println("Berhasil menambahkan user baru dengan ID: ", newUser.ID)
}
