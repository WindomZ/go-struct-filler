package gsf

import (
	. "github.com/WindomZ/go-struct-filler"
	"testing"
	"time"
)

func TestStructToStringArray(t *testing.T) {
	x := XXX{F0: "F0", F1: "F1", F2: 2, F3: time.Now(), F5: time.Now()}
	logStruct(t, "Before", x)
	rs, err := StructToStringArray(&x)
	if err != nil {
		t.Error(err)
	}
	logStruct(t, "After", rs)
}
