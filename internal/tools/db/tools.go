package db

import (
	"cmp"
	"reflect"

	"github.com/jackc/pgx/v5"
)

func fieldAddrs[T any](names []string, result *T) (addrs []interface{}) {
	if reflect.TypeOf(*result).Kind() == reflect.Struct {
		destVal := reflect.ValueOf(result).Elem()
		destType := destVal.Type()

		fields := make(map[string]interface{}, destVal.NumField())
		for i := 0; i < destVal.NumField(); i++ {
			fields[fieldName(destType.Field(i))] = destVal.Field(i).Addr().Interface()
		}

		addrs = make([]interface{}, len(names))
		for i := 0; i < len(addrs); i++ {
			addrs[i] = fields[names[i]]
		}

		return
	} else {
		return []interface{}{result}
	}
}

func fieldName(field reflect.StructField) string {
	return cmp.Or(field.Tag.Get("db"), field.Name)
}

func columnNames(row pgx.Rows) (descriptions []string) {
	fieldDescriptions := row.FieldDescriptions()
	columnNames := make([]string, len(fieldDescriptions))
	for i, fd := range fieldDescriptions {
		columnNames[i] = string(fd.Name)
	}
	return columnNames
}
