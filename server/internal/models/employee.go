package models

import "gorm.io/gorm"

type Employee struct {
	Id   uint32
	Name string
}

func GetEmployees(db *gorm.DB) ([]*Employee, error) {
	var employees []*User
	err := db.Model(&User{}).Where("role = ?", 1).Find(&employees).Error
	var employeesInfo []*Employee
	for _, employee := range employees {
		employeeInfo := Employee{
			Id:   employee.ID,
			Name: employee.DName,
		}
		employeesInfo = append(employeesInfo, &employeeInfo)
	}
	return employeesInfo, err
}
