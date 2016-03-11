package gsf_test

import (
	. "github.com/WindomZ/go-struct-filler"
	"testing"
	"time"
)

// go test -v base_test.go conv_struct_test.go

func TestConvertStruct(t *testing.T) {
	x := XXX{F0: "0", F1: "1", F2: 1, F3: time.Now(), F5: time.Now(), F10: []string{"111", "222", "333"}}
	y := YYY{}
	logStruct(t, "Reference", x)
	logStruct(t, "Before", y)
	err := ConvertStruct(&x, &y)
	if err != nil {
		t.Error(err)
	}
	logStruct(t, "After", y)
}
