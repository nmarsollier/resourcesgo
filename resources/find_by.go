package resources

import (
	"github.com/nmarsollier/resourcesgo/tools/db"
	"github.com/nmarsollier/resourcesgo/tools/logx"
)

func findBy(logenv logx.Fields, project string, language string, semVer string) (*Resource, error) {
	return db.QueryRow[Resource](
		logenv,
		`
		SELECT id, project, language, sem_ver, values, created, enabled 
		FROM resources 
		WHERE project = $1 AND language = $2 AND sem_ver = $3
		`,
		project, language, semVer,
	)
}
