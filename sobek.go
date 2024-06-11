package aws

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
	"unicode"

	"github.com/grafana/sobek"
)

// Populates the given struct [to] from a *sobek.Object.
func fromSobekObject(rt *sobek.Runtime, obj *sobek.Object, to any) error {
	if obj == nil || sobek.IsUndefined(obj) || sobek.IsNull(obj) {
		return nil
	}

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

		// Get the corresponding JavaScript value from the sobek object
		jsValue := obj.Get(fieldName)
		if jsValue == nil || sobek.IsUndefined(jsValue) || sobek.IsNull(jsValue) {
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
				if err := fromSobekObject(rt, jsValue.ToObject(rt), field.Interface()); err != nil {
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
			case reflect.Bool:
				boolVal := jsValue.ToBoolean()
				field.Set(reflect.ValueOf(&boolVal))
			}
		case reflect.Slice:
			elemType := fieldType.Type.Elem()
			jsArray := jsValue.ToObject(rt)
			length := jsArray.Get("length").ToInteger()
			slice := reflect.MakeSlice(fieldType.Type, int(length), int(length))

			for j := 0; j < int(length); j++ {
				jsElem := jsArray.Get(strconv.Itoa(j))
				elem := reflect.New(elemType).Elem()

				switch elemType.Kind() {
				case reflect.Ptr:
					innerElemType := elemType.Elem()
					switch innerElemType.Kind() {
					case reflect.Struct:
						if err := fromSobekObject(rt, jsElem.ToObject(rt), elem.Addr().Interface()); err != nil {
							return err
						}
					case reflect.String:
						str := jsElem.String()
						elem.Set(reflect.ValueOf(&str))
					case reflect.Int32:
						if val, err := strconv.ParseInt(jsElem.String(), 10, 32); err == nil {
							intVal := int32(val)
							elem.Set(reflect.ValueOf(&intVal))
						}
					case reflect.Bool:
						boolVal := jsElem.ToBoolean()
						elem.Set(reflect.ValueOf(&boolVal))
					}
				case reflect.Struct:
					if err := fromSobekObject(rt, jsElem.ToObject(rt), elem.Addr().Interface()); err != nil {
						return err
					}
				case reflect.String:
					elem.SetString(jsElem.String())
				case reflect.Int, reflect.Int32, reflect.Int64:
					if val, err := strconv.ParseInt(jsElem.String(), 10, elemType.Bits()); err == nil {
						elem.SetInt(val)
					}
				case reflect.Float32, reflect.Float64:
					if val, err := strconv.ParseFloat(jsElem.String(), elemType.Bits()); err == nil {
						elem.SetFloat(val)
					}
				case reflect.Bool:
					elem.SetBool(jsElem.ToBoolean())
				}
				slice.Index(j).Set(elem)
			}
			field.Set(slice)
		case reflect.Struct:
			// Direct conversion for specific struct types, possibly using custom converters
			if err := fromSobekObject(rt, jsValue.ToObject(rt), field.Addr().Interface()); err != nil {
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
		case reflect.Bool:
			field.SetBool(jsValue.ToBoolean())
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

// isObject determines whether the given [sobek.Value] is a [sobek.Object] or not.
func isObject(val sobek.Value) bool {
	return val != nil && val.ExportType() != nil && val.ExportType().Kind() == reflect.Map
}
