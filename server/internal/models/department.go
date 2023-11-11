package models

import "gorm.io/gorm"

type Department struct {
	Id   uint32
	Name string
}

func GetDepartments(db *gorm.DB) ([]*Department, error) {
	var departments []*User
	err := db.Model(&User{}).Where("role = ?", 1).Find(&departments).Error
	var departmentsInfo []*Department
	for _, department := range departments {
		departmentInfo := Department{
			Id:   department.ID,
			Name: department.Username,
		}
		departmentsInfo = append(departmentsInfo, &departmentInfo)
	}
	return departmentsInfo, err
}
