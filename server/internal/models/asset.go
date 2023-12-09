package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type GetAssetsParam struct {
	UserId uint32 `form:"user_id"`
}

type Asset struct {
	Id          uint32    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	Type        string    `gorm:"column:type" json:"type"`
	Status      string    `gorm:"column:status" json:"status"`
	StatusNote  string    `gorm:"column:status_note" json:"status_note"`
	Description string    `gorm:"column:description" json:"description"`
	UserId      uint32    `gorm:"column:user_id;foreignKey:UserId;" json:"user_id"`
	user        User      `gorm:"foreignKey:UserId;"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type IAssetRepository interface {
	GetAssets(param *GetAssetsParam) ([]*Asset, error)
	GetAssetById(id uint32) (*Asset, error)
	CreateAsset(asset *Asset) error
	UpdateAsset(asset *Asset) error
	DeleteAsset(id uint32) error
}

type AssetRepository struct {
	DB *gorm.DB
}

func NewAssetRepository(db *gorm.DB) IAssetRepository {
	return &AssetRepository{DB: db}
}

func (r *AssetRepository) GetAssets(param *GetAssetsParam) ([]*Asset, error) {
	var assets []*Asset

	fmt.Println(param.UserId)
	if param.UserId == 0 {
		err := r.DB.Find(&assets).Error
		return assets, err
	}
	err := r.DB.Where(fmt.Sprintf("user_id = %d", param.UserId)).Find(&assets).Error
	return assets, err
}

func (r *AssetRepository) GetAssetById(id uint32) (*Asset, error) {
	var asset Asset
	err := r.DB.First(&asset, id).Error
	return &asset, err
}

func (r *AssetRepository) CreateAsset(asset *Asset) error {
	return r.DB.Create(&asset).Error
}

func (r *AssetRepository) UpdateAsset(asset *Asset) error {
	return r.DB.Updates(&asset).Error
}

func (r *AssetRepository) DeleteAsset(id uint32) error {
	tx := r.DB.Delete(&Asset{}, id)
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil

}
