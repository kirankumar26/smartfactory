package config

import (
	"database/sql"
	"fmt"

	"github.com/kirankumar26/smartfactory/utility"

	_ "github.com/go-sql-driver/mysql"
)

/*Create mysql connection*/
func CreateCon() *sql.DB {
	utility.InitializeConfig()
	dbDriver := utility.Configuration.DatabaseDriver
	dbUser := utility.Configuration.DatabaseUserName
	dbPass := utility.Configuration.DatabasePassword
	dbName := utility.Configuration.DatabaseName
	dbHostName := utility.Configuration.DatabaseHostName
	dbPort := utility.Configuration.DatabasePort
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHostName+":"+dbPort+")/"+dbName)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}
	//defer db.Close()
	// make sure connection is available
	err = db.Ping()
	fmt.Println(err)
	if err != nil {
		fmt.Println("db is not connected")
		fmt.Println(err.Error())
	}
	return db
}

/*end mysql connection*/
