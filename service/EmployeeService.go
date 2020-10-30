package service

import (
	"github.com/kirankumar26/smartfactory/dao"
	"github.com/kirankumar26/smartfactory/model"
)

/* Fetch Employees*/
func GetEmployees() model.Employees {
	employees := dao.GetEmployees()
	return employees
}
