package sql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() string {
	// Create the database handle, confirm driver is present
	db, _ := sql.Open("mysql", "root:@/user")
	defer db.Close()

	// Connect and check the server version
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	return version
}
