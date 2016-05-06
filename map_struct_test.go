package gsf

import (
	"testing"
	"time"
)

// go test -v base_test.go map_struct_test.go

type testXXXYYYZZZ struct {
	F0 string    `json:"f0"`
	F1 string    `json:"f1"`
	F2 int64     `json:"f2"`
	F3 time.Time `json:"f3"`
	F4 int16     `json:"f4"`
	F5 time.Time `json:"f5"`
}

func TestMapStruct(t *testing.T) {
	m := make(map[string]interface{})
	m["F0"] = "F0"
	m["F1"] = "F1"
	m["f2"] = 2
	m["F3"] = time.Now()
	m["F4"] = 4
	m["F5"] = time.Now()
	var xyx testXXXYYYZZZ
	logStruct(t, "Reference", m)
	logStruct(t, "Before", xyx)
	err := MapToStruct(m, &xyx)
	if err != nil {
		t.Error(err)
	}
	logStruct(t, "After", xyx)
}
