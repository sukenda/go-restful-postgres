package entity

type Department struct {
	Entity
	Name        string `gorm:"type:varchar(255);not null;index"`
	Description string `gorm:"type:text"`
}
