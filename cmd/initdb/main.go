package main

import (
	"context"
	"fmt"
	"log"

	bindingsql "github.com/mattmoor/bindings/pkg/sql"

	"database/sql"
	_ "database/sql"

	_ "github.com/lib/pq"
)

type Entry struct {
	first string
	last  string
	email string
}

var entries = []Entry{
	{"Ville", "Aikas", "vaikas@vmware.com"},
	{"Scotty", "Nicholson", "snichols@vmware.com"},
	{"Matt", "Moore", "mattmoor@vmware.com"},
}

const fmtStr = "insert into users (first_name, last_name, email) values ('%s', '%s', '%s')"
const tableQuery = `SELECT tablename FROM pg_catalog.pg_tables where tablename = $1`

func checkTableExists(db *sql.DB) (bool, error) {
	rows, err := db.Query(tableQuery, "users")
	if err != nil {
		return false, err
	}
	defer rows.Close()
	for rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		if err != nil {
			return false, err
		}
		log.Printf("Found existing table: %q\n", tableName)
		return true, nil
	}
	return false, rows.Err()
}

func main() {
	db, err := bindingsql.Open(context.TODO(), "postgres")
	if err != nil {
		log.Fatal(err)
	}

	tableExists, err := checkTableExists(db)
	if err != nil {
		log.Fatal(err)
	}

	if !tableExists {
		log.Printf("Creating 'users' table")
		res, err := db.Exec(`CREATE TABLE users (
                  id SERIAL,
                  first_name VARCHAR(255),
                  last_name VARCHAR(255),
                  email TEXT
                );`)
		if err != nil {
			log.Fatalf("Failed to create table users: %v", err)
		}
		log.Printf("Created Table: %+v", res)
	}

	_, err = db.Exec("delete from users")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		_, err := db.Exec(fmt.Sprintf(fmtStr, e.first, e.last, e.email))
		if err != nil {
			log.Fatalf("Failed to insert entry %+v : %v", e, err)
		}
	}
}
