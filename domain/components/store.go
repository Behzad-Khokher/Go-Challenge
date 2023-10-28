package Component

import "github.com/Behzad-Khokher/Go-Challenge/domain"

type Store interface {
	Create(c *domain.Component) error
	Update(c *domain.Component) error
	Delete(c *domain.Component) error
	GetByID(id int) (*domain.Component, error)
	GetAll() (*domain.Components, error)
}
