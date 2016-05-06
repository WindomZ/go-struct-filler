package gsf

import (
	"testing"
	"time"
)

// go test -v -test.run conv_struct_test.go

func TestConvertStruct(t *testing.T) {
	x := testXXX{F0: "0", F1: "1", F2: 1, F3: time.Now(), F5: time.Now(), F10: []string{"111", "222", "333"}}
	y := testYYY{}
	logStruct(t, "Reference", x)
	logStruct(t, "Before", y)
	err := ConvertStruct(&x, &y)
	if err != nil {
		t.Error(err)
	}
	logStruct(t, "After", y)
}
