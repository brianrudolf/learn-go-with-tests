package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		// Check if field is struct, recursive through it

		// Check if field type is string
		if field.Kind() == reflect.String {
			fn(field.String())
		}
	}
}
