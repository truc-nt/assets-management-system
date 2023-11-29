package services

import (
	"server/internal/db"
	"server/internal/models"
)

var GetEmployees = func() ([]*models.Employee, error) {
	return models.GetEmployees(db.DB)
}
