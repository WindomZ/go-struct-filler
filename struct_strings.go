package gsf

import (
	"reflect"
	"strconv"
	"strings"
	"time"
)

func StructToStringSlice(src interface{}) (result []string, err error) {
	v, err := valueOf(src)
	if err != nil {
		return
	}
	tvs := getNameAndValues(v)
	result = make([]string, len(tvs))
	for i, tv := range tvs {
		switch tv.Value.Kind() {
		case reflect.Bool:
			result[i] = strconv.FormatBool(tv.Value.Bool())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			result[i] = strconv.FormatInt(tv.Value.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			result[i] = strconv.FormatUint(tv.Value.Uint(), 10)
		case reflect.Float32, reflect.Float64:
			result[i] = strconv.FormatFloat(tv.Value.Float(), 'f', 20, 64)
		case reflect.String:
			result[i] = tv.Value.String()
		case reflect.Struct:
			switch item := tv.Value.Interface().(type) {
			case time.Time:
				result[i] = item.Format(time.RFC3339)
			}
		}
	}
	return
}

type structTagValue struct {
	Tag   string
	Value reflect.Value
}

func StructToStringMap(src interface{}, tags ...string) (result map[string]string, err error) {
	var tag string
	if len(tags) == 0 {
		tag = "json"
	} else {
		tag = tags[0]
	}
	srcV, err := valueOf(src)
	if err != nil {
		return
	}
	srcTVs := getNameAndValues(srcV)
	tvs := make([]structTagValue, len(srcTVs))
	for _, sTV := range srcTVs {
		t := sTV.Type.Tag.Get(tag)
		if len(t) == 0 || strings.EqualFold(t, "-") {
			continue
		}
		tvs = append(tvs, structTagValue{Tag: t, Value: sTV.Value})
	}
	result = make(map[string]string, len(tvs))
	for _, tv := range tvs {
		switch tv.Value.Kind() {
		case reflect.Bool:
			result[tv.Tag] = strconv.FormatBool(tv.Value.Bool())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			result[tv.Tag] = strconv.FormatInt(tv.Value.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			result[tv.Tag] = strconv.FormatUint(tv.Value.Uint(), 10)
		case reflect.Float32, reflect.Float64:
			result[tv.Tag] = strconv.FormatFloat(tv.Value.Float(), 'f', 20, 64)
		case reflect.String:
			result[tv.Tag] = tv.Value.String()
		case reflect.Struct:
			switch item := tv.Value.Interface().(type) {
			case time.Time:
				result[tv.Tag] = item.Format(time.RFC3339)
			}
		}
	}
	return
}
