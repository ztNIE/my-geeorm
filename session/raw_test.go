package session

import (
	"database/sql"
	"geeorm/dialect"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var (
	TestDB      *sql.DB
	TestDial, _ = dialect.GetDialect("sqlite3")
)

func TestMain(m *testing.M) {
	TestDB, _ = sql.Open("sqlite3", "../../gee.db")
	code := m.Run()
	_ = TestDB.Close()
	os.Exit(code)
}

func NewSession() *Session {
	return New(TestDB, TestDial)
}

func setupTestSession(s *Session) {
	_, _ = s.Raw("DROP TABLE IF EXISTS ForTest;").Exec()
	_, _ = s.Raw("CREATE TABLE ForTest(Name text);").Exec()
}

func TestSessionExec(t *testing.T) {
	s := NewSession()
	setupTestSession(s)
	result, _ := s.Raw("INSERT INTO ForTest(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	if count, err := result.RowsAffected(); err != nil || count != 2 {
		t.Fatal("expect 2 but got", count)
	}
}

func TestSessionQueryRow(t *testing.T) {
	s := NewSession()
	setupTestSession(s)
	s.Raw("INSERT INTO ForTest(`Name`) values (?), (?)", "ztNIE", "Chris").Exec()
	result := s.Raw("SELECT count(*) FROM ForTest WHERE Name = (?)", "ztNIE").QueryRow()
	var count int
	if err := result.Scan(&count); err != nil {
		t.Fatal("failed to query db", err)
	}

	if count != 1 {
		t.Fatalf("too much rows affected: %d", count)
	}
}
