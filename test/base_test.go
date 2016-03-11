package gsf_test

import (
	"bytes"
	"testing"
	"time"
)

type XXX struct {
	F0  string    `json:"f0"`
	F1  string    `json:"f1"`
	F2  int       `json:"f2"`
	F3  time.Time `json:"f3"`
	F5  time.Time `json:"f5"`
	F10 []string  `json:"f10"`
}

type YYY struct {
	F0 int
	F1 string
	F2 int
	XXX
	F5  time.Time
	F10 []string
}

type ZZZ struct {
	XXX
	YYY
}

func logStruct(t *testing.T, title string, s interface{}) {
	l := int((50 - len(title)) / 2)
	if l < 5 {
		l = 5
	}
	var b bytes.Buffer
	for i := 0; i < l; i++ {
		b.WriteByte('-')
	}
	tmp := b.String()
	b.WriteString(title)
	b.WriteString(tmp)
	t.Log(b.String())
	t.Logf("%#v", s)
	b.Reset()
	len := l*2 + len(title)
	for i := 0; i < len; i++ {
		b.WriteByte('-')
	}
	t.Log(b.String())
}
