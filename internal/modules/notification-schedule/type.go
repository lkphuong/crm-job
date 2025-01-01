package notification_schedule

type NotificationCampaign struct {
	ID      int64  `boil:"notification_campaign_id" json:"notification_campaign_id"`
	Title   string `boil:"title" json:"title"`
	Summary string `boil:"summary" json:"summary"`
}

type FirebaseToken struct {
	FirebaseToken string `boil:"firebase_token" json:"firebase_token"`
}

type NotificationSchedule struct {
	CustomerCode        string `boil:"customer_code" json:"customer_code"`
	NotificationContent string `boil:"notification_content" json:"notification_content"`
}
