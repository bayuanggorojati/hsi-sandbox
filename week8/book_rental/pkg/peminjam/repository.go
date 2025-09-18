package peminjam

import (
	"book_rental_app/pkg/entities"

	"gorm.io/gorm"
)

type PeminjamRepository struct {
	db *gorm.DB
}

func NewPeminjamRepository(db *gorm.DB) *PeminjamRepository {
	return &PeminjamRepository{
		db: db,
	}
}

func (r *PeminjamRepository) Create(peminjam entities.Peminjam) error {
	return r.db.Create(&peminjam).Error
}

func (r *PeminjamRepository) GetAll() ([]entities.Peminjam, error) {
	var peminjams []entities.Peminjam
	err := r.db.Model(&entities.Peminjam{}).Preload("BukuPinjaman").Find(&peminjams).Error
	return peminjams, err
}

func (r *PeminjamRepository) GetByID(id int) (entities.Peminjam, error) {
	var peminjam entities.Peminjam
	err := r.db.Model(&entities.Peminjam{}).Preload("BukuPinjaman").First(&peminjam, id).Error
	return peminjam, err
}
