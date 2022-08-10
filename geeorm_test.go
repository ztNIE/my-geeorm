package geeorm

import (
	"errors"
	"geeorm/session"
	"reflect"

	_ "github.com/mattn/go-sqlite3"

	"testing"
)

func openDB(t *testing.T) *Engine {
	t.Helper()
	engine, err := NewEngine("sqlite3", "../gee.db")
	if err != nil {
		t.Fatal("fail to connect", err)
	}
	return engine
}

func TestNewEngine(t *testing.T) {
	engine := openDB(t)
	defer engine.Close()
}

type ForTest struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func TestEngine_Transaction(t *testing.T) {
	t.Run("rollback", func(t *testing.T) {
		transactionRollback(t)
	})
	t.Run("commit", func(t *testing.T) {
		transactionCommit(t)
	})
	t.Run("panic", func(t *testing.T) {
		transactionPanic(t)
	})
}

func transactionRollback(t *testing.T) {
	engine := openDB(t)
	defer engine.Close()
	s := engine.NewSession()
	_ = s.Model(&ForTest{}).DropTable()
	_, err := engine.Transaction(
		func(s *session.Session) (result interface{}, err error) {
			_ = s.Model(&ForTest{}).CreateTable()
			_, err = s.Insert(&ForTest{"ztNie", 23})
			return nil, errors.New("Error")
		},
	)
	if err == nil || s.HasTable() {
		t.Fatal("failed to rollback")
	}
}

func transactionCommit(t *testing.T) {
	engine := openDB(t)
	defer engine.Close()
	s := engine.NewSession()
	_ = s.Model(&ForTest{}).DropTable()
	_, err := engine.Transaction(func(s *session.Session) (result interface{}, err error) {
		_ = s.Model(&ForTest{}).CreateTable()
		_, err = s.Insert(&ForTest{"ztNie", 24})
		return
	})
	u := &ForTest{}
	_ = s.First(u)
	if err != nil || u.Name != "ztNie" {
		t.Fatal("failed to commit")
	}
}

func transactionPanic(t *testing.T) {
	engine := openDB(t)
	defer engine.Close()
	defer func() {
		if p := recover(); p != nil {
			t.Log("recover from panic")
		}
	}()
	s := engine.NewSession()
	s.Model(&ForTest{}).DropTable()
	engine.Transaction(func(s *session.Session) (result interface{}, err error) {
		s.Model(&ForTest{}).CreateTable()
		s.Insert(&ForTest{"ztNie", 24}, &ForTest{"Suan", 24})
		u := &ForTest{}
		err = s.Find(u)
		return
	})
}

func TestEngineMigrate(t *testing.T) {
	engine := openDB(t)
	defer engine.Close()
	s := engine.NewSession()
	s.Raw("DROP TABLE IF EXISTS ForTest;").Exec()
	s.Raw("CREATE TABLE ForTest(Name text PRIMARY KEY, xxx integer);").Exec()
	s.Raw("INSERT INTO ForTest(`Name`) values (?), (?)", "Tom", "Sam").Exec()

	engine.Migrate(&ForTest{})
	rows, _ := s.Raw("SELECT * FROM ForTest").QueryRows()
	columns, _ := rows.Columns()

	if !reflect.DeepEqual(columns, []string{"Name", "Age"}) {
		t.Fatal("Failed to migrate table User, got columns", columns)
	}
}
