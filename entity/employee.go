package entity

type Employee struct {
	Entity
	FullName    string `gorm:"type:varchar(255);not null;index"`
	Phone       string `gorm:"type:varchar(255);not null"`
	Age         int8   `gorm:"type:int"`
	Status      string `gorm:"type:varchar(20);not null"`
	UserID      string
	User        User         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Departments []Department `gorm:"many2many:employee_departments"`
}
