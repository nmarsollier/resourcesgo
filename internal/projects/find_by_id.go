package projects

import (
	"context"

	"github.com/nmarsollier/resourcesgo/internal/tools/db"
)

var dbQueryRow = db.QueryRow[Project]

func FindByID(ctx context.Context, id string) (*Project, error) {
	return dbQueryRow(
		ctx,
		`
      SELECT id, name, created, enabled
      FROM projects
      WHERE id = $1
    `,
		id,
	)
}
