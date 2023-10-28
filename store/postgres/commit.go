package pg

import (
	"github.com/Behzad-Khokher/Go-Challenge/domain"
	"github.com/Behzad-Khokher/Go-Challenge/domain/commit"
	"gorm.io/gorm"
)

type CommitStore struct {
	DB *gorm.DB
}

func NewCommitStore(db *gorm.DB) commit.Store {
	return &CommitStore{DB: db}
}

func (s *CommitStore) Create(c *domain.Commit) error {
	return s.DB.Create(c).Error
}

func (s *CommitStore) Update(c *domain.Commit) error {
	return s.DB.Save(c).Error
}

func (s *CommitStore) Delete(c *domain.Commit) error {
	return s.DB.Delete(c).Error
}

func (s *CommitStore) GetByID(id int) (*domain.Commit, error) {
	var c domain.Commit
	err := s.DB.First(&c, id).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (s *CommitStore) GetByName(name string) (*domain.Commit, error) {
	var c domain.Commit
	err := s.DB.Where("name = ?", name).First(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (s *CommitStore) GetByBranchID(branchID int) (domain.Commits, error) {
	var commits domain.Commits
	err := s.DB.Where("branch_id = ?", branchID).Find(&commits).Error
	if err != nil {
		return nil, err
	}
	return commits, nil
}

func (s *CommitStore) GetAll() (domain.Commits, error) {
	var commits domain.Commits
	err := s.DB.Find(&commits).Error
	if err != nil {
		return nil, err
	}
	return commits, nil
}
