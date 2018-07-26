// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package domains

import (
	"fmt"

	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
)

// datastoreDB persists Domains to Cloud Datastore.
// https://cloud.google.com/datastore/docs/concepts/overview
type datastoreDB struct {
	client *datastore.Client
}

// Ensure datastoreDB conforms to the DomainDB interface.
var _ DomainDB = &datastoreDB{}

// newDatastoreDB creates a new DomainDatabase backed by Cloud Datastore.
// See the datastore and google packages for details on creating a suitable Client:
// https://godoc.org/cloud.google.com/go/datastore
func newDatastoreDB(client *datastore.Client) (DomainDB, error) {
	ctx := context.Background()
	// Verify that we can communicate and authenticate with the datastore service.
	t, err := client.NewTransaction(ctx)
	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not connect: %v", err)
	}
	if err := t.Rollback(); err != nil {
		return nil, fmt.Errorf("datastoredb: could not connect: %v", err)
	}
	return &datastoreDB{client: client}, nil
}

// Close closes the database.
func (db *datastoreDB) Close() {
	// No op.
}

func (db *datastoreDB) datastoreKey(name string) *datastore.Key {
	return datastore.IDKey("Domain", name, nil)
}

// GetDomain retrieves a Domain by its ID.
func (db *datastoreDB) GetDomain(name string) (*Domain, error) {
	ctx := context.Background()
	k := db.datastoreKey(name)
	domain := &Domain{}
	if err := db.client.Get(ctx, k, domain); err != nil {
		return nil, fmt.Errorf("datastoredb: could not get domain: %v", err)
	}
	//domain.ID = id
	return domain, nil
}

// AddDomain saves a given domain, assigning it a new ID.
func (db *datastoreDB) AddDomain(d *Domain) (id int64, err error) {
	ctx := context.Background()
	k := datastore.IncompleteKey("domain", nil)
	k, err = db.client.Put(ctx, k, b)
	if err != nil {
		return 0, fmt.Errorf("datastoredb: could not put domain: %v", err)
	}
	return k.ID, nil
}

// DelDomain removes a given domain by its ID.
func (db *datastoreDB) DelDomain(id int64) error {
	ctx := context.Background()
	k := db.datastoreKey(id)
	if err := db.client.Delete(ctx, k); err != nil {
		return fmt.Errorf("datastoredb: could not delete domain: %v", err)
	}
	return nil
}

// UpdateDomain updates the entry for a given domain.
func (db *datastoreDB) UpdateDomain(b *Domain) error {
	ctx := context.Background()
	k := db.datastoreKey(b.ID)
	if _, err := db.client.Put(ctx, k, b); err != nil {
		return fmt.Errorf("datastoredb: could not update domain: %v", err)
	}
	return nil
}

// ListDomains returns a list of domains, ordered by title.
func (db *datastoreDB) ListDomains() (dt *DomainTable, err error) {
	ctx := context.Background()

	// Create our domain table
	dt = &DomainTable{
		Name:      "DomainDB",
		DomainMap: make(DomainMap, 100), // TODO be smarter than this
	}

	// create a new query to GetAll domains
	q := datastore.NewQuery("domain").Order("Title")

	domains := make([]*Domain, 0)
	keys, err := db.client.GetAll(ctx, q, domains)
	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not list domains: %v", err)
	}
	for _, d := range domains {
		dt.DomainMap[d.Name] = *d
	}
	return dt, nil
}
