package pegawai

import "fmt"

type Pegawai struct {
	Nama        string
	Posisi      string
	GajiBulanan float64
}

func (p Pegawai) HitungGajiTahunan() float64 {
	return 12 * p.GajiBulanan
}

type InformasiPegawai interface {
	TampilkanInformasi()
}

func (p Pegawai) TampilkanInformasi() {
	fmt.Println("Nama: " + p.Nama)
	fmt.Println("Posisi: " + p.Posisi)
	fmt.Printf("Gaji tahunan: %.2f\n", p.HitungGajiTahunan())
}
