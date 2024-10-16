package notification_schedule

const (
	DELETE_NOTIFICATION_SCHEDULE = `
			DELETE notification_schedule_tbl
			WHERE id in(
					SELECT
						TOP 1000 id FROM notification_schedule_tbl
					WHERE
						created_date < DATEADD (day, - 30, GETDATE ()))
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

	GET_NOTIFICATION_CAMPAIGN = `
			SELECT
				nc.notification_campaign_id,
				c.title,
				c.summary
			FROM
				notification_campaign_tbl nc
				LEFT JOIN notification_content_tbl c ON nc.content_id = c.notification_content_id
			WHERE
				nc.status = 1
				AND nc.delete_flag = 0
				AND nc.planned_send_date <= DATEADD(MINUTE, 1, GETDATE())
				AND planned_send_date > '2024-08-18 12:00:00.000'

	`

	GET_FIREBASE_TOKENS = `
			SELECT
				firebase_token
			FROM
				account_tokens_tbl att
			WHERE
				logged = 1
			ORDER BY
				account_token_id ASC
	`

	UPDATE_NOTIFICATION_CAMPAIGN = `
			UPDATE notification_campaign_tbl set status = 2 where notification_campaign_id = %d
	`
)
