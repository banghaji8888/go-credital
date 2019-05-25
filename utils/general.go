package utils

import (
	"reflect"
)

const tagName = "bson"

// GetFieldsName - get fields name
func GetFieldsName(m interface{}) []string {
	fields := []string{}
	t := reflect.TypeOf(m)

	// Iterate over all available fields and read the tag value
	for i := 0; i < t.NumField(); i++ {
		// Get the field, returns https://golang.org/pkg/reflect/#StructField
		field := t.Field(i)

		// Get the field tag value
		tag := field.Tag.Get(tagName)

		fields = append(fields, tag)
	}

	return fields
}
