package notification_schedule

import (
	"context"
	"database/sql"
)

var (
	repository Repository
)

type Service struct{}

func (s *Service) DeleteNotificationDraft(ctx context.Context, db *sql.DB) error {

	err := repository.DeleteNotificationDraft(ctx, db)

	if err != nil {
		return err
	}

	return nil
}
