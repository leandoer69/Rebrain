package convertor

import (
	"reflect"
	"strings"
)

func MapToStruct(mp map[string]interface{}, item interface{}) error {
	itemVal := reflect.ValueOf(item)
	elem := itemVal.Elem()

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Type().Field(i)

		tag := string(field.Tag)
		var key string
		if tag != "" {
			key = valueFromTag(tag)
		} else {
			key = field.Name
		}

		if _, exists := mp[key]; exists {
			elem.Field(i).Set(reflect.ValueOf(mp[key]))
		}

	}

	return nil
}

func valueFromTag(tag string) string {
	preValue := strings.Split(tag, ":")[1]
	value := strings.Split(preValue, "\"")[1]
	return value
}

func StructToMap(item interface{}) map[string]interface{} {
	itemValue := reflect.ValueOf(item)
	itemType := reflect.TypeOf(item)

	if itemType.Kind() != reflect.Struct {
		return nil
	}

	m := make(map[string]interface{})

	for i := 0; i < itemType.NumField(); i++ {
		field := itemType.Field(i)

		var key string
		tag := string(field.Tag)
		if tag != "" {
			key = valueFromTag(tag)
		} else {
			key = itemType.Field(i).Name
		}

		value := itemValue.Field(i)
		m[key] = value.Interface()
	}

	return m
}
