package utils

import "reflect"

func StructToAttr(item interface{}) *map[string][]string {
	res := make(map[string][]string)
	if item == nil {
		return &res
	}
	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get("attr")
		field := reflectValue.Field(i).Interface()
		if v, ok := field.(string); ok {
			res[tag] = []string{v}
		}
	}
	return &res
}
