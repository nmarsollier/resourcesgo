package resources

import (
	"context"
	"testing"

	"github.com/nmarsollier/resourcesgo/internal/tools/errs"
	"github.com/nmarsollier/resourcesgo/tests/tlog"
	"github.com/stretchr/testify/assert"
)

func TestDeleteById(t *testing.T) {
	mockfunc := dbQueryFindBy
	defer func() { dbQueryFindBy = mockfunc }()

	mockfunc2 := dbDeleteExec
	defer func() { dbDeleteExec = mockfunc2 }()

	res := newTestResource()

	dbQueryFindBy = func(ctx context.Context, query string, args ...interface{}) (*Resource, error) {
		assert.Equal(t, res.ProjectID, args[0])
		assert.Equal(t, res.LanguageID, args[1])
		assert.Equal(t, res.SemVer, args[2])
		return res, nil
	}

	dbDeleteExec = func(ctx context.Context, query string, args ...interface{}) error {
		assert.Equal(t, res.ID, args[0])
		return nil
	}

	err := Delete(tlog.TestContext, res.ProjectID, res.LanguageID, res.SemVer)
	assert.NoError(t, err)
}

func TestDeleteByIdFindError(t *testing.T) {
	mockfunc := dbQueryFindBy
	defer func() { dbQueryFindBy = mockfunc }()

	res := newTestResource()

	dbQueryFindBy = func(ctx context.Context, query string, args ...interface{}) (*Resource, error) {
		assert.Equal(t, res.ProjectID, args[0])
		assert.Equal(t, res.LanguageID, args[1])
		assert.Equal(t, res.SemVer, args[2])
		return res, errs.NotFound
	}

	err := Delete(tlog.TestContext, res.ProjectID, res.LanguageID, res.SemVer)
	assert.ErrorIs(t, err, errs.NotFound)
}

func TestDeleteByIdDeleteError(t *testing.T) {
	mockfunc := dbQueryFindBy
	defer func() { dbQueryFindBy = mockfunc }()

	mockfunc2 := dbDeleteExec
	defer func() { dbDeleteExec = mockfunc2 }()

	res := newTestResource()

	dbQueryFindBy = func(ctx context.Context, query string, args ...interface{}) (*Resource, error) {
		assert.Equal(t, res.ProjectID, args[0])
		assert.Equal(t, res.LanguageID, args[1])
		assert.Equal(t, res.SemVer, args[2])
		return res, nil
	}

	dbDeleteExec = func(ctx context.Context, query string, args ...interface{}) error {
		return errs.Internal
	}

	err := Delete(tlog.TestContext, res.ProjectID, res.LanguageID, res.SemVer)
	assert.ErrorIs(t, err, errs.Internal)
}
