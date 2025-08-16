package main

import "assignment/pegawai"

func main() {
	// isi data dummy pegawai
	pegawaiBayu := pegawai.Pegawai{
		Nama:        "Bayu Anggorojati",
		Posisi:      "CEO",
		GajiBulanan: 100000000,
	}

	pegawaiBayu.TampilkanInformasi()
}
