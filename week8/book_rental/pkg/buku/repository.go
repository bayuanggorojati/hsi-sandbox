package buku

import (
	"book_rental_app/pkg/entities"

	"gorm.io/gorm"
)

type BukuRepository struct {
	db *gorm.DB
}

func NewBukuRepository(db *gorm.DB) *BukuRepository {
	return &BukuRepository{
		db: db,
	}
}
func (r *BukuRepository) Create(buku entities.Buku) error {
	return r.db.Create(&buku).Error
}

func (r *BukuRepository) GetAll() ([]entities.Buku, error) {
	var bukus []entities.Buku
	err := r.db.Find(&bukus).Error
	return bukus, err
}

// get buku by ID --> use db.First methodß
func (r *BukuRepository) GetByID(id int) (entities.Buku, error) {
	var buku entities.Buku
	err := r.db.First(&buku, id).Error
	return buku, err
}
