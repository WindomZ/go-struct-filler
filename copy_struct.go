package gsf

import (
	"bytes"
	"encoding/gob"
	"errors"
)

func CopyStruct(src, dst interface{}) error {
	if src == nil || dst == nil {
		return errors.New("gsf: cannot copy nil pointer")
	}
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}
