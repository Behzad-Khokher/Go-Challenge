package usecase

import (
	"errors"

	"github.com/Behzad-Khokher/Go-Challenge/domain"
	CommitDomain "github.com/Behzad-Khokher/Go-Challenge/domain/commit"
)

type CreateCommitReq struct {
	Name     string
	BranchID int
}

type CommitUsecase interface {
	Create(req CreateCommitReq) error
	GetByID(id int) (*domain.Commit, error)
	GetByName(name string) (*domain.Commit, error)
	GetByBranchID(branchID int) (domain.Commits, error)
	GetAll() (domain.Commits, error)
	Delete(id int) error
}

type commitUsecase struct {
	store CommitDomain.Store
}

func NewCommitUsecase(cs CommitDomain.Store) CommitUsecase {
	return &commitUsecase{store: cs}
}

func (u *commitUsecase) Create(req CreateCommitReq) error {
	if req.Name == "" {
		return errors.New("commit name is required")
	}
	commit := &domain.Commit{
		Name:     req.Name,
		BranchID: req.BranchID,
	}
	return u.store.Create(commit)
}

func (u *commitUsecase) GetByID(id int) (*domain.Commit, error) {
	return u.store.GetByID(id)
}

func (u *commitUsecase) GetByName(name string) (*domain.Commit, error) {
	return u.store.GetByName(name)
}

func (u *commitUsecase) GetByBranchID(branchID int) (domain.Commits, error) {
	return u.store.GetByBranchID(branchID)
}

func (u *commitUsecase) GetAll() (domain.Commits, error) {
	return u.store.GetAll()
}

func (u *commitUsecase) Delete(id int) error {
	commit := &domain.Commit{
		ID: id,
	}
	return u.store.Delete(commit)
}
