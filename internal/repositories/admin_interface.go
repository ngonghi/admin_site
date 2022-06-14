package repositories

import "github.com/ngonghi/admin_site/internal/models"

type Admin interface {
	First(m *models.Admin) error
	Find(m *[]models.Admin) error
	Create(m *models.Admin) error
	Ping() error
}
