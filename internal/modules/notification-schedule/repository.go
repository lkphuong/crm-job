package notification_schedule

import (
	"context"
	"database/sql"
)

type Repository struct{}

func (r *Repository) DeleteNotificationDraft(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, DELETE_NOTIFICATION_SCHEDULE)

	if err != nil {
		return err
	}

	return nil
}
