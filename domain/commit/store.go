package commit

import "github.com/Behzad-Khokher/Go-Challenge/domain"

type Store interface {
	Create(c *domain.Commit) error
	Update(c *domain.Commit) error
	Delete(c *domain.Commit) error
	GetByID(id int) (*domain.Commit, error)
	GetByName(name string) (*domain.Commit, error)
	GetByBranchID(branchID int) (domain.Commits, error)
	GetAll() (domain.Commits, error)
}
