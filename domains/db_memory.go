// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package domains

import (
	"errors"
	"fmt"
	"sort"
	"sync"
)

// Ensure memoryDB conforms to the domainDatabase interface.
var _ DomainDB = &memoryDB{}

// memoryDB is a simple in-memory persistence layer for domains.
type memoryDB struct {
	mu        sync.Mutex
	nextID    int64 // next ID to assign to a domain.
	domainMap DomainMap
}

func newMemoryDB() *memoryDB {
	return &memoryDB{
		domainMap: NewDomainMap(100),
		nextID:    1,
	}
}

// Close closes the database.
func (db *memoryDB) Close() {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.domainMap = nil
}

// Getdomain retrieves a domain by its ID.
func (db *memoryDB) GetDomain(name string) (*Domain, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	domain, ok := db.domainMap[id]
	if !ok {
		return nil, fmt.Errorf("memorydb: domain not found with ID %d", id)
	}
	return domain, nil
}

// AddDomain saves a given domain, assigning it a new ID.
func (db *memoryDB) AddDomain(b *Domain) (id int64, err error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	b.ID = db.nextID
	db.domains[b.ID] = b
	db.nextID++
	return b.ID, nil
}

// DelDomain removes a given domain by its ID.
func (db *memoryDB) DelDomain(id int64) error {
	if id == 0 {
		return errors.New("memorydb: domain with unassigned ID passed into DelDomain")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.domains[id]; !ok {
		return fmt.Errorf("memorydb: could not delete domain with ID %d, does not exist", id)
	}
	delete(db.domains, id)
	return nil
}

// UpdateDomain updates the entry for a given domain.
func (db *memoryDB) UpdateDomain(b *Domain) error {
	if b.ID == 0 {
		return errors.New("memorydb: domain with unassigned ID passed into UpdateDomain")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	db.domains[b.ID] = b
	return nil
}

// domainsByTitle implements sort.Interface, ordering domains by Title.
// https://golang.org/pkg/sort/#example__sortWrapper
type domainsByRegistrar []*Domain

func (s domainsByRegistrar) Less(i, j int) bool { return s[i].Name < s[j].Name }
func (s domainsByRegistrar) Len() int           { return len(s) }
func (s domainsByRegistrar) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// ListDomains returns a list of domains, ordered by title.
func (db *memoryDB) ListDomains() (*DomainTable, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	var domains []*Domain
	for _, b := range db.domains {
		domains = append(domains, b)
	}

	sort.Sort(domainsByRegistrar(domains))
	return domains, nil
}
