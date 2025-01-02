package projects

import (
	"context"

	"github.com/nmarsollier/resourcesgo/internal/tools/db"
)

func FindByID(ctx context.Context, id string) (*Project, error) {
	return db.QueryRow[Project](
		ctx,
		`
      SELECT id, name, created, enabled
      FROM projects
      WHERE id = $1
    `,
		id,
	)
}
