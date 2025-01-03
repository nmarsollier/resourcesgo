package resources

import (
	"context"
	"testing"

	"github.com/nmarsollier/resourcesgo/internal/tools/errs"
	"github.com/nmarsollier/resourcesgo/tests/tlog"
	"github.com/stretchr/testify/assert"
)

func TestFindBy(t *testing.T) {
	mockfunc := dbQueryFindBy
	defer func() { dbQueryFindBy = mockfunc }()

	res := newTestResource()

	dbQueryFindBy = func(ctx context.Context, query string, args ...interface{}) (*Resource, error) {
		assert.Equal(t, res.ProjectID, args[0])
		assert.Equal(t, res.LanguageID, args[1])
		assert.Equal(t, res.SemVer, args[2])
		return res, nil
	}

	result, err := findBy(tlog.TestContext, res.ProjectID, res.LanguageID, res.SemVer)

	assert.NoError(t, err)
	assert.Equal(t, res.ID, result.ID)
}

func TestFindByError(t *testing.T) {
	mockfunc := dbQueryFindBy
	defer func() { dbQueryFindBy = mockfunc }()

	res := newTestResource()

	dbQueryFindBy = func(ctx context.Context, query string, args ...interface{}) (*Resource, error) {
		assert.Equal(t, res.ProjectID, args[0])
		assert.Equal(t, res.LanguageID, args[1])
		assert.Equal(t, res.SemVer, args[2])
		return nil, errs.NotFound
	}

	result, err := findBy(tlog.TestContext, res.ProjectID, res.LanguageID, res.SemVer)

	assert.ErrorIs(t, err, errs.NotFound)
	assert.Nil(t, result)
}
