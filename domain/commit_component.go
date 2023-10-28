package domain

type CommitComponent struct {
	CommitID    int `gorm:"column:commit_id"`
	ComponentID int `gorm:"column:component_id"`
}
