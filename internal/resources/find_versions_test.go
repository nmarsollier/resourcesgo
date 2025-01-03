package resources

import (
	"context"
	"testing"

	"github.com/nmarsollier/resourcesgo/internal/tools/errs"
	"github.com/nmarsollier/resourcesgo/tests/tlog"
	"github.com/stretchr/testify/assert"
)

func TestFindVersions(t *testing.T) {
	mockfunc := dbQueryFindVersions
	defer func() { dbQueryFindVersions = mockfunc }()

	dbQueryFindVersions = func(ctx context.Context, query string, args ...interface{}) ([]*string, error) {
		assert.Equal(t, "test", args[0])
		assert.Equal(t, "en", args[1])

		a := "123"
		b := "456"
		return []*string{&a, &b}, nil
	}

	result, err := FindVersions(tlog.TestContext, "test", "en")

	assert.NoError(t, err)
	assert.Equal(t, result[0], "123")
	assert.Equal(t, result[1], "456")
}

func TestFindVersionsError(t *testing.T) {
	mockfunc := dbQueryFindVersions
	defer func() { dbQueryFindVersions = mockfunc }()

	dbQueryFindVersions = func(ctx context.Context, query string, args ...interface{}) ([]*string, error) {
		assert.Equal(t, "test", args[0])
		assert.Equal(t, "en", args[1])
		return nil, errs.NotFound
	}

	result, err := FindVersions(tlog.TestContext, "test", "en")

	assert.ErrorIs(t, err, errs.NotFound)
	assert.Empty(t, result)
}
