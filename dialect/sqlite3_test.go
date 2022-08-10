package dialect

import (
	"reflect"
	"testing"
	"time"
)

func TestDataTypeOf(t *testing.T) {
	var (
		boolValue   bool
		intValue    int
		bigintValue int64
		floatValue  float32
		strValue    string
		arrValue    []int
		timeValue   time.Time
	)

	eachTest(t, boolValue, "bool")
	eachTest(t, intValue, "integer")
	eachTest(t, bigintValue, "bigint")
	eachTest(t, floatValue, "real")
	eachTest(t, strValue, "text")
	eachTest(t, arrValue, "blob")
	eachTest(t, timeValue, "datetime")
}

func TestWrongData(t *testing.T) {
	type newStruct struct {
		f1 int
		f2 string
	}

	var structValue newStruct
	var result string
	defer func ()  {
		if r := recover(); r == nil {
			t.Fatalf("Failed to panic, get %s", result)
		}
	}()

	result = (&sqlite3{}).DataTypeOf(reflect.ValueOf(structValue))

}

func eachTest(t *testing.T, input interface{}, wanted string) {
	var result string
	if result = (&sqlite3{}).DataTypeOf(reflect.ValueOf(input)); result != wanted {
		t.Fatalf("Fail test, wanted %s, get %s", wanted, result)
	}

}
