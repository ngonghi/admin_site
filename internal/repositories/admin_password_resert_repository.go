package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/ngonghi/admin_site/internal/models"
)

type AdminPasswordResetRepository struct {
	DB *gorm.DB
}

func (r *AdminPasswordResetRepository) First(m *models.AdminPasswordReset, where interface{}) error {
	return r.DB.First(m, where).Error
}

func (r *AdminPasswordResetRepository) Create(m *models.AdminPasswordReset) error {
	return r.DB.Create(m).Error
}
