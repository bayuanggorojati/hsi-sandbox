package main

import "fmt"

// Mendefiniskan Interface/Kontrak/Method Yang Harus Dimiliki oleh Struct
// Konsep Interface itu adalah Polimorphism
type FormPendaftaranInterface interface {
	ValidasiUsia(usia int) bool
	ValidasiGender(gender string) bool
}

type FormPendaftaran struct {
	NamaLengkap string
	Email       string
	Gender      string
	Usia        int
}

type FormPendaftaranUsiaSenja struct {
	NamaLengkap    string
	Email          string
	Gender         string
	Usia           int
	PenyakitBawaan string
}

//Implementasi Untuk Struct Reguler
func (f FormPendaftaran) ValidasiUsia(usia int) bool {
	if usia < 15 || usia > 75 {
		return false
	}
	return true
}

func (f FormPendaftaran) ValidasiGender(gender string) bool {
	if gender != "Laki-laki" && gender != "Perempuan" {
		return false
	}
	return true
}

// Implementasi Untuk Struct Usia Senja
func (f FormPendaftaranUsiaSenja) ValidasiUsia(usia int) bool {
	if usia > 75 {
		return true
	}
	return false
}

func (f FormPendaftaranUsiaSenja) ValidasiGender(gender string) bool {
	if gender != "Laki-laki" && gender != "Perempuan" {
		return false
	}
	return true
}

func ValidasiUsiaForm(formInt FormPendaftaranInterface, usia int) bool {
	return formInt.ValidasiUsia(usia)
}

func main() {
	pendaftaran1 := FormPendaftaran{
		NamaLengkap: "Budi Santoso",
		Email:       "email@email.com",
		Usia:        35,
	}

	pendaftaran2 := FormPendaftaranUsiaSenja{
		NamaLengkap:    "Abu Fulan",
		Email:          "Fulan@email.com",
		Usia:           80,
		PenyakitBawaan: "Diabetes",
	}

	apakahUsiaValid1 := ValidasiUsiaForm(pendaftaran1, pendaftaran1.Usia)
	fmt.Println("Apakah usia valid?", apakahUsiaValid1)

	apakahUsiaValid2 := ValidasiUsiaForm(pendaftaran2, pendaftaran2.Usia)
	fmt.Println("Apakah usia Senja valid?", apakahUsiaValid2)
}
