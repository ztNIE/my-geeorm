package dialect

import (
	"reflect"
	"testing"
)

func TestDialect(t *testing.T) {
	RegisterDialect("sqlite3", &sqlite3{})
	dialect, ok := GetDialect("sqlite3")
	if !ok {
		t.Fatalf("RegisterDialect failed")
	}

	if reflect.Indirect(reflect.ValueOf(dialect)).Type().String() != "dialect.sqlite3" {
		t.Fatalf("Dialect fail! Want %s, get %s", "dialect.sqlite3", reflect.Indirect(reflect.ValueOf(dialect)).Type().String())
	}
}
