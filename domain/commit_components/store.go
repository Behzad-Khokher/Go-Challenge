package commitcomponents

import "github.com/Behzad-Khokher/Go-Challenge/domain"

type Store interface {
	Create(cc *domain.CommitComponent) error
	DeleteByCommitIDAndComponentID(commitID, componentID int) error
	GetByCommitID(commitID int) ([]domain.CommitComponent, error)
	GetByComponentID(componentID int) ([]domain.CommitComponent, error)
}
