package util

import (
	"fmt"
	"net/url"
	"reflect"
)

// StructToQuery converts struct into URL query string
func StructToQuery(i interface{}) *string {
	values := url.Values{}

	structToQueryParams(reflect.ValueOf(i), "", values)

	result := values.Encode()

	return &result
}

// structToQueryParams converts struct into string recursively.
func structToQueryParams(v reflect.Value, prefix string, values url.Values) {
	t := v.Type()

	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return
		}

		v = v.Elem()
		t = v.Type()
	}

	switch v.Kind() {

	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			value := v.Field(i)

			// Skip field query:"-"
			tag := field.Tag.Get("query")
			if tag == "-" {
				continue
			}

			if tag == "" {
				tag = field.Name
			}

			if prefix != "" {
				tag = prefix + "[" + tag + "]"
			}

			structToQueryParams(value, tag, values)
		}
	case reflect.Map:
		iterator := v.MapRange()

		for iterator.Next() {
			key := iterator.Key()

			if key.Kind() != reflect.String {
				continue
			}

			tag := prefix
			if tag == "" {
				tag = key.String()
			} else {
				tag = prefix + "[" + key.String() + "]"
			}
			structToQueryParams(iterator.Value(), tag, values)
		}
	default:
		if v.IsZero() {
			return
		}

		var str string

		switch v.Kind() {
		case reflect.Bool:
			str = fmt.Sprintf("%v", v.Bool())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			str = fmt.Sprintf("%v", v.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			str = fmt.Sprintf("%v", v.Uint())
		case reflect.Float32, reflect.Float64:
			str = fmt.Sprintf("%v", v.Float())
		case reflect.String:
			str = v.String()
		default:
			return
		}

		if prefix != "" {
			values.Add(prefix, str)
		}
	}

}
