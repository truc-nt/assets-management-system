package services

import (
	"server/internal/db"
	"server/internal/models"
)

var GetDepartments = func() ([]*models.Department, error) {
	return models.GetDepartments(db.DB)
}
