package voucher_gift

import (
	"context"

	"github.com/robfig/cron"
)

func InsertVoucherGiftSalePublicCode(ctx context.Context) error {
	job := cron.New()

	job.AddFunc("0 */10 * * * *", func() {
		err := InsertVoucherGift(ctx)

		if err != nil {
			return
		}
	})

	job.Start()

	return nil

}

func VoucherGiftExpire(ctx context.Context) error {
	job := cron.New()

	job.AddFunc("0 0 * * * *", func() {
		err := UpdateVoucherGiftExpire(ctx)

		if err != nil {
			return
		}
	})

	job.Start()

	return nil
}

func voucherGiftUsed(ctx context.Context) error {
	job := cron.New()

	job.AddFunc("0 0 */4 * * *", func() {
		err := UpdateVoucherUsed(ctx)

		if err != nil {
			return
		}
	})

	job.Start()

	return nil
}
