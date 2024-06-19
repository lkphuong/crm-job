package notification_schedule

import (
	"context"

	"github.com/robfig/cron"
)

var (
	service *Service
)

func DeleteNotificationDraft(ctx context.Context) {
	job := cron.New()

	job.AddFunc("0 * * * * *", func() {
		service.DeleteNotificationDraft(ctx, db)
	})

	job.Start()
}
