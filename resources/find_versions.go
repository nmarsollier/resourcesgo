package resources

import (
	"github.com/nmarsollier/resourcesgo/tools/db"
	"github.com/nmarsollier/resourcesgo/tools/logx"
)

func FindVersions(logenv logx.Fields, project string, language string) ([]*string, error) {
	return db.Query[string](
		logenv,
		"SELECT sem_ver FROM resources WHERE project = $1 AND language = $2 AND enabled = true",
		project, language,
	)
}
