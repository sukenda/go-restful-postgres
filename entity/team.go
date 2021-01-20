package entity

type Team struct {
	Entity
	Name       string `gorm:"type:varchar(255);not null"`
	EmployeeID string
	Employee   Employee `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProjectID  string
	Project    Project    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Employees  []Employee `gorm:"many2many:employee_works"`
}
