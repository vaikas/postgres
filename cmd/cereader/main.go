package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"log"

	bindingsql "github.com/mattmoor/bindings/pkg/sql"

	_ "database/sql"

	_ "github.com/lib/pq"
)

type reader struct {
	dbName string
}

func (a *reader) handler(w http.ResponseWriter, r *http.Request) {
	log.Print("helloworld: received a request")
	db, err := bindingsql.Open(context.TODO(), "postgres", a.dbName)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT time, type, source FROM events order by time")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var eventtype, source string
		var ts time.Time
		err = rows.Scan(&ts, &eventtype, &source)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "%q %q %q\n", ts, source, eventtype)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbName := os.Getenv("DBNAME")
	if dbName == "" {
		dbName = "cloudevents"
	}

	r := reader{dbName: dbName}
	http.HandleFunc("/", r.handler)
	log.Printf("helloworld: listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
