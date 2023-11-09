package services

import (
	"server/internal/db"
	"server/internal/models"
)

var GetAssets = func() ([]*models.Assets, error) {
	if _, err := models.GetAssets(db.DB); err != nil {
		return nil, err
	}

	return models.GetAssets(db.DB)
}
