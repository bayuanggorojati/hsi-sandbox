package peminjam

import (
	"book_rental_app/pkg/entities"
)

type PeminjamRepo interface {
	Create(peminjam entities.Peminjam) error
	GetAll() ([]entities.Peminjam, error)
	GetByID(id int) (entities.Peminjam, error)
}

type PeminjamService struct {
	repo PeminjamRepo
}

func NewPeminjamService(repo PeminjamRepo) *PeminjamService {
	return &PeminjamService{repo: repo}
}

func (s *PeminjamService) CreatePeminjam(peminjam entities.Peminjam) error {
	return s.repo.Create(peminjam)
}

func (s *PeminjamService) GetAllPeminjam() ([]entities.Peminjam, error) {
	return s.repo.GetAll()
}

func (s *PeminjamService) GetPeminjamByID(id int) (entities.Peminjam, error) {
	return s.repo.GetByID(id)
}
