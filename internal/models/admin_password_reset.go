package models

import "time"

type AdminPasswordReset struct {
	ID        int    `gorm:"type:int AUTO_INCREMENT"`
	Email     string `sql:"type:varchar(255);not null"`
	CreatedAt time.Time
}
