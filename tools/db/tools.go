package db

import "reflect"

func fieldAddrs[T any](result *T) (addrs []interface{}) {
	if reflect.TypeOf(*result).Kind() == reflect.Struct {
		destVal := reflect.ValueOf(result).Elem()
		addrs = make([]interface{}, destVal.NumField())
		for i := 0; i < destVal.NumField(); i++ {
			addrs[i] = destVal.Field(i).Addr().Interface()
		}
		return
	} else {
		return []interface{}{result}
	}
}
