package domain

type Branches []Branch
type Branch struct {
	ID   int    `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name;unique"`
}
