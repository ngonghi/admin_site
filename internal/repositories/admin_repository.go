package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/ngonghi/admin_site/internal/models"
)

type AdminRepository struct {
	DB *gorm.DB
}

func (s *AdminRepository) First(m *models.Admin) error {
	return s.DB.First(m).Error
}

func (s *AdminRepository) Create(m *models.Admin) error {
	return s.DB.Create(m).Error
}

func (s *AdminRepository) Find(m *[]models.Admin) error {
	return s.DB.Find(m).Error
}

func (s *AdminRepository) Ping() error {
	return s.DB.DB().Ping()
}
