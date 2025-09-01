package pegawai

import "fmt"

type Employee struct {
	ID          uint `gorm:"primaryKey"`
	Nama        string
	Posisi      string
	GajiBulanan float64
}

func (e Employee) HitungGajiTahunan() float64 {
	return 12 * e.GajiBulanan
}

type InformasiPegawai interface {
	TampilkanInformasi()
}

func (e Employee) TampilkanInformasi() {
	fmt.Println(e.ID)
	fmt.Println(e.Nama)
	fmt.Println(e.Posisi)
	fmt.Printf("Gaji tahunan: %.2f\n", e.HitungGajiTahunan())
}
