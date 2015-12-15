package gsf

import (
	"time"
)

type XXX struct {
	F0 string    `json:"f0"`
	F1 string    `json:"f1"`
	F2 int       `json:"f2"`
	F3 time.Time `json:"f3"`
	F5 time.Time `json:"f5"`
}

type YYY struct {
	F0 int
	F1 string
	F2 int
	XXX
	F5 time.Time
}

type ZZZ struct {
	XXX
	YYY
}
