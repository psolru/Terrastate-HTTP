package sqlite3

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/psolru/terrastate-http/env"
	"log"
)

// Open creates and returns the filehandle to the sqlite3 database
func Open() *sqlx.DB {
	db, err := sqlx.Open("sqlite3", env.WorkDir()+"/sqlite3/states.sqlite")
	if err != nil {
		log.Fatalf("[SQLITE3] %v", err)
	}

	return db
}

// createTable creates the sqlite3 table if it does not exists
func createTable() {
	log.Println("[SQLITE3] Create table if not exists...")

	// make it quick and dirty at startup for now
	_, err := Exec(`
		CREATE TABLE IF NOT EXISTS tf_state (
			ident VARCHAR(255) PRIMARY KEY,
			data TEXT,
			serial INT,
			version INT,
			lock INT DEFAULT 1,
			lock_id VARCHAR(255)
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
}

// Init makes some startup tasks
func Init() {
	log.Println("[SQLITE3] Init...")
	createTable()
}

// Exec handles a sql.Prepare and sql.Exec in one function for sake of simplicity
// It returns the exec sql.Result and error
func Exec(query string, values ...interface{}) (sql.Result, error) {
	db := Open()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	return stmt.Exec(values...)
}

// QueryRow handles a sql.QueryRow
// It returns the gotten *sql.Rows and error
func QueryRow(query string, values ...interface{}) *sql.Row {
	db := Open()
	return db.QueryRow(query, values...)
}

// Query handles a sql.QueryRow
// It returns the gotten *sql.Rows and error
func Query(query string, values ...interface{}) (*sql.Rows, error) {
	db := Open()
	return db.Query(query, values...)
}
