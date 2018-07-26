// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package domains

import (
	"testing"
	//"github.com/GoogleCloudPlatform/golang-samples/internal/testutil"
)

func testDB(t *testing.T, db DomainDB) {
	defer db.Close()

	b := &Domain{
		Name: "example.com",
	}

	id, err := db.AddDomain(b)
	if err != nil {
		t.Fatal(err)
	}

	b.ID = id
	if err := db.UpdateDomain(b); err != nil {
		t.Error(err)
	}

	gotdomain, err := db.GetDomain(id)
	if err != nil {
		t.Error(err)
	}
	if gotdomain == nil {
		t.Error("Can't find no domain")
	}
	if err := db.DelDomain(id); err != nil {
		t.Error(err)
	}

	if _, err := db.GetDomain(id); err == nil {
		t.Error("want non-nil err")
	}
}

func TestMemoryDB(t *testing.T) {
	testDB(t, newMemoryDB())
}
