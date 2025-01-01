package resources

import (
	"github.com/nmarsollier/resourcesgo/tools/db"
	"github.com/nmarsollier/resourcesgo/tools/logx"
)

func FindByID(fields logx.Fields, id string) (*Resource, error) {
	return db.QueryRow[Resource](
		fields,
		`
		SELECT id, project, language, sem_ver, values, created, enabled 
		FROM resources 
		WHERE id = $1
		`,
		id,
	)
}
