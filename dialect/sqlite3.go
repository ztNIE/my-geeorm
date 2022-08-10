package dialect

import (
	"fmt"
	"reflect"
	"time"

	"geeorm/log"
)

type sqlite3 struct{}

// check if sqlite3 has implemented interface Dialect
var _ Dialect = (*sqlite3)(nil)

func init() {
	RegisterDialect("sqlite3", &sqlite3{})
}

func (s *sqlite3) DataTypeOf(typ reflect.Value) string {
	switch typ.Kind() {
	case reflect.Bool:
		return "bool"

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
		return "integer"

	case reflect.Int64, reflect.Uint64:
		return "bigint"

	case reflect.Float32, reflect.Float64:
		return "real"

	case reflect.String:
		return "text"

	case reflect.Array, reflect.Slice:
		return "blob"

	case reflect.Struct:
		if _, ok := typ.Interface().(time.Time); ok {
			return "datetime"
		}
	}

	log.Errorf("invalid sql type %s (%s)", typ.Type().Name(), typ.Kind())
	panic(fmt.Sprintf("invalid sql type %s (%s)", typ.Type().Name(), typ.Kind()))
}

func (s *sqlite3) TableExistSQL(tableName string) (string, []interface{}) {
	args := []interface{}{tableName}
	return "select name from sqlite_master where type = 'table' and name = ?", args
}
