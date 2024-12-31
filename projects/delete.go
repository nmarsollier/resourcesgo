package projects

import (
	"github.com/nmarsollier/resourcesgo/tools/db"
	"github.com/nmarsollier/resourcesgo/tools/logx"
)

func Delete(logenv logx.Fields, id string) error {
	return db.Exec(
		logenv,
		"UPDATE projects SET enabled=false WHERE id=$1",
		id,
	)
}
