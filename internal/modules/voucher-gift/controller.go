package voucher_gift

import (
	"context"
	"database/sql"

	"github.com/robfig/cron"
)

func UpdateVoucher(ctx context.Context, db *sql.DB) error {
	job := cron.New()

	job.AddFunc("0 0 0 * * *", func() {
		err := UpdateVoucherGiftExpire(ctx, db)

		if err != nil {
			return
		}
	})

	return nil
}
