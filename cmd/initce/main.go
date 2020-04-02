package main

import (
	"context"
	"os"

	"log"

	bindingsql "github.com/mattmoor/bindings/pkg/sql"

	_ "database/sql"

	_ "github.com/lib/pq"
)

func main() {
	dbName := os.Getenv("DBNAME")
	if dbName == "" {
		log.Fatal("No database name specified, bailing")
	}

	db, err := bindingsql.Open(context.TODO(), "postgres", dbName)
	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Exec(`CREATE TABLE events (
                  id SERIAL,
                  type text,
                  source text,
                  eventid text,
                  time timestamptz,
                  schema text,
                  contenttype text,
                  data text
                );`)
	if err != nil {
		log.Fatalf("Failed to create table events: %v", err)
	}

	log.Printf("Created Table: %+v", res)
}
