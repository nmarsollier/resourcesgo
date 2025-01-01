package projects

import (
	"github.com/nmarsollier/resourcesgo/tools/db"
	"github.com/nmarsollier/resourcesgo/tools/errs"
	"github.com/nmarsollier/resourcesgo/tools/logx"
)

func Create(logenv logx.Fields, id string, name string) (string, error) {
	project := newProject(id, name)

	if err := project.ValidateSchema(); err != nil {
		return "", err
	}

	err := db.Exec(
		logenv,
		"INSERT INTO projects (id, name, created, enabled) VALUES ($1, $2, $3, $4)",
		project.ID, project.Name, project.Created, project.Enabled)

	if err != nil {
		switch db.ErrorCode(err) {
		case db.ERR_EXIST:
			return "", errs.AlreadyExist
		}

		return "", err
	}

	return project.ID, nil
}
