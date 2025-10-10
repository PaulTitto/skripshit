package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uint      `gorm:"primary_key;AUTO_INCREMENT"`
	UUID        uuid.UUID `gorm:"type:uuid;not null"`
	Name        string    `gorm:"type:varchar(100); not null"`
	Username    string    `gorm:"type:varchar(100); not null"`
	Password    string    `gorm:"type:varchar(255); not null"`
	PhoneNumber string    `gorm:"type:varchar(15); not null"`
	Email       string    `gorm:"type:varchar(100); not null"`
	RoleID      uint      `gorm:"type:uint; not null"`
	Streak      int       `gorm:"type:int;not null"`
	CreateAt    time.Time
	UpdateAt    time.Time
	Role        Role `gorm:"foreignKey:role_id; references:id; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
