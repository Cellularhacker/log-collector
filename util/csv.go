package util

import (
	"fmt"
	"reflect"
	"strings"
)

func ExtractCSVHeaders(v interface{}) string {
	headers := make([]string, 0)

	val := reflect.ValueOf(v)
	for i := 0; i < val.Type().NumField(); i++ {
		h := val.Type().Field(i).Tag.Get("csv")
		if h == "" {
			h = val.Type().Field(i).Tag.Get("json")
		}

		headers = append(headers, h)
	}

	return fmt.Sprintf("%s\n", strings.Join(headers, ","))
}

func ExtractCSVData(v interface{}) string {
	data := make([]string, 0)

	val := reflect.ValueOf(v)
	for i := 0; i < val.NumField(); i++ {
		data = append(data, fmt.Sprintf("%v", val.Field(i).Interface()))
	}

	return fmt.Sprintf("%s\n", strings.Join(data, ","))
}
