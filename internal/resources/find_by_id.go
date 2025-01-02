package resources

import (
	"context"

	"github.com/nmarsollier/resourcesgo/internal/tools/db"
)

func FindByID(ctx context.Context, id string) (*Resource, error) {
	return db.QueryRow[Resource](
		ctx,
		`
		SELECT id, project, language, sem_ver, values, created, enabled 
		FROM resources 
		WHERE id = $1
		`,
		id,
	)
}
