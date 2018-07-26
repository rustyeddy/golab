package domains

import (
	"testing"
)

func Test_BadCSV(t *testing.T) {
	fname := "/thisname-does-not-exist/foobar"
	recs, err := ReadCSVFile(&fname)
	if recs != nil {
		t.Error("expected recs to be nil got %v", recs)
	}
	if err == nil {
		t.Error("expected a NOEXIST error got nothing")
	}
}

func Test_GoodCSV(t *testing.T) {
	fname := "etc/domains.csv"
	recs, err := ReadCSVFile(&fname)
	if recs == nil {
		t.Error("expected recs but got nothing")
	}
	if err == nil {
		t.Error("expected recs got an error %v", err)
	}
	if len(recs) < 1 {
		t.Error("hmmm, we seemed to failed collecting domains")
	}
}
