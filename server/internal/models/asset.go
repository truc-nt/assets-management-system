package models

import (
	"time"

	"gorm.io/gorm"
)

type Asset struct {
	Id          uint32    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	Type        string    `gorm:"column:type" json:"type"`
	Status      string    `gorm:"column:status" json:"status"`
	StatusNote  string    `gorm:"column:status_note" json:"status_note"`
	Description string    `gorm:"column:description" json:"description"`
	EmployeeId  uint32    `gorm:"column:employee_id" json:"employee_id"`
	employee    Employee  `gorm:"foreignKey:EmployeeId;"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func GetAssetsByEmployeeId(db *gorm.DB, id uint32) ([]*Asset, error) {
	var assets []*Asset
	err := db.Where("employee_id = ?", id).Find(&assets).Error
	return assets, err
}

func GetAssetById(db *gorm.DB, id uint32) (*Asset, error) {
	var asset Asset
	err := db.First(&asset, id).Error
	return &asset, err
}

func CreateAsset(db *gorm.DB, asset *Asset) error {
	return db.Create(&asset).Error
}

func UpdateAsset(db *gorm.DB, asset *Asset) error {
	return db.Updates(&asset).Error
}

func DeleteAsset(db *gorm.DB, id uint32) error {
	tx := db.Delete(&Asset{}, id)
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil

}
