package earning_point

import (
	"context"

	"github.com/robfig/cron"
)

var (
	service *Service
)

func EarningPointBillHoangDieu2(ctx context.Context) {
	job := cron.New()

	job.AddFunc("0 */10 * * * *", func() {
		service.EaringPointHoangDieu2(ctx)
	})

	job.Start()
}

func EaringPoint(ctx context.Context) {
	job := cron.New()

	job.AddFunc("0 */5 * * * *", func() {
		service.EaringPoint(ctx)
	})

	job.Start()
}

func ExpiredPoint30Days(ctx context.Context) {
	job := cron.New()

	job.AddFunc("0 0 10 * * *", func() {
		service.ExpiredPoint(ctx)
	})

	job.Start()
}
