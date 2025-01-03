package projects

import (
	"context"
	"testing"

	"github.com/nmarsollier/resourcesgo/internal/tools/errs"
	"github.com/nmarsollier/resourcesgo/tests/terr"
	"github.com/nmarsollier/resourcesgo/tests/tlog"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	mockfunc := dbExec
	defer func() { dbExec = mockfunc }()

	dbExec = func(ctx context.Context, query string, args ...interface{}) error {
		assert.Equal(t, "test", args[0])
		assert.Equal(t, "Test Project", args[1])
		return nil
	}

	result, err := Create(tlog.TestContext, "test", "Test Project")

	assert.NoError(t, err)
	assert.Equal(t, "test", result)
}

func TestInvalidIdError(t *testing.T) {
	mockfunc := dbExec
	defer func() { dbExec = mockfunc }()

	dbExec = func(ctx context.Context, query string, args ...interface{}) error {
		return nil
	}

	result, err := Create(tlog.TestContext, "", "Test Project")

	assert.ErrorContains(t, err, "Field validation for 'ID' failed on the 'required'")
	assert.Empty(t, result)
}

func TestInvalidNameError(t *testing.T) {
	mockfunc := dbExec
	defer func() { dbExec = mockfunc }()

	dbExec = func(ctx context.Context, query string, args ...interface{}) error {
		return nil
	}

	result, err := Create(tlog.TestContext, "test", "")

	assert.ErrorContains(t, err, "Field validation for 'Name' failed on the 'required'")
	assert.Empty(t, result)
}

func TestCreateError(t *testing.T) {
	mockfunc := dbExec
	defer func() { dbExec = mockfunc }()

	dbExec = func(ctx context.Context, query string, args ...interface{}) error {
		return terr.PgErrorExist
	}

	result, err := Create(tlog.TestContext, "es", "Spanish")

	assert.ErrorIs(t, err, errs.AlreadyExist)
	assert.Empty(t, result)
}

func TestOtherError(t *testing.T) {
	mockfunc := dbExec
	defer func() { dbExec = mockfunc }()

	dbExec = func(ctx context.Context, query string, args ...interface{}) error {
		return errs.Internal
	}

	result, err := Create(tlog.TestContext, "es", "Spanish")

	assert.ErrorIs(t, err, errs.Internal)
	assert.Empty(t, result)
}
