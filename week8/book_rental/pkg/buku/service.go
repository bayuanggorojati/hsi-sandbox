package buku

import (
	"book_rental_app/pkg/entities"
)

type BukuRepo interface {
	Create(buku entities.Buku) error
	GetAll() ([]entities.Buku, error)
	GetByID(id int) (entities.Buku, error)
}

type BukuService struct {
	repo BukuRepo
}

func NewBukuService(repo BukuRepo) *BukuService {
	return &BukuService{repo: repo}
}

func (s *BukuService) CreateBuku(buku entities.Buku) error {
	return s.repo.Create(buku)
}

func (s *BukuService) GetAllBuku() ([]entities.Buku, error) {
	return s.repo.GetAll()
}

func (s *BukuService) GetBukuByID(id int) (entities.Buku, error) {
	return s.repo.GetByID(id)
}
