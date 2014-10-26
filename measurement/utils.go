package measurement

import (
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// From https://gist.github.com/tonyhb/5819315
// TODO: include omitempty support
func Flatten(i interface{}) (values url.Values) {
	values = url.Values{}
	iVal := reflect.ValueOf(i).Elem()
	typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		f := iVal.Field(i)

		tag := typ.Field(i).Tag.Get("param")
		// omitEmpty := false
		if strings.Contains(tag, ",") {
			tag = strings.Split(tag, ",")[0]
			// omitEmpty = true
		}
		// Convert each type into a string for the url.Values string map
		var v string
		switch f.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = strconv.FormatInt(f.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			v = strconv.FormatUint(f.Uint(), 10)
		case float32:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
		case float64:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
		case []byte:
			v = string(f.Bytes())
		case string:
			v = f.String()
		}

		values.Set(tag, v)

	}
	return
}
