package resources

import (
	"context"

	"github.com/nmarsollier/resourcesgo/internal/tools/db"
)

var dbQueryFindByID = db.QueryRow[Resource]

func FindByID(ctx context.Context, id string) (*Resource, error) {
	return dbQueryFindByID(
		ctx,
		`
		SELECT id, project, language, sem_ver, values, created, enabled 
		FROM resources 
		WHERE id = $1
		`,
		id,
	)
}
