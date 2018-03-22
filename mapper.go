// Package mapper allows you to easily map from one struct to another using tags.
package mapper

import (
	"fmt"
	"reflect"
	"strings"
)

var (
	errNotPointer        = "expected parameter 'out' to be of type pointer, but was %s"
	errInvalidMap        = "invalid map specified: %s"
	errInvalidConversion = "cannot map %s to %s"
	errInvalidType       = "type %s is not supported"
)

// Map reads the struct tags on the 'in' parameter and attempts to write the values to
// the provided 'out' parameter. The 'out' parameter must be a pointer to a struct.
// To define a mappable struct, declare your struct's fields with tags in the following
// format:
//
// `toMap:"<target>:<field>"`
//
// The source and target fields must be of the same type. Currently supported types are:
// string, int, float32, float64, bool & map.
func Map(in, out interface{}) error {
	// Get the type of the output struct.
	outType := reflect.TypeOf(out)

	// If the target struct is not a pointer to a struct,
	// return an error.
	if outType.Kind() != reflect.Ptr {
		return fmt.Errorf(errNotPointer, outType)
	}

	inType := reflect.TypeOf(in)

	// Get the fields of the source struct
	fields := getFields(inType)

	// Create a map of the source field names to target field names
	fieldMap, err := getFieldMap(fields, outType)

	if err != nil {
		return err
	}

	// Copy values accordingly.
	mapValues(in, out, fieldMap)

	return nil
}

func mapValues(in, out interface{}, m map[string]string) error {
	// Get the values of the source & target.
	inValue := reflect.ValueOf(in)
	outValue := reflect.Indirect(reflect.ValueOf(out))

	for source, target := range m {
		// Get the source and target fields
		inField := inValue.FieldByName(source)
		outField := outValue.FieldByName(target)

		// Ensure the types match. Otherwise, return an error.
		if inField.Type() != outField.Type() {
			return fmt.Errorf(errInvalidConversion, inField.Type(), outField.Type())
		}

		kind := inField.Type().Kind()

		// Copy the value from the source to the target
		// field based on the type of field.
		switch kind {
		case reflect.String:
			outField.SetString(inField.String())
		case reflect.Int:
			outField.SetInt(inField.Int())
		case reflect.Bool:
			outField.SetBool(inField.Bool())
		case reflect.Float32:
			outField.SetFloat(inField.Float())
		case reflect.Float64:
			outField.SetFloat(inField.Float())
		case reflect.Map:
			outField.Set(inField)
		default:
			return fmt.Errorf(errInvalidType, kind.String())
		}
	}

	// Copy the new interface to the old one.
	out = outValue.Interface()

	return nil
}

func getFields(t reflect.Type) []reflect.StructField {
	var out []reflect.StructField

	// Get all fields for the given type.
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		out = append(out, field)
	}

	return out
}

func getFieldMap(fs []reflect.StructField, outType reflect.Type) (map[string]string, error) {
	out := make(map[string]string)

	for _, field := range fs {
		// Obtain the tag string
		tag := field.Tag.Get("map")

		if tag == "" {
			continue
		}

		// Split the tags by struct type. Multiple maps can be added using a
		// semicolon. For example: `mapTo:"A:Field;B:Field"`
		maps := strings.Split(tag, ";")

		for _, m := range maps {
			// Seperate the type & field
			parts := strings.Split(m, ":")

			if len(parts) != 2 {
				return nil, fmt.Errorf(errInvalidMap, tag)
			}

			// Get the simplified type name.
			typeName := getTypeName(outType)

			// If the type doesn't match what we're mapping to
			// skip it.
			if parts[0] != typeName {
				continue
			}

			// Otherwise, add it to the map.
			out[field.Name] = parts[1]
		}
	}

	return out, nil
}

func getTypeName(t reflect.Type) string {
	// Simplify the type name to obtain just
	// the name of the struct.
	full := t.String()
	parts := strings.Split(full, ".")

	return parts[len(parts)-1]
}
