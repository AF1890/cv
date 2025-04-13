package repository

import "cv-api/domain"

type CVRepository interface {
	GetCV() (*domain.CV, error)
}
