package db

import (
	"context"
	"reflect"

	"github.com/nmarsollier/resourcesgo/tools/logx"
)

func QueryRow[T any](logenv logx.Fields, query string, args ...interface{}) (*T, error) {
	conn, err := getDBConn(logenv)
	if err != nil {
		logx.Error(logenv, err)
		return nil, err
	}

	row := conn.QueryRow(context.Background(), query, args...)

	result := new(T)
	if err := row.Scan(fieldAddrs(result)...); err != nil {
		checkConnectionError(err)
		logx.Error(logenv, err)
		return nil, err
	}

	return result, nil
}

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
