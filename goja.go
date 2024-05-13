package aws

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
	"unicode"

	"github.com/dop251/goja"
)

// Populates the given struct [to] from a *goja.Object.
func fromGojaObject(obj *goja.Object, to any) error {
	// Validate that 'to' is a pointer to a struct
	rv := reflect.ValueOf(to)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return errors.New("[to] parameter must be a pointer to a struct")
	}

	// Get the struct value and type
	structValue := rv.Elem()
	structType := structValue.Type()

	// Iterate over the fields of the struct
	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)
		if !field.CanSet() {
			continue // Skip unexported fields
		}

		fieldType := structType.Field(i)
		fieldName := pascalToSnake(fieldType.Name)

		// Get the corresponding JavaScript value from the goja object
		jsValue := obj.Get(fieldName)
		if jsValue == nil || goja.IsUndefined(jsValue) || goja.IsNull(jsValue) {
			continue
		}

		// Depending on the field type, we need to convert the JS value appropriately
		switch field.Kind() {
		case reflect.Ptr:
			// Handle pointer types to basic types like *string, *int32
			elemType := fieldType.Type.Elem()
			switch elemType.Kind() {
			case reflect.String:
				str := jsValue.String()
				field.Set(reflect.ValueOf(&str))
			case reflect.Int32:
				if val, err := strconv.ParseInt(jsValue.String(), 10, 32); err == nil {
					intVal := int32(val)
					field.Set(reflect.ValueOf(&intVal))
				}
			}
		case reflect.Slice:
			// Additional logic needed to handle slice types
		case reflect.Struct:
			// Direct conversion for specific struct types, possibly using custom converters
		case reflect.String:
			field.SetString(jsValue.String())
		case reflect.Int, reflect.Int32, reflect.Int64:
			if val, err := strconv.ParseInt(jsValue.String(), 10, field.Type().Bits()); err == nil {
				field.SetInt(val)
			}
		case reflect.Float32, reflect.Float64:
			if val, err := strconv.ParseFloat(jsValue.String(), field.Type().Bits()); err == nil {
				field.SetFloat(val)
			}
		}
	}

	return nil
}

// pascalToSnake converts a PascalCase string to snake_case.
func pascalToSnake(s string) string {
	var snakeCase strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				snakeCase.WriteByte('_')
			}
			snakeCase.WriteRune(unicode.ToLower(r))
		} else {
			snakeCase.WriteRune(r)
		}
	}
	return snakeCase.String()
}
