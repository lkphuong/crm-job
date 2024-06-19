package customer

import (
	"database/sql"

	"github.com/lkphuong/crm-job/internal/modules"
)

var (
	db *sql.DB
)

func init() {
	db = modules.GetDB()
}
