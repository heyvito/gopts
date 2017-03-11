package interman

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

// LoadEnvs is an alias to LoadEnvsWithPrefix("", baseObj)
func LoadEnvs(baseObj interface{}) interface{} {
	return LoadEnvsWithPrefix("", baseObj)
}

// LoadEnvsWithPrefix loads data from environment variables into a managed
// provided by you. Useful for loading settings stored in the OS environment.
//
// The following types are currently supported:
// - Bool
// - Int64
// - []string
// - String
//
// Preflight will use the fields of the provided object to match environment
// keys. For instance, it expects that a field named `APIKey" is available
// as API_KEY in the environment. Assuming the `prefix" argument is filled
// with a non-empty value `PREFLIGHT", the library will look for a env named
// PREFLIGHT_API_KEY.
//
// Setting a `interman:"-"` tag prevents Preflight from filling the associated
// field.
//
// A `default" tag may also be set. Its value will be set to the field in case
// it is absent from the environment vars set.
//
// For instance, take the following struct:
//
// type Settings struct {
//     Username 	string
//     SecretKey 	string	`default:"s3cr37"`
//     AutoRestart 	bool	`default:"true"`
//     IgnoredField string	`interman:"-"`
// }
//
// and the following environment variables:
// - PREF_USERNAME=Rob
// - PREF_AUTO_RESTART=false
//
// running the following snippet:
//
// 	settings := interman.LoadEnvsWithPrefix("pref", Settings{}).(Settings)
//
// will yield the following result:
// {
// 		Username: 		"Rob",
// 		SecretKey: 		"s3cr37",
// 		AutoRestart: 	false,
// 		IgnoredField:	""
// }
//
func LoadEnvsWithPrefix(prefix string, baseObj interface{}) interface{} {
	baseObjType := reflect.TypeOf(baseObj)
	objPtr := reflect.New(baseObjType)
	obj := reflect.Indirect(objPtr)
	for i := 0; i < baseObjType.NumField(); i++ {
		field := baseObjType.Field(i)
		var def *string
		if prefData, ok := field.Tag.Lookup("interman"); ok {
			if prefData == "-" {
				continue
			}
		}
		if alias, ok := field.Tag.Lookup("default"); ok {
			if alias == "" {
				def = nil
			} else {
				def = &alias
			}
		}
		envName := snakeCase(field.Name)
		if prefix != "" {
			envName = fmt.Sprintf("%s_%s", prefix, envName)
		}
		envValue := os.Getenv(strings.ToUpper(envName))
		if envValue == "" && def != nil {
			envValue = *def
		}

		targetField := obj.Field(i)
		switch field.Type.Kind() {
		case reflect.Bool:
			targetField.SetBool(boolFromString(envValue))
		case reflect.Int64:
			if intValue, err := strconv.Atoi(envValue); err == nil {
				targetField.SetInt(int64(intValue))
			}
		case reflect.Slice:
			if reflect.TypeOf(field.Type.Elem()).Kind() == reflect.Ptr {
				// Assuming as String, this will have to change anytime soon
				// we support other slice types
				values := strings.Split(envValue, ",")
				targetField.Set(reflect.ValueOf(values))
				break
			}
			fallthrough
		default:
			targetField.SetString(envValue)
		}
	}
	return objPtr.Elem().Convert(baseObjType).Interface()
}

func snakeCase(in string) string {
	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) && runes[i-1] != '_' {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return string(out)
}

func boolFromString(in string) bool {
	in = strings.ToLower(in)
	return in == "yes" || in == "true" || in == "y" || in == "1" || in == "on"
}
