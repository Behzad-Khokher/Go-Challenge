package pg

import (
	"errors"

	"github.com/Behzad-Khokher/Go-Challenge/domain"
	Component "github.com/Behzad-Khokher/Go-Challenge/domain/components"
	"gorm.io/gorm"
)

type ComponentStore struct {
	DB *gorm.DB
}

func NewComponentStore(db *gorm.DB) Component.Store {
	return &ComponentStore{DB: db}

}

func (store *ComponentStore) Create(c *domain.Component) error {
	return store.DB.Create(c).Error
}

func (store *ComponentStore) Update(c *domain.Component) error {
	return store.DB.Save(c).Error
}

func (store *ComponentStore) Delete(c *domain.Component) error {
	return store.DB.Delete(c).Error
}

func (store *ComponentStore) GetByID(id int) (*domain.Component, error) {
	var comp domain.Component
	if err := store.DB.First(&comp, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &comp, nil
}

func (store *ComponentStore) GetAll() (*domain.Components, error) {
	var comps domain.Components
	if err := store.DB.Find(&comps).Error; err != nil {
		return nil, err
	}
	return &comps, nil
}
