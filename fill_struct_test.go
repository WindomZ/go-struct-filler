package gsf

import (
	"testing"
	"time"
)

// go test -v base_test.go fill_struct_test.go

func TestFillStruct(t *testing.T) {
	x := testXXX{F0: "F0", F1: "F1", F2: 2, F3: time.Now(), F5: time.Now()}
	y := testYYY{}
	z := testZZZ{testXXX: x, testYYY: y}
	logStruct(t, "Before", z)
	err := FillStruct(&z)
	if err != nil {
		t.Error(err)
	}
	logStruct(t, "After", z)
}
