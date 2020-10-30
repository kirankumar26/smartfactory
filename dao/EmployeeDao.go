package dao

import (
	"database/sql"
	"fmt"

	"github.com/kirankumar26/smartfactory/config"
	"github.com/kirankumar26/smartfactory/model"
)

var con *sql.DB

func GetEmployees() model.Employees {
	con := config.CreateCon()
	//db.CreateCon()
	sqlStatement := "SELECT id, name, email FROM employee order by id"

	rows, err := con.Query(sqlStatement)
	fmt.Println(rows)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		//return c.JSON(http.StatusCreated, u);
	}
	defer rows.Close()
	result := model.Employees{}

	for rows.Next() {
		employee := model.Employee{}

		err2 := rows.Scan(&employee.ID, &employee.Name, &employee.Email)
		// Exit if we get an error
		if err2 != nil {
			fmt.Print(err2)
		}
		result.Employees = append(result.Employees, employee)
	}
	return result
}
