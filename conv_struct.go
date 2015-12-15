package gsf

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func ConvertStruct(src, dest interface{}) error {
	t1, err := typeOf(src)
	if err != nil {
		return err
	}
	t2, err := typeOf(dest)
	if err != nil {
		return err
	}
	v1, err := valueOf(src)
	if err != nil {
		return err
	}
	v2, err := valueOf(dest)
	if err != nil {
		return err
	}
	for i := 0; i < t1.NumField(); i++ {
		f1 := t1.Field(i)
		if f1.Anonymous {
			continue
		}
		for j := 0; j < t2.NumField(); j++ {
			f2 := t2.Field(j)
			if f2.Anonymous {
				continue
			}
			if f1.Type.Kind() != f2.Type.Kind() {
				continue
			}
			if strings.EqualFold(f1.Name, f2.Name) {
				v2.Field(j).Set(v1.Field(i))
			}
		}
	}
	return nil
}

func mustPtr(s interface{}) error {
	if s == nil {
		return errors.New(fmt.Sprintf("Must not be nil"))
	}
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Ptr {
		return errors.New(fmt.Sprintf("%v must be reflect.Ptr", v.Type().Name()))
	}
	if v.IsNil() {
		return errors.New(fmt.Sprintf("Must not be nil"))
	}
	return nil
}

func typeOf(s interface{}) (reflect.Type, error) {
	err := mustPtr(s)
	if err != nil {
		return reflect.TypeOf(s), err
	}
	return reflect.TypeOf(s).Elem(), nil
}

func valueOf(s interface{}) (reflect.Value, error) {
	err := mustPtr(s)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(s).Elem(), nil
}
