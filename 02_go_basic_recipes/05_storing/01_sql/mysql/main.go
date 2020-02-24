package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"net/http"
)

// Program Setup (TODO: Move into config reader)
const (
	port int = 8080
	user string = "testuser"
	password string = "testpass"
	dbname string = "testdb"
)

// Global variables we need
var db *sql.DB  // This ensures all functions have access to the database handle

func init() {

}

func main() {
	var err error

	// Set up the database connection
	db, err = sql.Open("mysql", fmt.Sprintf(
		"%s:%s@/%s", user, password, dbname))
	check(err, "Open database", "get connection string") // Helper function to check errors i only want to trace
	defer db.Close()


	// Set up the handlers
	http.HandleFunc("/", baseHandler)

	// Start the server
	err = http.ListenAndServe(":8080", nil)
	check(err, "http.ListenAndServer", "start server to listen on assigned port")
}

// -------------------- QUERY FUNCTIONS ---------------------------->>>>
func friends(w http.ResponseWriter, req * http.Request) {
	rows, err := db.Query("SELECT name FROM friends")
	check(err, "friends", "query for list of all friends")

	// data to be used in query
	var s, name string
	s = "RETRIEVED RECORDS:\n"

	//query
	for rows.Next() {
		err = rows.Scan(&name)
		check(err, "rows.Next()", "retrieve name from database per row")
		s += name + "\n"
	}
}

func createTable(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare("CREATE TABLE customer (name VARCHAR(20));")
	check(err, "db.Prepare", "create table customet with value name")

	r, err := stmt.Exec()
	check(err, "stmt.Exec", "Execute prepared sql create statement")

	n, err := r.RowsAffected()
	check(err, "r.RowsAffected", "select the number of affected rows from table creation")

	fmt.Fprintf(w, "INSERTED %d records into the databse", n)
}

func read(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT * FROM customer;`)
	check(err, "db.Query", "read full customer table")

	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		check(err, "rows.Scan", "copy scanned value into name string variable")
		fmt.Println(name)

		fmt.Fprintln(w, "RETRIEVED RECORD: ", name)
	}
}

func insert(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`INSERT INTO customer VALUES("James");`)
	check(err, "db.Prepare", "insert single row into customer table")

	r, err = stmt.Exec()
	check(err, "stmt.Exec", "compile check insert statement")

	n, err := r.RowsAffected()
	check(err, "r.RowsAffected", "get number of affected rows by insert")

	fmt.Fprintln(w, "INSERTED RECORDS: ", n)
}


// -------------------- HELPER FUNCTIONS --------------------------->>>>>
func check(err error, fun string, task string) {
	if err != nil {
		log.Printf("%s:: Unable to %s", fun, task)
	}
}

// -------------------- HANDLER FUNCTIONS --------------------------->>>>>
func baseHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Welcome home")
}

