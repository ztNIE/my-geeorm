package clause

import (
	"reflect"
	"testing"
)

func testSelect(t *testing.T) {
	var clause Clause
	clause.Set(LIMIT, 3)
	clause.Set(SELECT, "ForTest", []string{"*"})
	clause.Set(WHERE, "Name = ?", "Tom")
	clause.Set(ORDERBY, "Age ASC")
	sql, vars := clause.Build(SELECT, WHERE, ORDERBY, LIMIT)
	t.Log(sql, vars)
	wantedStr := "SELECT * FROM ForTest WHERE Name = ? ORDER BY Age ASC LIMIT ?" 
	if sql != wantedStr {
		t.Fatalf("failed to build SQL\nwanted: %s\nresult:%s\n", wantedStr, sql)
	}
	if !reflect.DeepEqual(vars, []interface{}{"Tom", 3}) {
		t.Fatal("failed to build SQLVars")
	}
}

func TestClauseBuild(t *testing.T) {
	t.Run("select", func (t *testing.T) {
		testSelect(t)
	}) 
}
