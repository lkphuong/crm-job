package customer

import (
	"context"
	"database/sql"

	"github.com/robfig/cron"
)

func UpdateCustomer(ctx context.Context, db *sql.DB) error {
	job := cron.New()

	job.AddFunc("0 0 0 * * *", func() {
		err := UpdateCustomerDuplicate(ctx, db)

		if err != nil {
			return
		}
	})

	return nil
}
