package pg

import (
	"github.com/Behzad-Khokher/Go-Challenge/domain"
	Branch "github.com/Behzad-Khokher/Go-Challenge/domain/branch"
	"gorm.io/gorm"
)

type BranchStore struct {
	DB *gorm.DB
}

func NewBranchStore(db *gorm.DB) Branch.Store {
	return &BranchStore{DB: db}
}

func (bs *BranchStore) Create(branch *domain.Branch) error {
	return bs.DB.Create(branch).Error
}

func (bs *BranchStore) GetByID(id int) (*domain.Branch, error) {
	var branch domain.Branch
	if err := bs.DB.First(&branch, id).Error; err != nil {
		return nil, err
	}
	return &branch, nil
}

func (bs *BranchStore) GetByName(name string) (*domain.Branch, error) {
	var branch domain.Branch
	if err := bs.DB.Where("name = ?", name).First(&branch).Error; err != nil {
		return nil, err
	}
	return &branch, nil
}

func (bs *BranchStore) GetAll() (domain.Branches, error) {
	var branches domain.Branches
	if err := bs.DB.Find(&branches).Error; err != nil {
		return nil, err
	}
	return branches, nil
}

func (bs *BranchStore) Update(branch *domain.Branch) error {
	return bs.DB.Save(branch).Error
}

func (bs *BranchStore) Delete(id int) error {
	return bs.DB.Delete(&domain.Branch{}, id).Error
}
