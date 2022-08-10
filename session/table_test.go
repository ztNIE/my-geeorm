package session

import (
	"geeorm/schema"
	"testing"
)

func TestCreateTable(t *testing.T) {
	s := NewSession().Model(&schema.ForTest{})
	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HasTable() {
		t.Fatal("Failed to create table ForTest")
	}
}
