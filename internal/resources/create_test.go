package resources

import (
	"context"
	"testing"

	"github.com/nmarsollier/resourcesgo/internal/tools/errs"
	"github.com/nmarsollier/resourcesgo/internal/tools/logx"
	"github.com/nmarsollier/resourcesgo/tests/terr"
	"github.com/nmarsollier/resourcesgo/tests/tlog"
	"github.com/stretchr/testify/assert"
)

var tContext = logx.CtxWithFields(context.Background(), logx.NewFields())

func TestCreate(t *testing.T) {
	mockfunc := dbExec
	defer func() { dbExec = mockfunc }()

	var res = newTestResource()

	dbExec = func(ctx context.Context, query string, args ...interface{}) error {
		assert.Equal(t, res.ID, args[0])
		assert.Equal(t, res.ProjectID, args[1])
		assert.Equal(t, res.LanguageID, args[2])
		assert.Equal(t, res.SemVer, args[3])
		assert.Equal(t, res.Values, args[4])
		assert.Equal(t, res.Created, args[5])
		assert.Equal(t, res.Enabled, args[6])

		return nil
	}

	result, err := Create(tContext, res)

	assert.NoError(t, err)
	assert.Equal(t, res.ID, result)
}

func TestCreateError(t *testing.T) {
	mockfunc := dbExec
	defer func() { dbExec = mockfunc }()

	var res = newTestResource()

	dbExec = func(f context.Context, query string, args ...interface{}) error {
		return terr.PgErrorExist
	}

	result, err := Create(tContext, res)

	assert.ErrorIs(t, err, errs.AlreadyExist)
	assert.Empty(t, result)
}

func TestProjectNotExistError(t *testing.T) {
	mockfunc := dbExec
	defer func() { dbExec = mockfunc }()
	var res = newTestResource()

	dbExec = func(f context.Context, query string, args ...interface{}) error {
		return terr.PgErrorProjectForeign
	}

	result, err := Create(tContext, res)

	assert.ErrorIs(t, err, errs.ErrProjectNotExist)
	assert.Empty(t, result)
}

func TestLanguageNotExistError(t *testing.T) {
	mockfunc := dbExec
	defer func() { dbExec = mockfunc }()
	var res = newTestResource()

	dbExec = func(f context.Context, query string, args ...interface{}) error {
		return terr.PgErrorLanguageForeign
	}

	result, err := Create(tContext, res)

	assert.ErrorIs(t, err, errs.ErrLanguageNotExist)
	assert.Empty(t, result)
}

func TestOtherError(t *testing.T) {
	mockfunc := dbExec
	defer func() { dbExec = mockfunc }()
	var res = newTestResource()

	dbExec = func(f context.Context, query string, args ...interface{}) error {
		return errs.Internal
	}

	result, err := Create(tContext, res)

	assert.ErrorIs(t, err, errs.Internal)
	assert.Empty(t, result)
}

func newTestResource() *Resource {
	values := make(map[string]string)
	values["test"] = "Test Project"

	return NewResource(
		"test",
		"es",
		"0.0.1",
		values,
	)
}

func TestInvalidSemverError(t *testing.T) {
	mockfunc := dbExec
	defer func() { dbExec = mockfunc }()

	var res = newTestResource()
	res.SemVer = ""

	dbExec = func(ctx context.Context, query string, args ...interface{}) error {
		return nil
	}

	result, err := Create(tlog.TestContext, res)

	assert.ErrorContains(t, err, "Field validation for 'SemVer' failed on the 'required'")
	assert.Empty(t, result)
}
