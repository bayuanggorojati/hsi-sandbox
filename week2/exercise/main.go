package main

import "fmt"

type PesertaSandbox struct {
	NamaLengkap  string
	Gender       string
	Usia         int
	SudahMenikah bool
	Domisili     string
}

// diteruskan sebaga pass By Value
// Struct dibuat baru atau dibuat salinan (copy)
func bertambahUsia(peserta PesertaSandbox) {
	peserta.Usia = peserta.Usia + 1
	fmt.Println("Usia peserta setelah melewati function bertambahUsia:", peserta.Usia)
}

// diteruskan sebagai pass by reference
// struct tidak dibuat baru, hanya mengirimkan alamat dari struct
// sehingga perubahan pada struct akan berpengaruh pada struct aslinya
func bertambahUsiaPointer(peserta *PesertaSandbox) {
	peserta.Usia = peserta.Usia + 1
	fmt.Println("Usia peserta setelah melewati function bertambahUsiaPointer:", peserta.Usia)
}

func main() {
	// Inisialisasi struct dengan zero value
	var pesertaDarizki PesertaSandbox // Nilai awalnya akan diisi dengan zero value
	pesertaDarizki.NamaLengkap = "Darizki Pratama"
	pesertaDarizki.Gender = "Ikhwan"
	pesertaDarizki.Usia = 40
	pesertaDarizki.SudahMenikah = true
	pesertaDarizki.Domisili = "Jakarta"

	// Inisialiasi struct dengan literal
	// Default Value dari Sebuah Struct
	pesertaStructVariable2 := PesertaSandbox{
		NamaLengkap:  "Siti Aisyah",
		Gender:       "Akhwat",
		Usia:         30,
		SudahMenikah: true,
		Domisili:     "Bandung",
	}
	// Cara akses struct
	fmt.Println("Nama Lengkap Peserta dari struct pesertaDarizki:", pesertaDarizki.NamaLengkap)
	fmt.Println(pesertaDarizki.NamaLengkap, " berjenis kelamin: ", pesertaDarizki.Gender, ", berusia: ", pesertaDarizki.Usia)
	fmt.Println("Nama Lengkap Peserta dari struct pesertaStructVariable 2:", pesertaStructVariable2.NamaLengkap)

	//Pass By Value, dimana Struct akan di copy paste
	bertambahUsia(pesertaDarizki)
	bertambahUsiaPointer(&pesertaDarizki) // Mengirimkan alamat dari struct pesertaDarizki
	//Struct aslinya
	fmt.Println(pesertaDarizki.Usia) // 40

	var alamatAsli string = "Duren Sawit Jakata Timur"
	var pointerAlamat *string = &alamatAsli // Pointer ke alamatAsli
	fmt.Println("Nilai dari variable alamatAsli:", alamatAsli)
	fmt.Println("Nilai dari variable Pointer ke alamat asli:", pointerAlamat) // outputnya adalah alamat dari variabel alamatAsli
	fmt.Println("Isi dari variable pointerAlamat:", *pointerAlamat)           // outputnya adalah isi dari alamatAsli

}
