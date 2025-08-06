package main

import "fmt"

// Inisialisasi structure bernama FormPendaftaran
type FormPendaftaran struct {
	NamaLengkap string
	Email       string
	Gender      string
	Usia        int
}

// Method yang besifat pass by value
func (f FormPendaftaran) MencetakNamaLengkap() {
	fmt.Println("Nama Lengkap:", f.NamaLengkap)
}

// Method yang bersifat pass by reference
func (f *FormPendaftaran) MerubahUsia(usiaBaru int) {
	f.Usia = usiaBaru
	fmt.Println("Usia telah diubah menjadi:", f.Usia)
}

func MencetakNamaLengkapDenganFunction(form FormPendaftaran) {
	fmt.Println("Nama Lengkap:", form.NamaLengkap)
}

func main() {
	pendaftaran1 := FormPendaftaran{
		NamaLengkap: "Budi Santoso",
		Email:       "email@email.com",
		Gender:      "Laki-laki",
		Usia:        30,
	}

	// method
	pendaftaran1.MencetakNamaLengkap()
	// method pass by reference
	pendaftaran1.MerubahUsia(31)
}
