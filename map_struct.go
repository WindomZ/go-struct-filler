package gsf

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func MapStruct(m map[string]interface{}, dst interface{}) error {
	return MapToStruct(m, dst)
}

func MapToStruct(m map[string]interface{}, dst interface{}) error {
	for k, v := range m {
		err := setField(v, dst, k)
		if err != nil {
			return err
		}
	}
	return nil
}

func setField(src interface{}, dest interface{}, name string) error {
	v := reflect.ValueOf(dest).Elem()
	fv := v.FieldByName(name)
	if !fv.IsValid() {
		for i := 0; i < v.NumField(); i++ {
			f := v.Type().Field(i)
			if strings.EqualFold(strings.ToLower(name), strings.ToLower(f.Tag.Get("json"))) {
				return setField(src, dest, f.Name)
			}
		}
		return fmt.Errorf("No such field: %s in obj", name)
	}
	if !fv.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}
	val := reflect.ValueOf(src)
	if fv.Kind() != val.Kind() {
		switch val.Kind() {
		case reflect.Float32, reflect.Float64:
			switch fv.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				fv.SetInt(int64(val.Float()))
				return nil
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			switch fv.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				fv.SetInt(int64(val.Int()))
				return nil
			}
		}
		return errors.New("Provided value type didn't match obj field type")
	} else {
		fv.Set(val)
	}
	return nil
}
