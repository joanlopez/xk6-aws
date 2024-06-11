package aws

import (
	"errors"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
	"unicode"

	"github.com/grafana/sobek"
	"go.k6.io/k6/js/modules"
	"go.k6.io/k6/js/modules/k6/experimental/streams"
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

// Translates the given struct [from] to a [sobek.Value].
func toSobekObject(vu modules.VU, from any) (val sobek.Value, err error) {
	rt := vu.Runtime()
	rv := reflect.ValueOf(from)
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return sobek.Null(), nil
		}
		rv = rv.Elem()
	}

	// Validate that 'from' is a struct or a pointer to a struct.
	if rv.Kind() != reflect.Struct {
		return nil, errors.New("[from] parameter must be a struct or a pointer to a struct")
	}

	// Get the struct value and type.
	structType := rv.Type()

	// Create a new *sobek.Object.
	obj := rt.NewObject()

	// Iterate over the fields of the struct.
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		fieldType := structType.Field(i)
		fieldName := pascalToSnake(fieldType.Name)

		// Skip unexported fields.
		if !field.CanInterface() {
			continue
		}

		// Depending on the field type, we need to convert the Go value to a sobek.Value.
		var jsValue sobek.Value

		switch field.Kind() {
		case reflect.Ptr:
			if field.IsNil() {
				jsValue = sobek.Null()
				break
			}

			elem := field.Elem()
			switch elem.Kind() {
			case reflect.Struct:
				jsValue, err = toSobekObject(vu, elem.Addr().Interface())
			case reflect.String:
				jsValue = rt.ToValue(elem.String())
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				jsValue = rt.ToValue(elem.Int())
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				jsValue = rt.ToValue(elem.Uint())
			case reflect.Float32, reflect.Float64:
				jsValue = rt.ToValue(elem.Float())
			case reflect.Bool:
				jsValue = rt.ToValue(elem.Bool())
			default:
				err = fmt.Errorf("unsupported field type: %s", elem.Kind())
			}

		case reflect.Slice:
			length := field.Len()
			elems := make([]interface{}, length)
			for j := 0; j < length; j++ {
				elem := field.Index(j)
				var elemVal sobek.Value
				switch elem.Kind() {
				case reflect.Ptr:
					elemVal, err = toSobekObject(vu, elem.Interface())
				case reflect.Struct:
					elemVal, err = toSobekObject(vu, elem.Interface())
				case reflect.String:
					elemVal = rt.ToValue(elem.String())
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					elemVal = rt.ToValue(elem.Int())
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					elemVal = rt.ToValue(elem.Uint())
				case reflect.Float32, reflect.Float64:
					elemVal = rt.ToValue(elem.Float())
				case reflect.Bool:
					elemVal = rt.ToValue(elem.Bool())
				default:
					err = fmt.Errorf("unsupported field type: %s", elem.Kind())
				}

				// If there was any error while parsing the field,
				// we early stop the array population.
				if err != nil {
					break
				}

				elems[j] = elemVal
			}

			jsValue = rt.NewArray(elems...)

		case reflect.Interface:
			if fieldType.Type == reflect.TypeOf((*io.ReadCloser)(nil)).Elem() {
				jsValue = rt.ToValue(streams.NewReadableStreamFromReader(vu, field.Interface().(io.ReadCloser)))
			}
		case reflect.Struct:
			jsValue, err = toSobekObject(vu, field.Addr().Interface())
		case reflect.String:
			jsValue = rt.ToValue(field.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			jsValue = rt.ToValue(field.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		case reflect.Float32, reflect.Float64:
			jsValue = rt.ToValue(field.Float())
		case reflect.Bool:
			jsValue = rt.ToValue(field.Bool())
		default:
			jsValue = rt.ToValue(field.Interface())
		}

		// If there was any error while parsing the field,
		// we early stop and return the error.
		if err != nil {
			return nil, err
		}

		// The same if there is any error setting the field.
		if err := obj.Set(fieldName, jsValue); err != nil {
			return nil, err
		}
	}

	val = rt.ToValue(obj)
	return val, err
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
