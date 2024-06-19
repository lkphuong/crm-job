package notification_schedule

import (
	"context"
	"database/sql"
	"fmt"
)

var (
	repository Repository
)

type Service struct{}

func (s *Service) DeleteNotificationDraft(ctx context.Context, db *sql.DB) error {

	err := repository.DeleteNotificationDraft(ctx, db)

	fmt.Println("job clear draft notification: ", err)

	if err != nil {
		return err
	}

	return nil
}
