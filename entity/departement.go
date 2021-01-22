package entity

type Department struct {
	Entity
	Name        string    `gorm:"type:varchar(255);unique;not null"`
	Description string    `gorm:"type:text"`
	Projects    []Project
}
