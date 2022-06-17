package repositories

import "github.com/ngonghi/admin_site/internal/models"

type Admin interface {
	First(m *models.Admin, where ...interface{}) error
	Find(m *[]models.Admin, where ...interface{}) error
	Create(m *models.Admin) error
}
