package jen

import (
	"testing"
)

var fmatter = `---
title: Test Markdown string with YAML (YML) Front Matter (Fmatter)
description: this data is used to test the Fmatter extractor
---
# Heading 1
A paragraph for heading one with Fmatter
## Heading 2
We have some data still no Fmatter
`

var nofmatter = `
# Heading 1
A paragraph for heading one with Fmatter
## Heading 2
We have some data still no Fmatter
`

func TestGetFrontMatter(t *testing.T) {
	fm := GetFrontMatter()
	if fm == nil {
		t.Errorf("our bugger is not empty")
	}
}

// TestNoMatter test markdown file with no matter
func TestNoFMatter(t *testing.T) {
	fm, body, err := DoFrontMatter(nofmatter)
	if err != nil {
		t.Fatalf("no front matter - expected nil error but got on: %v", err)
	}
	if fm != nil {
		t.Fatalf("expected no front matter got some ...")
	}
	if string(body) != nofmatter {
		t.Error("are body is not ours")
	}
}

// TestFmatter the front matter
func TestFmatter(t *testing.T) {
	fm, body, err := DoFrontMatter(fmatter)
	if err != nil {
		t.Fatalf("do front matter - expected nil error but got  %v", err)
	}
	if fm == nil {
		t.Fatalf("this can NOT go on also ...")
	}
	if body == "" {
		t.Error("we have failed our body")
	}
	if len(body) <= 0 {
		t.Error("we seem to have meta, but should not")
	}
}
