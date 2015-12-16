package gsf

import (
	"testing"
	"time"
)

func TestFillStruct(t *testing.T) {
	x := XXX{F0: "F0", F1: "F1", F2: 2, F3: time.Now(), F5: time.Now()}
	y := YYY{}
	z := ZZZ{XXX: x, YYY: y}
	logStruct(t, "Before", z)
	err := FillStruct(&z)
	if err != nil {
		t.Error(err)
	}
	logStruct(t, "After", z)
}
