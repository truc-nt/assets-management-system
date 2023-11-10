package services

import (
	"server/internal/db"
	"server/internal/models"
)

var GetAssetsByDepartmentId = func(id uint32) ([]*models.Asset, error) {
	return models.GetAssetsByDepartmentId(db.DB, id)
}

var GetAssetById = func(id uint32) (*models.Asset, error) {
	return models.GetAssetById(db.DB, id)
}

var CreateAsset = func(assets *models.Asset) error {
	return models.CreateAsset(db.DB, assets)
}

var UpdateAsset = func(id uint32, asset *models.Asset) error {
	_, err := models.GetAssetById(db.DB, id)
	if err != nil {
		return err
	}
	asset.Id = id
	return models.UpdateAsset(db.DB, asset)
}

var DeleteAsset = func(id uint32) error {
	return models.DeleteAsset(db.DB, id)
}
