package models

import "gorm.io/gorm"

type Department struct {
	Id   uint32 `gorm:"column:id;primaryKey;autoIncrement"`
	Name string `gorm:"column:name" json:"name"`
}

func CreateDepartment(db *gorm.DB, department *Department) error {
	return db.Create(department).Error
}
