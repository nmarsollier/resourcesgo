package projects

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

	dbQueryRow = func(ctx context.Context, query string, args ...interface{}) (*Project, error) {
		assert.Equal(t, "test", args[0])
		return newProject("test", "Test"), nil
	}

	result, err := FindByID(tlog.TestContext, "test")

	assert.NoError(t, err)
	assert.Equal(t, "test", result.ID)
	assert.Equal(t, "Test", result.Name)
}

func TestFindByIdError(t *testing.T) {
	mockfunc := dbQueryRow
	defer func() { dbQueryRow = mockfunc }()

	dbQueryRow = func(ctx context.Context, query string, args ...interface{}) (*Project, error) {
		assert.Equal(t, "test", args[0])
		return nil, errs.NotFound
	}

	result, err := FindByID(tlog.TestContext, "test")

	assert.ErrorIs(t, err, errs.NotFound)
	assert.Nil(t, result)
}
