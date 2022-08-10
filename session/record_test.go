package session

import (
	"testing"

	"geeorm/schema"
)

type ForTest schema.ForTest

var (
	test1 = &ForTest{"ad", 1}
	test2 = &ForTest{"ztNie", 13}
	test3 = &ForTest{"Suan", 12}
)

func testRecordInit(t *testing.T) *Session {
	t.Helper()
	s := NewSession().Model(&ForTest{})
	err1 := s.DropTable()
	err2 := s.CreateTable()
	_, err3 := s.Insert(test1, test2)
	if err1 != nil || err2 != nil || err3 != nil {
		t.Fatal("failed init test records")
	}
	return s
}

func TestLimit(t *testing.T) {
	s := testRecordInit(t)
	var tests []ForTest
	err := s.Limit(1).Find(&tests)
	if err != nil || len(tests) != 1 {
		t.Fatal("failed to query with limit condition")
	}
}

func TestUpdate(t *testing.T) {
	s := testRecordInit(t)
	affected, err := s.Where("Name = ?", "ztNie").Update("Age", 24)
	if err != nil {
		t.Fatal("Fail to update")
	}

	u := &ForTest{}
	s.Where("Age = ?", 24).First(u)
	if affected != 1 || u.Name != "ztNie" {
		t.Fatal("failed to update")
	}
}

func TestDeleteAndCount(t *testing.T) {
	s := testRecordInit(t)
	affected, err1 := s.Where("Name = ?", "ztNie").Delete()
	count, err2 := s.Count()
	if affected != 1 || count != 1 || err1 != nil || err2 != nil {
		t.Fatal("failed to delete or count")
	}

}

func TestInsert(t *testing.T) {
	s := testRecordInit(t)
	affected, err := s.Insert(test3)
	if err != nil || affected != 1 {
		t.Fatal("failed to create record")
	}
}

func TestFind(t *testing.T) {
	s := testRecordInit(t)
	var tests []ForTest
	if err := s.Find(&tests); err != nil || len(tests) != 2 {
		t.Fatal("failed to query all")
	}
}
