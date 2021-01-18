package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Entity struct {
	Id        uuid.UUID  `gorm:"type:varchar(50);primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}

func (u *Entity) BeforeCreate(db *gorm.DB) (err error) {
	u.Id = uuid.New()
	return
}
