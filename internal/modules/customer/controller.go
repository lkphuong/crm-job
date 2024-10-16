package customer

import (
	"context"

	"github.com/robfig/cron"
)

func UpdateCustomer(ctx context.Context) error {
	job := cron.New()

	job.AddFunc("0 0 10 * * *", func() {
		err := UpdateCustomerDuplicate(ctx)

		if err != nil {
			return
		}
	})

	job.Start()

	return nil
}

func UpdateJob(ctx context.Context) error {
	job := cron.New()

	job.AddFunc("0 0 7 * * *", func() {
		err := UpdateJobFromCustomerToKhanhHang(ctx)

		if err != nil {
			return
		}
	})

	job.Start()

	return nil
}
