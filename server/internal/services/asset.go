package services

import (
	"server/internal/db"
	"server/internal/models"
)

type IAssetService interface {
	GetAssets(param *models.GetAssetsParam) ([]*models.Asset, error)
	GetAssetById(id uint32) (*models.Asset, error)
	CreateAsset(asset *models.Asset) error
	UpdateAsset(id uint32, asset *models.Asset) error
	DeleteAsset(id uint32) error
}

type AssetService struct {
	Repository models.IAssetRepository
}

func NewAssetService() IAssetService {
	return &AssetService{
		Repository: models.NewAssetRepository(db.DB),
	}
}

func (s *AssetService) GetAssets(param *models.GetAssetsParam) ([]*models.Asset, error) {
	return s.Repository.GetAssets(param)
}

func (s *AssetService) GetAssetById(id uint32) (*models.Asset, error) {
	return s.Repository.GetAssetById(id)
}

func (s *AssetService) CreateAsset(assets *models.Asset) error {
	return s.Repository.CreateAsset(assets)
}

func (s *AssetService) UpdateAsset(id uint32, asset *models.Asset) error {
	_, err := s.Repository.GetAssetById(id)
	if err != nil {
		return err
	}
	asset.Id = id
	return s.Repository.UpdateAsset(asset)
}

func (s *AssetService) DeleteAsset(id uint32) error {
	return s.Repository.DeleteAsset(id)
}
