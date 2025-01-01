package languages

import (
	"github.com/nmarsollier/resourcesgo/tools/db"
	"github.com/nmarsollier/resourcesgo/tools/errs"
	"github.com/nmarsollier/resourcesgo/tools/logx"
)

var dbExec = db.Exec

func Create(fields logx.Fields, id string, name string) (string, error) {
	language := newLanguage(id, name)

	if err := language.ValidateSchema(); err != nil {
		return "", err
	}

	err := dbExec(
		fields,
		"INSERT INTO languages (id, name, created, enabled) VALUES ($1, $2, $3, $4)",
		language.ID,
		language.Name,
		language.Created,
		language.Enabled,
	)

	if err != nil {
		switch db.ErrorCode(err) {
		case db.ERR_EXIST:
			return "", errs.AlreadyExist
		}

		return "", err
	}

	return language.ID, nil
}
