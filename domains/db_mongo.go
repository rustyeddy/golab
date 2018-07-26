// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package domains

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type mongoDB struct {
	conn *mgo.Session
	c    *mgo.Collection
}

// Ensure mongoDB conforms to the domainDatabase interface.
var _ DomainDB = &mongoDB{}

// newMongoDB creates a new domainDatabase backed by a given Mongo server,
// authenticated with given credentials.
func newMongoDB(addr string, cred *mgo.Credential) (DomainDB, error) {
	conn, err := mgo.Dial(addr)
	if err != nil {
		return nil, fmt.Errorf("mongo: could not dial: %v", err)
	}

	if cred != nil {
		if err := conn.Login(cred); err != nil {
			return nil, err
		}
	}

	return &mongoDB{
		conn: conn,
		c:    conn.DB("domainshelf").C("domains"),
	}, nil
}

// Close closes the database.
func (db *mongoDB) Close() {
	db.conn.Close()
}

// GetDomain retrieves a domain by its ID.
func (db *mongoDB) GetDomain(id int64) (*Domain, error) {
	b := &Domain{}
	if err := db.c.Find(bson.D{{Name: "id", Value: id}}).One(b); err != nil {
		return nil, err
	}
	return b, nil
}

var maxRand = big.NewInt(1<<63 - 1)

// randomID returns a positive number that fits within an int64.
func randomID() (int64, error) {
	// Get a random number within the range [0, 1<<63-1)
	n, err := rand.Int(rand.Reader, maxRand)
	if err != nil {
		return 0, err
	}
	// Don't assign 0.
	return n.Int64() + 1, nil
}

// AddDomain saves a given domain, assigning it a new ID.
func (db *mongoDB) AddDomain(b *Domain) (id int64, err error) {
	id, err = randomID()
	if err != nil {
		return 0, fmt.Errorf("mongodb: could not assign an new ID: %v", err)
	}

	b.ID = id
	if err := db.c.Insert(b); err != nil {
		return 0, fmt.Errorf("mongodb: could not add domain: %v", err)
	}
	return id, nil
}

// DelDomain removes a given domain by its ID.
func (db *mongoDB) DelDomain(id int64) error {
	return db.c.Remove(bson.D{{Name: "id", Value: id}})
}

// UpdateDomain updates the entry for a given domain.
func (db *mongoDB) UpdateDomain(b *Domain) error {
	return db.c.Update(bson.D{{Name: "id", Value: b.ID}}, b)
}

// ListDomains returns a list of domains, ordered by title.
func (db *mongoDB) ListDomains() ([]*Domain, error) {
	var result []*Domain
	if err := db.c.Find(nil).Sort("title").All(&result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListDomainsCreatedBy returns a list of domains, ordered by title, filtered by
// the user who created the domain entry.
func (db *mongoDB) ListDomainsByRegistrar(regID string) ([]*Domain, error) {
	var result []*Domain
	if err := db.c.Find(bson.D{{Name: "registrar", Value: regID}}).Sort("title").All(&result); err != nil {
		return nil, err
	}
	return result, nil
}
