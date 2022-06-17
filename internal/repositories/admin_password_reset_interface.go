package repositories

import "github.com/ngonghi/admin_site/internal/models"

type AdminPasswordResetInterface interface {
	Create(m *models.AdminPasswordReset) error
	First(m *models.AdminPasswordReset, where ...interface{})
}
