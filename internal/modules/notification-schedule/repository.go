package notification_schedule

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/volatiletech/sqlboiler/queries"
)

type Repository struct{}

func (r *Repository) DeleteNotificationDraft(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, DELETE_NOTIFICATION_SCHEDULE)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetNotificationCampaign(ctx context.Context) ([]NotificationCampaign, error) {
	var notificationCampaign []NotificationCampaign

	err := queries.Raw(GET_NOTIFICATION_CAMPAIGN).Bind(ctx, db, &notificationCampaign)

	if err != nil {
		return nil, err
	}

	return notificationCampaign, nil
}

func (r *Repository) GetFirebaseTokens(ctx context.Context) ([]FirebaseToken, error) {
	var firebaseTokens []FirebaseToken

	err := queries.Raw(GET_FIREBASE_TOKENS).Bind(ctx, db, &firebaseTokens)

	if err != nil {
		return nil, err
	}

	return firebaseTokens, nil
}

func (r *Repository) InsertNotificationSchedule(ctx context.Context, db *sql.DB, customerCode string, content string) error {
	_, err := db.ExecContext(ctx, fmt.Sprintf(INSERT_NOTIFICATION_SCHEDULE, customerCode, content))

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateNotificationCampaign(ctx context.Context, db *sql.DB, notificationCampaignID int64) error {
	_, err := db.ExecContext(ctx, fmt.Sprintf(UPDATE_NOTIFICATION_CAMPAIGN, notificationCampaignID))

	if err != nil {
		return err
	}

	return nil
}
