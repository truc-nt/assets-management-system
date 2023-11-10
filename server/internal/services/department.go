package services

import (
	"server/internal/db"
	"server/internal/models"
)

var CreateDepartment = func(department *models.Department) error {
	return models.CreateDepartment(db.DB, department)
}
