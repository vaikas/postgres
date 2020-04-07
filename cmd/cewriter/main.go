package main

import (
	"context"
	"log"

	cloudevents "github.com/cloudevents/sdk-go/v1"
	bindingsql "github.com/mattmoor/bindings/pkg/sql"
	"knative.dev/eventing/pkg/kncloudevents"

	_ "database/sql"

	_ "github.com/lib/pq"
)

const insertStmt = "insert into events (type, source, eventid, time, schema, contenttype, data) values ($1, $2, $3, $4, $5, $6, $7)"

func store(event cloudevents.Event) error {
	ctx := event.Context.AsV1()
	var data string
	if err := event.DataAs(&data); err != nil {
		log.Printf("Got data error: %v\n", err)
		return err
	}

	db, err := bindingsql.Open(context.TODO(), "postgres")
	if err != nil {
		log.Printf("Failed to open db : %v", err)
		return err
	}

	_, err = db.Exec(insertStmt, ctx.GetType(), ctx.GetSource(), ctx.GetID(), ctx.GetTime(), ctx.GetDataSchema(), ctx.GetDataContentType(), data)
	if err != nil {
		log.Printf("Failed to insert entry %+v : %v", event.String(), err)
		return err
	}
	return nil
}

func main() {
	c, err := kncloudevents.NewDefaultClient()
	if err != nil {
		log.Fatal("Failed to create client, ", err)
	}

	log.Fatal(c.StartReceiver(context.Background(), store))
}
