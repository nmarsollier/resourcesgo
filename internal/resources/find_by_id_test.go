package resources

import (
	"context"
	"testing"

	"github.com/nmarsollier/resourcesgo/internal/tools/errs"
	"github.com/nmarsollier/resourcesgo/tests/tlog"
	"github.com/stretchr/testify/assert"
)

func TestFindById(t *testing.T) {
	mockfunc := dbQueryFindByID
	defer func() { dbQueryFindByID = mockfunc }()

	res := newTestResource()

	dbQueryFindByID = func(ctx context.Context, query string, args ...interface{}) (*Resource, error) {
		assert.Equal(t, "test", args[0])
		return res, nil
	}

	result, err := FindByID(tlog.TestContext, "test")

	assert.NoError(t, err)
	assert.Equal(t, res.ID, result.ID)
}

func TestFindByIdError(t *testing.T) {
	mockfunc := dbQueryFindByID
	defer func() { dbQueryFindByID = mockfunc }()

	dbQueryFindByID = func(ctx context.Context, query string, args ...interface{}) (*Resource, error) {
		assert.Equal(t, "test", args[0])
		return nil, errs.NotFound
	}

	result, err := FindByID(tlog.TestContext, "test")

	assert.ErrorIs(t, err, errs.NotFound)
	assert.Nil(t, result)
}
