package usecase

import (
	"errors"
	"fmt"

	"github.com/Behzad-Khokher/Go-Challenge/domain"
	CommitComponentDomain "github.com/Behzad-Khokher/Go-Challenge/domain/commit_components"
)

type AddComponentsToCommitReq struct {
	CommitID     int
	ComponentIDs []int
}

type CommitComponentUsecase interface {
	AddComponentsToCommit(req AddComponentsToCommitReq) error
	RemoveComponentFromCommit(commitID, componentID int) error
	GetComponentsByCommitID(commitID int) ([]domain.CommitComponent, error)
	GetCommitsByComponentID(componentID int) ([]domain.CommitComponent, error)
}

type commitComponentUsecase struct {
	store CommitComponentDomain.Store
}

func NewCommitComponentUsecase(cs CommitComponentDomain.Store) CommitComponentUsecase {
	return &commitComponentUsecase{store: cs}
}

func (u *commitComponentUsecase) AddComponentsToCommit(req AddComponentsToCommitReq) error {
	if req.CommitID == 0 {
		return errors.New("commit id is required")
	}
	if len(req.ComponentIDs) == 0 {
		return errors.New("at least one component id is required")
	}

	for _, componentID := range req.ComponentIDs {
		cc := &domain.CommitComponent{
			CommitID:    req.CommitID,
			ComponentID: componentID,
		}
		if err := u.store.Create(cc); err != nil {
			return fmt.Errorf("failed to add component with id %d to commit: %v", componentID, err)
		}
	}
	return nil
}

func (u *commitComponentUsecase) RemoveComponentFromCommit(commitID, componentID int) error {
	return u.store.DeleteByCommitIDAndComponentID(commitID, componentID)
}

func (u *commitComponentUsecase) GetComponentsByCommitID(commitID int) ([]domain.CommitComponent, error) {
	return u.store.GetByCommitID(commitID)
}

func (u *commitComponentUsecase) GetCommitsByComponentID(componentID int) ([]domain.CommitComponent, error) {
	return u.store.GetByComponentID(componentID)
}
