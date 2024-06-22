package voucher_gift

import (
	"context"

	"github.com/robfig/cron"
)

func InsertVoucherGiftSalePublicCode(ctx context.Context) error {
	job := cron.New()

	job.AddFunc("* * * * * *", func() {
		err := InsertVoucherGift(ctx, db)

		if err != nil {
			return
		}
	})

	job.Start()

	return nil

}

func VoucherGiftExpire(ctx context.Context) error {
	job := cron.New()

	job.AddFunc("0 * * * * *", func() {
		err := UpdateVoucherGiftExpire(ctx, db)

		if err != nil {
			return
		}
	})

	job.Start()

	return nil
}
