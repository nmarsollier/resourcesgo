package resources

import (
	"context"

	"github.com/nmarsollier/resourcesgo/internal/tools/db"
)

func findBy(ctx context.Context, project string, language string, semVer string) (*Resource, error) {
	return db.QueryRow[Resource](
		ctx,
		`
		SELECT id, project, language, sem_ver, values, created, enabled 
		FROM resources 
		WHERE project = $1 AND language = $2 AND sem_ver = $3
		`,
		project, language, semVer,
	)
}