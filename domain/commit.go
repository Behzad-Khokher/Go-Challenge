package domain

import "time"

// import "time"

type Commits []Commit

// Commit represents a snapshot of changes.
type Commit struct {
	ID        int       `gorm:"column:id;primaryKey"`
	Name      string    `gorm:"column:name"`
	BranchID  int       `gorm:"column:branch_id"`
	Timestamp time.Time `gorm:"column:timestamp;default:CURRENT_TIMESTAMP"`

	// Components Components `gorm:"foreignKey:commit_id;references:id"`
}
