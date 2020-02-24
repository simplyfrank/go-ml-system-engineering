package config

import (
	"database/sql"
	"fmt"
	"html/template"
)
// Let's define the connection string syntax as Stringf strings
const (
	// Define a list of template constants to use to fill in the conn strings
	mySQLConString = "{{.user}}:{{.password}}@/{{.dbname}}"
	// TODO: Insert other common constring formats
)
// Lets define the database connections here


func NewSQL(dbType string) {
	switch dbType {
	case "mysql":
		fmt.Println("Got a mysql instance")
	case "postgres", "postgresQL":
		fmt.Println("Got a postgres instance")
	}
}

// Define the connector structs
type mySQLConnector struct {
	username string
	password string
	dbname string
}
// Formats the information to pass as a connection string
func (*mySQLConnector) String() string {
	t, _ := template.New("").Parse(mySQLConString)
	t.
}
// TODO: Think about generalizing the connector to work with all types of db systems


// Define the connection functions for each type of database
// system here separetely
func newMSQL(dbCon mySQLConnector) *sql.DB {
	db, err := sql.Open(
		"mysql",
		dbCon.String())
}