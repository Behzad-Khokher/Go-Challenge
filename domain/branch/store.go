package Branch

import (
	"github.com/Behzad-Khokher/Go-Challenge/domain"
)

type Store interface {
	Create(branch *domain.Branch) error
	GetByID(id int) (*domain.Branch, error)
	GetByName(name string) (*domain.Branch, error)
	GetAll() (domain.Branches, error)
	Update(branch *domain.Branch) error
	Delete(id int) error
}
