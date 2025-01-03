package resources

import (
	"context"

	"github.com/nmarsollier/resourcesgo/internal/tools/db"
)

var dbDeleteExec = db.Exec

func Delete(ctx context.Context, project string, language string, semver string) error {
	resource, err := findBy(ctx, project, language, semver)
	if err != nil {
		return err
	}

	return dbDeleteExec(
		ctx,
		"UPDATE resources SET enabled=false WHERE id=$1",
		resource.ID,
	)
}
