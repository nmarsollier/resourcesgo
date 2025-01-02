package languages

import (
	"context"

	"github.com/nmarsollier/resourcesgo/internal/tools/db"
	"github.com/nmarsollier/resourcesgo/internal/tools/errs"
)

var dbExec = db.Exec

func Create(ctx context.Context, id string, name string) (string, error) {
	language := newLanguage(id, name)

	if err := language.ValidateSchema(); err != nil {
		return "", err
	}

	err := dbExec(
		ctx,
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
