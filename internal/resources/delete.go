package resources

import (
	"github.com/nmarsollier/resourcesgo/internal/tools/db"
	"github.com/nmarsollier/resourcesgo/internal/tools/logx"
)

func Delete(fields logx.Fields, project string, language string, semver string) {
	if resource, err := findBy(fields, project, language, semver); err == nil {
		db.Exec(
			fields,
			"UPDATE resources SET enabled=false WHERE id=$1",
			resource.ID,
		)
	}
}
