package entity

type User struct {
	Entity
	Username string `gorm:"type:varchar(100);unique"`
	Password string `gorm:"not null"`
	Email    string `gorm:"type:varchar(50);not null;unique"`
	Phone    string `gorm:"type:varchar(15)"`
	Profile  string `gorm:"type:varchar(255)"`
}
