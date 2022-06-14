package models

import "time"

type Admin struct {
	ID        int    `gorm:"type:int AUTO_INCREMENT"`
	Name      string `sql:"type:varchar(255)"`
	Email     string `sql:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
