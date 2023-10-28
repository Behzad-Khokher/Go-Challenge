package usecase

import (
	"errors"
	"fmt"

	"github.com/Behzad-Khokher/Go-Challenge/domain"
	Branch "github.com/Behzad-Khokher/Go-Challenge/domain/branch"
)

type CreateBranchReq struct {
	Name string
}

type EditBranchReq struct {
	ID   int
	Name *string
}

type BranchUsecase interface {
	Create(req CreateBranchReq) error
	Edit(req EditBranchReq) error
	GetByID(id int) (*domain.Branch, error)
	GetAll() (domain.Branches, error)
}

type branchUsecase struct {
	store Branch.Store
}

func NewBranchUsecase(bs Branch.Store) BranchUsecase {
	return &branchUsecase{store: bs}
}

func (u *branchUsecase) Create(req CreateBranchReq) error {
	if req.Name == "" {
		return errors.New("branch name is required")
	}
	branch := &domain.Branch{
		Name: req.Name,
	}
	return u.store.Create(branch)
}

func (u *branchUsecase) Edit(req EditBranchReq) error {
	branch, err := u.store.GetByID(req.ID)
	if err != nil {
		return err
	}
	if branch == nil {
		return fmt.Errorf("branch with id %d not found", req.ID)
	}
	if req.Name != nil {
		branch.Name = *req.Name
	}
	return u.store.Update(branch)
}

func (u *branchUsecase) GetByID(id int) (*domain.Branch, error) {
	return u.store.GetByID(id)
}

func (u *branchUsecase) GetAll() (domain.Branches, error) {
	return u.store.GetAll()
}
