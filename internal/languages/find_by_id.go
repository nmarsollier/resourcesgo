package languages

import (
	"context"

	"github.com/nmarsollier/resourcesgo/internal/tools/db"
)

var dbQueryRow = db.QueryRow[Language]

func FindByID(ctx context.Context, id string) (*Language, error) {
	return dbQueryRow(
		ctx,
		`
      SELECT id, name, created, enabled
      FROM languages
      WHERE id = $1
    `,
		id,
	)
}
