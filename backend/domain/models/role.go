package models

import "time"

type Role struct {
	ID        uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Code      string `gorm:"varchar(15); not null"`
	Name      string `gorm:"varchar(15); not null"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
