package entity

type Project struct {
	Entity
	Name         string `gorm:"type:varchar(255);not null"`
	Description  string `gorm:"type:text"`
	DepartmentID string
	Department   Department `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
