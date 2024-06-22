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

	job.AddFunc("* * * * * *", func() {
		service.EaringPointHoangDieu2(ctx)
	})

	job.Start()
}

func EaringPoint(ctx context.Context) {
	job := cron.New()

	job.AddFunc("* * * * * *", func() {
		service.EaringPoint(ctx)
	})

	job.Start()
}
