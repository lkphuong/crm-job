package notification_schedule

const (
	DELETE_NOTIFICATION_SCHEDULE = `
			DELETE notification_schedule_tbl
			WHERE id in(
					SELECT
						TOP 1000 id FROM notification_schedule_tbl
					WHERE
						id NOT IN(
							SELECT
								notification_schedule_tbl.id FROM account_tokens_tbl
								JOIN customer_tbl ON account_tokens_tbl.account_id = customer_tbl.account_id
								JOIN notification_schedule_tbl ON customer_tbl.customer_code = notification_schedule_tbl.customer_code
							WHERE
								account_tokens_tbl. [logged] = 1) or created_date < DATEADD(day, -30, GETDATE()))
	`

	INSERT_NOTIFICATION_SCHEDULE = `
			SELECT
				notification_schedule_tbl.id
			FROM
				account_tokens_tbl
				JOIN customer_tbl ON account_tokens_tbl.account_id = customer_tbl.account_id
				JOIN notification_schedule_tbl ON customer_tbl.customer_code = notification_schedule_tbl.customer_code
			WHERE
				account_tokens_tbl. [logged] = 1
				AND notification_schedule_tbl.created_date > DATEADD (day, - 30, GETDATE ())
	`
)
