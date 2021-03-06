package common

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

type QCloudArg interface {
	EncodeStructWithPrefix(prefix string, val reflect.Value, v *url.Values) error
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

func EncodeStruct(i interface{}, v *url.Values) error {
	val := reflect.ValueOf(i)
	return encodeStructWithPrefix("", val, v)
}

func encodeStructWithPrefix(prefix string, val reflect.Value, v *url.Values) error {
	if !reflect.Indirect(val).IsValid() {
		val = reflect.New(val.Type().Elem())
	}
	qcloudArg, ok := val.Interface().(QCloudArg)
	if ok {
		return qcloudArg.EncodeStructWithPrefix(prefix, val, v)
	}
	switch val.Kind() {
	case reflect.Struct:
		{
			typ := val.Type()
			for index := 0; index < val.NumField(); index++ {
				tag, opts := parseTag(typ.Field(index).Tag.Get("qcloud_arg"))
				fieldVal := val.Field(index)
				if fieldVal.Kind() == reflect.Ptr {
					if fieldVal.IsNil() {
						if opts.Contains("required") {
							return errors.New(fmt.Sprintf("field %s of %s should not be nil", tag, typ))
						}
						continue
					}
				}
				if opts.Contains("omitempty") && isEmptyValue(fieldVal) {
					continue
				}
				p := strings.Join([]string{prefix, tag}, ".")
				if err := encodeStructWithPrefix(p, fieldVal, v); err != nil {
					return err
				}
			}
		}
	case reflect.Array, reflect.Slice:
		{
			for index := 0; index < val.Len(); index++ {
				p := strings.Join([]string{prefix, fmt.Sprint(index)}, ".")
				if err := encodeStructWithPrefix(p, val.Index(index), v); err != nil {
					return err
				}
			}
		}
	case reflect.Ptr, reflect.Interface:
		if err := encodeStructWithPrefix(prefix, val.Elem(), v); err != nil {
			return err
		}
	case reflect.String:
		v.Set(strings.TrimLeft(prefix, "."), val.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v.Set(strings.TrimLeft(prefix, "."), fmt.Sprint(val))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		v.Set(strings.TrimLeft(prefix, "."), fmt.Sprint(val))
	case reflect.Float32, reflect.Float64:
		v.Set(strings.TrimLeft(prefix, "."), fmt.Sprint(val))
	case reflect.Bool:
		v.Set(strings.TrimLeft(prefix, "."), fmt.Sprint(val))
	}
	return nil
}

type tagOptions []string

func (tOpts tagOptions) Contains(opt string) bool {
	for _, o := range tOpts {
		if o == opt {
			return true
		}
	}
	return false
}

func parseTag(tag string) (string, tagOptions) {
	parts := strings.Split(tag, ",")
	return parts[0], parts[1:]
}
