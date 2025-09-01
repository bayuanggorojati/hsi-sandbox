package main

import (
	"assignment/pegawai"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	// Import Aplikasi Standar

	"log"
)

func main() {
	// creating arrays of employeenames, positions and salaries
	employeeNames := [5]string{"Prabowo Subianto", "Gibran Rakabuming Raka", "Puan Maharani", "Ahmad Muzani", "Sufmi Dasco Ahmad"}
	positions := [5]string{"Presiden", "Wakil Presiden", "Ketua DPR", "Ketua MPR", "Wakil Ketua DPR"}
	salaries := [5]float64{250000000, 200000000, 215000000, 215000000, 200000000}

	// initiate an array of employees
	var daftarPegawai [5]pegawai.Employee

	// populating the values into list of employees using for loop
	for i := 0; i < len(employeeNames); i++ {
		daftarPegawai[i].Nama = employeeNames[i]
		daftarPegawai[i].Posisi = positions[i]
		daftarPegawai[i].GajiBulanan = salaries[i]
	}

	// definisikan koneksi database
	dsn := "root:password123@tcp(172.16.202.130:3306)/hsisandbox?charset=utf8mb4&parseTime=True&loc=Local"
	// 2. Membuat Koneksi Database
	koneksiDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Koneksi Gagal, terjadi kesalahan pada saat berkomunikasi dengan DB", err.Error())
	}
	//3. Migrasi Struktur Table dari Struct
	koneksiDb.AutoMigrate(&pegawai.Employee{})

	for i := 0; i < len(positions); i++ {
		// CRUD -> create / insert new data into database
		createNewData := koneksiDb.Create(&daftarPegawai[i])

		if createNewData.Error != nil {
			log.Fatal("Gagal menambahkan data baru", createNewData.Error.Error())
		}
		fmt.Println("Berhasil menambahkan data baru dengan ID: ", daftarPegawai[i].ID)

		// show annual salary of the newly added employee
		fmt.Printf("Gaji tahunan pegawai yang baru ditambahkan: %.f\n", daftarPegawai[i].HitungGajiTahunan())
		// retrieve and show a newly added employee from the database
		var dataPegawai pegawai.Employee
		koneksiDb.First(&dataPegawai, daftarPegawai[i].ID)

		fmt.Println("Data pegawai yang baru saja ditambahkan:", dataPegawai)
	}

	updateSalaries := [5]float64{50000000, 40000000, 41000000, 41000000, 40000000}
	// CRUD -> update salary of all employees
	for i := 0; i < len(salaries); i++ {
		var updateSalary pegawai.Employee
		koneksiDb.First(&updateSalary, daftarPegawai[i])
		updateSalary.GajiBulanan = updateSalaries[i]
		koneksiDb.Save(updateSalary)
		fmt.Println("Data dengan gaji bulanan pegawai setelah diupdate", updateSalary)
		fmt.Printf("Gaji tahunan setelah diupdate %.f\n", updateSalary.HitungGajiTahunan())
	}

	// delete one of the employees from database
	var hapusEmployee2 pegawai.Employee
	koneksiDb.First(&hapusEmployee2, 2)
	koneksiDb.Delete(&hapusEmployee2)
	fmt.Println("Data pegawai ID:2 dihapus", hapusEmployee2)
	fmt.Printf("Gaji tahunan pegawai yang baru dihapus %.f\n", hapusEmployee2.HitungGajiTahunan())
}
