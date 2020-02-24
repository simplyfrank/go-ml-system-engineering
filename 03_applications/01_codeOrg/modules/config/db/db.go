package db

import (
	"bytes"
	"database/sql"
	"fmt"
	"html/template"
	"log"
)



// Let's define the connection string syntax as Stringf strings
const (
	// Define a list of template constants to use to fill in the conn strings
	mySQLConString = "{{.user}}:{{.password}}@/{{.dbname}}"
	// TODO: Insert other common constring formats
)
// Lets define the database connections here


func NewSQL(dbType string, dbCon connector) (*sql.DB, error) {
	switch dbType {
	case "mysql":
		// Do some necessary initialization for MySQL
		fmt.Println("Got a mysql instance")
		return dbCon.connect()
	case "postgres", "postgresQL":
		// Do some necessary initialization for PostgreSQL
		//fmt.Println("Got a postgres instance")
		return dbCon.connect()

	// ----- NoSQL Variants ---------->
	case "mongo", "mongodb":
		// Do some necessary initialization for MongoDB
		log.Println("Requested a new mongoDB instance")
		return dbCon.connect()
	}
}

type connector interface {
	connect() (*sql.DB, error)
}


// Define the connector structs
type SQLConnector struct {
	Username string
	Password string
	DBname string
}
// Formats the information to pass as a connection string
func (s *SQLConnector) connect() (*sql.DB, error) {
	t, _ := template.New("").Parse(mySQLConString)
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, s); err != nil { return nil, err }

	// Connect to db
	db, err := sql.Open("mysql", tpl.String())
	if err != nil {
		return nil, err
	}
	return db, nil
}
// Just for illustration
func(s *SQLConnector) String() string {
	return fmt.Sprintf("This is a %s type instance connection", "mysql")
}

// Define the connection functions for each type of database
type MySQLConnector struct {
	SQLConnector
	// Additional fields only necessary for MySQL
}
type MongoDBConnector struct {
	SQLConnector
	// Additional fields ony necessary for MongoDB
}

