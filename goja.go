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
func fromGojaObject(rt *goja.Runtime, obj *goja.Object, to any) error {
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
			elemType := fieldType.Type.Elem()
			switch elemType.Kind() {
			// Handle pointer to struct
			case reflect.Struct:
				if field.IsNil() {
					field.Set(reflect.New(elemType))
				}
				if err := fromGojaObject(rt, jsValue.ToObject(rt), field.Interface()); err != nil {
					return err
				}
			// Handle pointer types to basic types like *string, *int32
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
			elemType := fieldType.Type.Elem()
			if elemType.Kind() == reflect.Struct {
				jsArray := jsValue.ToObject(rt)
				length := int(jsArray.Get("length").ToInteger())
				slice := reflect.MakeSlice(fieldType.Type, length, length)
				for j := 0; j < length; j++ {
					jsElem := jsArray.Get(strconv.Itoa(j))
					elem := reflect.New(elemType).Elem()
					if err := fromGojaObject(rt, jsElem.ToObject(rt), elem.Addr().Interface()); err != nil {
						return err
					}
					slice.Index(j).Set(elem)
				}
				field.Set(slice)
			}
		case reflect.Struct:
			// Direct conversion for specific struct types, possibly using custom converters
			if err := fromGojaObject(rt, jsValue.ToObject(rt), field.Addr().Interface()); err != nil {
				return err
			}
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
