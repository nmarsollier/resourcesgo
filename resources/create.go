package resources

import (
	"github.com/go-playground/validator/v10"
	"github.com/nmarsollier/resourcesgo/tools/db"
	"github.com/nmarsollier/resourcesgo/tools/errs"
	"github.com/nmarsollier/resourcesgo/tools/logx"
)

var dbExec = db.Exec

func Create(
	fields logx.Fields,
	resource *Resource,
) (string, error) {
	if err := validator.New().Struct(resource); err != nil {
		return "", err
	}

	resource, err := insert(fields, resource)

	if err != nil {
		switch db.ErrorCode(err) {
		case db.ERR_FOREIGN_KEY:
			return "", errs.ErrProjectNotExist
		case db.ERR_EXIST:
			return "", errs.AlreadyExist
		}
		return "", err
	}

	return resource.ID, nil
}

func insert(
	fields logx.Fields,
	resource *Resource,
) (*Resource, error) {
	err := dbExec(fields,
		`
		INSERT INTO resources (id, project, language, sem_ver, values, created, enabled)
		VALUES ($1, $2, $3, $4, $5, $6, $7);
		`,
		resource.ID, resource.ProjectID, resource.LanguageID, resource.SemVer, resource.Values, resource.Created, resource.Enabled)

	if err != nil {
		return nil, err
	}
	return resource, nil
}
