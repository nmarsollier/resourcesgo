package resources

import (
	"github.com/nmarsollier/resourcesgo/tools/db"
	"github.com/nmarsollier/resourcesgo/tools/logx"
)

func FindByID(logenv logx.Fields, id string) (project *Resource, err error) {
	project, err = db.QueryRow[Resource](
		logenv,
		`
		SELECT id, project, language, sem_ver, values, created, enabled 
		FROM resources 
		WHERE id = $1
		`,
		id,
	)

	return
}
