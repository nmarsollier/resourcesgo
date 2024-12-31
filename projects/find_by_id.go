package projects

import (
	"github.com/nmarsollier/resourcesgo/tools/db"
	"github.com/nmarsollier/resourcesgo/tools/logx"
)

func FindByID(logenv logx.Fields, id string) (project *Project, err error) {
	project, err = db.QueryRow[Project](
		logenv,
		`
      SELECT id, name, created, enabled
      FROM projects
      WHERE id = $1
    `,
		id,
	)

	return
}
