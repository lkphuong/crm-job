package notification_schedule

import (
	"context"
	"fmt"

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

func PushNotificationCampaign(ctx context.Context) {
	job := cron.New()

	job.AddFunc("0 * * * * *", func() {
		fmt.Println("health check")
		service.PushNotificationCampaign(ctx, db)
	})

	job.Start()

}
