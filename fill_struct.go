package gsf

import (
	"reflect"
	"strings"
	"time"
)

func FillStruct(dest interface{}) error {
	v, err := valueOf(dest)
	if err != nil {
		return err
	}
	tvs := getNameAndValues(v)
	for i, tv1 := range tvs {
		for j, tv2 := range tvs {
			if i >= j {
				continue
			}
			if tv1.Type.Type != tv2.Type.Type {
				continue
			}
			if strings.EqualFold(tv1.Type.Name, tv2.Type.Name) {
				v1 := tv1.Value
				v2 := tv2.Value
				switch v1.Kind() {
				case reflect.Bool:
					if v1.Bool() {
						v2.Set(v1)
					} else if v2.Bool() {
						v1.Set(v2)
					}
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					if v1.Int() > v2.Int() {
						v2.Set(v1)
					} else {
						v1.Set(v2)
					}
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					if v1.Uint() > v2.Uint() {
						v2.Set(v1)
					} else {
						v1.Set(v2)
					}
				case reflect.Float32, reflect.Float64:
					if v1.Float() > v2.Float() {
						v2.Set(v1)
					} else {
						v1.Set(v2)
					}
				case reflect.String:
					if len(v1.String()) > len(v2.String()) {
						v2.Set(v1)
					} else {
						v1.Set(v2)
					}
				case reflect.Struct:
					switch item1 := v1.Interface().(type) {
					case time.Time:
						if item1.After(v2.Interface().(time.Time)) {
							v2.Set(v1)
						} else {
							v1.Set(v2)
						}
					}
				}
			}
		}
	}
	return nil
}

type structTypeValue struct {
	Type  reflect.StructField
	Value reflect.Value
}

func getNameAndValues(v reflect.Value) []structTypeValue {
	tvs := make([]structTypeValue, v.NumMethod())
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Anonymous {
			vf := v.Field(i)
			if vf.CanSet() && t.Field(i).Type.Kind() == reflect.Struct {
				tvs = append(tvs, getNameAndValues(vf)...)
			}
		} else {
			tvs = append(tvs, structTypeValue{Type: t.Field(i), Value: v.Field(i)})
		}
	}
	return tvs
}
