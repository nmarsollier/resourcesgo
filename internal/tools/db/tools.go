package db

import (
	"cmp"
	"context"
	"errors"
	"reflect"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/nmarsollier/resourcesgo/internal/tools/strs"
)

// structFieldPointers takes a slice of field names and a pointer to a struct of any type T,
// and returns a slice of pointers to the corresponding fields in the struct in the given names order.
//
// Returns:
//   - A slice of interface{} containing pointers to the specified fields in the struct, or
//     a slice containing the result itself if it is not a struct.
func structFieldPointers[T any](names []string, result *T) (addrs []interface{}) {
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

// fieldName returns the name of a struct field based on its "db" tag.
// If the "db" tag is not present, it returns the field's name.
func fieldName(field reflect.StructField) string {
	return cmp.Or(field.Tag.Get("db"), field.Name)
}

// columnNames extracts and returns the column names from the given pgx.Rows.
// It iterates over the field descriptions of the rows and converts each field name
// to a string, returning a slice of these column names.
func columnNames(row pgx.Rows) (descriptions []string) {
	fieldDescriptions := row.FieldDescriptions()
	columnNames := make([]string, len(fieldDescriptions))
	for i, fd := range fieldDescriptions {
		columnNames[i] = string(fd.Name)
	}
	return columnNames
}

// checkConnectionError checks the provided error to determine if it is a
// PostgreSQL connection error or a context deadline/cancellation error.
// If error sets the instance to nil, to force reconnect.
func checkConnectionError(err error) {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "08000", "08003", "08006", "08001", "08004", "08007", "08P01":
			instance = nil
		}
	}

	if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
		instance = nil
	}
}

// ErrorCode extracts and returns the error code from a PostgreSQL error.
//
// Returns:
//
//	int - the extracted error code, or 0 if the error is not a PostgreSQL error
func ErrorCode(err error) int {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return strs.AtoiZero(pgErr.Code)
	}
	return 0
}
