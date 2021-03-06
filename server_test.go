package main

import (
	"fmt"
	"testing"
)

var sanitizetests = []struct {
	in  JobParams
	out JobParams
}{
	{JobParams{}, JobParams{}},
	{JobParams{"", nil, "", "", ""}, JobParams{"", nil, "", "", ""}},
	{JobParams{"postgres://user:pass@host.com/db", nil, "", "", ""}, JobParams{"postgres://user:pass@host.com/db", nil, "", "", ""}},
	{JobParams{"", nil, "crane", "", ""}, JobParams{"", nil, "crane", "", ""}},
	{JobParams{"", nil, "cr@ne", "", ""}, JobParams{"", nil, "", "", ""}},
	{JobParams{"", nil, "", "sushi", ""}, JobParams{"", nil, "", "sushi", ""}},
	{JobParams{"", nil, "", "su$hi", ""}, JobParams{"", nil, "", "", ""}},
	{JobParams{"", nil, "", "", "HEROKU_POSTGRESQL_RED_URL"}, JobParams{"", nil, "", "", "HEROKU_POSTGRESQL_RED_URL"}},
	{JobParams{"", nil, "", "", "&EROKU_POSTGRESQL_RED_URL"}, JobParams{"", nil, "", "", ""}},
}

func TestSanitizeJopParams(t *testing.T) {
	for i, tt := range sanitizetests {
		tt.in.sanitize()
		if fmt.Sprintf("%v", tt.in) != fmt.Sprintf("%v", tt.out) {
			t.Errorf("%d. Expected to sanitize to %v, but was %v", i, tt.out, tt.in)
		}
	}
}

func TestRemovePassword(t *testing.T) {
	input := "postgres://user:pass@host:5432/dbname"
	expected := "postgres://user:@host:5432/dbname"
	output := removePassword(input)
	if output != expected {
		t.Errorf("Expected %v, but was %v", expected, output)
	}

	input = "postgres://usewhathappened"
	expected = ""
	output = removePassword(input)
	if output != expected {
		t.Errorf("Expected %v, but was %v", expected, output)
	}

}
