package resources

import (
	"context"

	"github.com/nmarsollier/resourcesgo/internal/tools/db"
)

func Delete(ctx context.Context, project string, language string, semver string) {
	if resource, err := findBy(ctx, project, language, semver); err == nil {
		db.Exec(
			ctx,
			"UPDATE resources SET enabled=false WHERE id=$1",
			resource.ID,
		)
	}
}
