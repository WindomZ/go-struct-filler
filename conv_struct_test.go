package gsf

import (
	"testing"
	"time"
)

func TestConvertStruct(t *testing.T) {
	x := XXX{F0: "0", F1: "1", F2: 1, F3: time.Now(), F5: time.Now()}
	y := YYY{}
	err := ConvertStruct(&x, &y)
	if err != nil {
		t.Error(err)
	}
	t.Log(x)
	t.Log(y)
}
