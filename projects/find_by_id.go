package projects

import (
	"github.com/nmarsollier/resourcesgo/tools/db"
	"github.com/nmarsollier/resourcesgo/tools/logx"
)

func FindByID(fields logx.Fields, id string) (*Project, error) {
	return db.QueryRow[Project](
		fields,
		`
      SELECT id, name, created, enabled
      FROM projects
      WHERE id = $1
    `,
		id,
	)
}
