package entity

type User struct {
	Entity
	Username   string `gorm:"type:varchar(100);unique"`
	Password   string `gorm:"not null"`
	Email      string `gorm:"type:varchar(50);unique;not null"`
	EmployeeID string
	Employee   Employee `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
