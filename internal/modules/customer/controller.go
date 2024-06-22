package customer

import (
	"context"
	"fmt"

	"github.com/robfig/cron"
)

func UpdateCustomer(ctx context.Context) error {
	job := cron.New()

	fmt.Println("Update customer duplicate")

	job.AddFunc("0 0 * * * *", func() {
		err := UpdateCustomerDuplicate(ctx)

		fmt.Println("clear customer duplicate: ", err)

		if err != nil {
			return
		}
	})

	job.Start()

	return nil
}
