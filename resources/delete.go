package resources

import (
	"github.com/nmarsollier/resourcesgo/tools/db"
	"github.com/nmarsollier/resourcesgo/tools/logx"
)

func Delete(logenv logx.Fields, proejct string, language string, semver string) {
	if resource, err := findBy(logenv, proejct, language, semver); err == nil {
		db.Exec(
			logenv,
			"UPDATE resources SET enabled=false WHERE id=$1",
			resource.ID,
		)
	}
}
