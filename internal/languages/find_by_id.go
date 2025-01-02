package languages

import (
	"context"

	"github.com/nmarsollier/resourcesgo/internal/tools/db"
)

func FindByID(ctx context.Context, id string) (*Language, error) {
	return db.QueryRow[Language](
		ctx,
		`
      SELECT id, name, created, enabled
      FROM languages
      WHERE id = $1
    `,
		id,
	)
}
