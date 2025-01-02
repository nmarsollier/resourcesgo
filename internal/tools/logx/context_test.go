package logx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCtxWithFields(t *testing.T) {
	ctx := context.Background()
	fields := NewFields().Add("key", "value")

	// Agregar campos al contexto
	ctxWithFields := CtxWithFields(ctx, fields)

	// Verificar que los campos se hayan agregado correctamente
	retrievedFields := ctxWithFields.Value(fieldsKey).(Fields)
	assert.Equal(t, fields, retrievedFields)

	// Verificar que agregar campos a un contexto que ya tiene campos no los sobrescriba
	ctxWithFieldsAgain := CtxWithFields(ctxWithFields, NewFields().Add("newKey", "newValue"))
	retrievedFieldsAgain := ctxWithFieldsAgain.Value(fieldsKey).(Fields)
	assert.Equal(t, fields, retrievedFieldsAgain)
}
