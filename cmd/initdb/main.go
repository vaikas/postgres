package main

import (
	"context"
	"fmt"
	"os"

	"log"

	bindingsql "github.com/mattmoor/bindings/pkg/sql"

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

func main() {
	dbName := os.Getenv("DBNAME")
	if dbName == "" {
		dbName = "knative"
	}

	db, err := bindingsql.Open(context.TODO(), "postgres", dbName)
	if err != nil {
		log.Fatal(err)
	}

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

	for _, e := range entries {
		res, err = db.Exec(fmt.Sprintf(fmtStr, e.first, e.last, e.email))
		if err != nil {
			log.Fatalf("Failed to insert entry %+v : %v", e, err)
		}
	}
}
