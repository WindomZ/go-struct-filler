package gsf

import (
	"testing"
	"time"
)

// go test -v -test.run copy_struct_test.go

func TestCopyStruct(t *testing.T) {
	x := testXXX{F0: "0", F1: "1", F2: 1, F3: time.Now(), F5: time.Now(), F10: []string{"111", "222", "333"}}
	y := testXXX{}
	logStruct(t, "Reference", x)
	logStruct(t, "Before", y)
	if err := CopyStruct(&x, &y); err != nil {
		t.Error(err)
	}
	logStruct(t, "After", y)
}
