package logx

import "context"

type contextKey string

const fieldsKey = contextKey("fields")

// CtxWithFields returns a new context with the provided fields added to it.
// If the context already contains fields, it returns the original context.
//
// Parameters:
//
//	ctx - The original context.
//	fields - The fields to add to the context.
//
// Returns:
//
//	A new context with the fields added, or the original context if it already contains fields.
func CtxWithFields(ctx context.Context, fields Fields) context.Context {
	_, ok := ctx.Value(fieldsKey).(Fields)
	if ok {
		return ctx
	}

	return context.WithValue(ctx, fieldsKey, fields)
}

// CtxFields retrieves the Fields from the given context.
// If the context does not contain any Fields, it returns a new Fields instance.
//
// Parameters:
//
//	ctx - the context from which to retrieve the Fields
//
// Returns:
//
//	Fields - the Fields retrieved from the context, or a new Fields instance if none are found
func CtxFields(ctx context.Context) Fields {
	fields, ok := ctx.Value(fieldsKey).(Fields)
	if !ok {
		return NewFields()
	}
	return fields
}
