package schema

import (
	"geeorm/dialect"
	"testing"
)

var TestDial, _ = dialect.GetDialect("sqlite3")

func TestParse(t *testing.T) {
	schema := Parse(&ForTest{}, TestDial)

	if schema.Name != "ForTest" || len(schema.Fields) != 2 {
		t.Fatal("Failed to parse ForTest struct")
	}

	if schema.GetField("Name").Tag != "PRIMARY KEY" {
		t.Fatal("failed to parse primary key")
	}

}

/* func TestRecordValues(t *testing.T) {
 *     schema := Parse(&ForTest{}, TestDial)
 * } */
