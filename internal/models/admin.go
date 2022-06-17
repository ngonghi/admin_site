package models

import "time"

type Admin struct {
	ID        int    `gorm:"type:int AUTO_INCREMENT"`
	Name      string `sql:"type:varchar(255);not null"`
	Email     string `sql:"type:varchar(255);not null"`
	Password  string `sql:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
