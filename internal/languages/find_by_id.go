package languages

import (
	"github.com/nmarsollier/resourcesgo/internal/tools/db"
	"github.com/nmarsollier/resourcesgo/internal/tools/logx"
)

func FindByID(fields logx.Fields, id string) (*Language, error) {
	return db.QueryRow[Language](
		fields,
		`
      SELECT id, name, created, enabled
      FROM languages
      WHERE id = $1
    `,
		id,
	)
}
