package main

// To download dependencies and include them in go.mod:
//   go get github.com/go-sql-driver/mysql
// Add the package path to the import section

// What changed to the files go.mod and go.sum?
// go.mod: Package name and version number
// go.sum: Package name, version number, and file hash
// (You should commit both files to Git)

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

func main() {
	config := mysql.NewConfig()
	config.User = "root"
	config.Passwd = ""
	config.DBName = "db"

	// sql.Open() returns (*DB, error)
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil { // If err is not nil, an error occurred
		panic(err)
	}
	defer db.Close() // Deferred calls are run at end of function; useful for cleanup functions
}

/*
https://gobyexample.com/errors
https://gobyexample.com/defer
*/
