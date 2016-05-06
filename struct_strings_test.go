package gsf

import (
	"testing"
	"time"
)

// go test -v -test.run struct_strings_test.go

func TestStructToStringSlice(t *testing.T) {
	x := testXXX{F0: "F0", F1: "F1", F2: 2, F3: time.Now(), F5: time.Now()}
	logStruct(t, "Before", x)
	rs, err := StructToStringSlice(&x)
	if err != nil {
		t.Error(err)
	}
	logStruct(t, "After", rs)
}

func TestStructToStringMap(t *testing.T) {
	x := testXXX{F0: "F000", F1: "F111", F2: 3, F3: time.Now(), F5: time.Now()}
	logStruct(t, "Before", x)
	rs, err := StructToStringMap(&x)
	if err != nil {
		t.Error(err)
	}
	logStruct(t, "After", rs)
}
