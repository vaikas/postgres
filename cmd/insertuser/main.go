package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	bindingsql "github.com/mattmoor/bindings/pkg/sql"

	_ "database/sql"

	_ "github.com/lib/pq"
)

const insertStmt = "insert into users (first_name, last_name, email) values ($1, $2, $3)"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handler)
	log.Printf("insertuser: listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	firsts, ok := r.URL.Query()["first"]
	if !ok || len(firsts[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Missing required query param 'first'"))
		return
	}

	lasts, ok := r.URL.Query()["last"]
	if !ok || len(lasts[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Missing required query param 'last'"))
		return
	}

	emails, ok := r.URL.Query()["email"]
	if !ok || len(emails[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Missing required query param 'email'"))
		return
	}

	first := firsts[0]
	last := lasts[0]
	email := emails[0]

	db, err := bindingsql.Open(context.TODO(), "postgres")
	defer db.Close()
	if err != nil {
		log.Fatalf("Failed to open db : %v\n", err)
	}

	_, err = db.Exec(insertStmt, first, last, email)
	if err != nil {
		log.Fatalf("Failed to insert entry: %v", err)
	}
}
