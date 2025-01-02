package terr

import (
	"strconv"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/nmarsollier/resourcesgo/internal/tools/db"
)

var PgErrorExist = &pgconn.PgError{
	Code: strconv.Itoa(db.ERR_EXIST),
}

var PgErrorLanguageForeign = &pgconn.PgError{
	Code:    strconv.Itoa(db.ERR_FOREIGN_KEY),
	Message: "language",
}

var PgErrorProjectForeign = &pgconn.PgError{
	Code:    strconv.Itoa(db.ERR_FOREIGN_KEY),
	Message: "project",
}
