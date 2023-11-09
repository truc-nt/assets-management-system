package models

import (
	"time"

	"gorm.io/gorm"
)

type Assets struct {
	ID         uint32    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`
	Type       string    `gorm:"column:type" json:"type"`
	Status     string    `gorm:"column:status" json:"status"`
	Department string    `gorm:"column:department" json:"department"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

func GetAssets(db *gorm.DB) ([]*Assets, error) {
	var assets []*Assets = make([]*Assets, 0)
	err := db.Model(&Assets{}).Find(&assets).Error

	if err != nil {
		return nil, err
	}
	return assets, nil
}
