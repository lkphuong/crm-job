package modules

import (
	"database/sql"

	"github.com/volatiletech/sqlboiler/boil"

	sqlserver "github.com/lkphuong/crm-job/configs/database"
)

func GetDB() *sql.DB {
	db := sqlserver.ConnectionSqlServcer()

	boil.SetDB(db)

	return db
}
