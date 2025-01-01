package terr

import (
	"strconv"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/nmarsollier/resourcesgo/tools/db"
)

var PgErrorExist = &pgconn.PgError{
	Code: strconv.Itoa(db.ERR_EXIST),
}

var PgErrorForeign = &pgconn.PgError{
	Code: strconv.Itoa(db.ERR_FOREIGN_KEY),
}
