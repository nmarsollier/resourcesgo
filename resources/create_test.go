package resources

import (
	"testing"

	"github.com/nmarsollier/resourcesgo/tests/terr"
	"github.com/nmarsollier/resourcesgo/tools/errs"
	"github.com/nmarsollier/resourcesgo/tools/logx"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	mockfunc := dbExec
	defer func() { dbExec = mockfunc }()

	var res = newTestResource()

	dbExec = func(fields logx.Fields, query string, args ...interface{}) error {
		assert.Equal(t, res.ID, args[0])
		assert.Equal(t, res.ProjectID, args[1])
		assert.Equal(t, res.LanguageID, args[2])
		assert.Equal(t, res.SemVer, args[3])
		assert.Equal(t, res.Values, args[4])
		assert.Equal(t, res.Created, args[5])
		assert.Equal(t, res.Enabled, args[6])

		return nil
	}

	result, err := Create(logx.Fields{}, res)

	assert.NoError(t, err)
	assert.Equal(t, res.ID, result)
}

func TestCreateError(t *testing.T) {
	mockfunc := dbExec
	defer func() { dbExec = mockfunc }()

	var res = newTestResource()

	dbExec = func(f logx.Fields, query string, args ...interface{}) error {
		return terr.PgErrorExist
	}

	result, err := Create(logx.Fields{}, res)

	assert.ErrorIs(t, err, errs.AlreadyExist)
	assert.Empty(t, result)
}

func TestConstraintError(t *testing.T) {
	mockfunc := dbExec
	defer func() { dbExec = mockfunc }()
	var res = newTestResource()

	dbExec = func(f logx.Fields, query string, args ...interface{}) error {
		return terr.PgErrorForeign
	}

	result, err := Create(logx.Fields{}, res)

	assert.ErrorIs(t, err, errs.ErrProjectNotExist)
	assert.Empty(t, result)
}

func TestOtherError(t *testing.T) {
	mockfunc := dbExec
	defer func() { dbExec = mockfunc }()
	var res = newTestResource()

	dbExec = func(f logx.Fields, query string, args ...interface{}) error {
		return errs.Internal
	}

	result, err := Create(logx.Fields{}, res)

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
