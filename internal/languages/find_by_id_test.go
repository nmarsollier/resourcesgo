package languages

import (
	"context"
	"testing"

	"github.com/nmarsollier/resourcesgo/internal/tools/errs"
	"github.com/nmarsollier/resourcesgo/tests/tlog"
	"github.com/stretchr/testify/assert"
)

func TestFindById(t *testing.T) {
	mockfunc := dbQueryRow
	defer func() { dbQueryRow = mockfunc }()

	dbQueryRow = func(ctx context.Context, query string, args ...interface{}) (*Language, error) {
		assert.Equal(t, "es", args[0])
		return newLanguage("es", "Test"), nil
	}

	result, err := FindByID(tlog.TestContext, "es")

	assert.NoError(t, err)
	assert.Equal(t, "es", result.ID)
	assert.Equal(t, "Test", result.Name)
}

func TestFindByIdError(t *testing.T) {
	mockfunc := dbQueryRow
	defer func() { dbQueryRow = mockfunc }()

	dbQueryRow = func(ctx context.Context, query string, args ...interface{}) (*Language, error) {
		assert.Equal(t, "es", args[0])
		return nil, errs.NotFound
	}

	result, err := FindByID(tlog.TestContext, "es")

	assert.ErrorIs(t, err, errs.NotFound)
	assert.Nil(t, result)
}
