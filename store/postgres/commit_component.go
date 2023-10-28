package pg

import (
	"github.com/Behzad-Khokher/Go-Challenge/domain"
	commitcomponents "github.com/Behzad-Khokher/Go-Challenge/domain/commit_components"
	"gorm.io/gorm"
)

type CommitComponentStore struct {
	DB *gorm.DB
}

func NewCommitComponentStore(db *gorm.DB) commitcomponents.Store {
	return &CommitComponentStore{DB: db}
}

func (s *CommitComponentStore) Create(cc *domain.CommitComponent) error {
	return s.DB.Create(cc).Error
}

func (s *CommitComponentStore) DeleteByCommitIDAndComponentID(commitID, componentID int) error {
	return s.DB.Where("commit_id = ? AND component_id = ?", commitID, componentID).Delete(&domain.CommitComponent{}).Error
}

func (s *CommitComponentStore) GetByCommitID(commitID int) ([]domain.CommitComponent, error) {
	var ccs []domain.CommitComponent
	err := s.DB.Where("commit_id = ?", commitID).Find(&ccs).Error
	if err != nil {
		return nil, err
	}
	return ccs, nil
}

func (s *CommitComponentStore) GetByComponentID(componentID int) ([]domain.CommitComponent, error) {
	var ccs []domain.CommitComponent
	err := s.DB.Where("component_id = ?", componentID).Find(&ccs).Error
	if err != nil {
		return nil, err
	}
	return ccs, nil
}
