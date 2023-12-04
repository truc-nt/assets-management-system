package services

import (
	"server/internal/db"
	"server/internal/models"
)

type IAssetService interface {
	GetAssetsByDepartmentId(id uint32) ([]*models.Asset, error)
	GetAssetById(id uint32) (*models.Asset, error)
	CreateAsset(asset *models.Asset) error
	UpdateAsset(id uint32, asset *models.Asset) error
	DeleteAsset(id uint32) error
}

type AssetService struct {
}

func NewAssetService() *AssetService {
	return &AssetService{}
}

func (s *AssetService) GetAssetsByDepartmentId(id uint32) ([]*models.Asset, error) {
	return models.GetAssetsByDepartmentId(db.DB, id)
}

func (s *AssetService) GetAssetById(id uint32) (*models.Asset, error) {
	return models.GetAssetById(db.DB, id)
}

func (s *AssetService) CreateAsset(assets *models.Asset) error {
	return models.CreateAsset(db.DB, assets)
}

func (s *AssetService) UpdateAsset(id uint32, asset *models.Asset) error {
	_, err := models.GetAssetById(db.DB, id)
	if err != nil {
		return err
	}
	asset.Id = id
	return models.UpdateAsset(db.DB, asset)
}

func (s *AssetService) DeleteAsset(id uint32) error {
	return models.DeleteAsset(db.DB, id)
}
