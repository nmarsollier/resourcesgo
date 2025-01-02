package resources

import (
	"github.com/nmarsollier/resourcesgo/internal/tools/db"
	"github.com/nmarsollier/resourcesgo/internal/tools/logx"
)

func FindVersions(fields logx.Fields, project string, language string) ([]string, error) {
	data, err := db.Query[string](
		fields,
		"SELECT sem_ver FROM resources WHERE project = $1 AND language = $2 AND enabled = true",
		project, language,
	)

	if err != nil {
		return make([]string, 0), err
	}

	versions := make([]string, len(data))
	for i, v := range data {
		versions[i] = *v
	}

	return versions, nil

}
