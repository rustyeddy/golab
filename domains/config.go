package domains

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
)

// =============== Global Variables ===============
var (
	DB DomainDB
)

func init() {
	var err error
	DB, err = configureDatastoreDB("clowdops")
	if err != nil {
		log.Fatal(err)
	}
}

// configureDatastore
func configureDatastoreDB(projectID string) (DomainDB, error) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	db, err := newDatastoreDB(client)
	return db, err
}

func init() {
}
