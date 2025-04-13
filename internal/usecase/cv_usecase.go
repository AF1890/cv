package usecase

import (
	"cv-api/domain"
	"cv-api/internal/repository"
)

type CVUsecase interface {
	GetCV() (*domain.CV, error)
}

type cvUsecase struct {
	cvRepo repository.CVRepository
}

func NewCVUsecase(cvRepo repository.CVRepository) CVUsecase {
	return &cvUsecase{
		cvRepo: cvRepo,
	}
}

func (u *cvUsecase) GetCV() (*domain.CV, error) {
	return u.cvRepo.GetCV()
}
