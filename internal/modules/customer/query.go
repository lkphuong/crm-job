package customer

const (
	GET_CUSTOMER_DUPLICATE = `
				SELECT customer_code
				FROM customer_tbl
				WHERE delete_flag = 0
				GROUP BY customer_code
				HAVING COUNT(customer_code) > 1
			`

	UPDATE_CUSTOMER_DUPLICATE = `
				UPDATE customer_tbl
				SET delete_flag = 1
				WHERE customer_code IN (
					SELECT customer_code
					FROM customer_tbl
					WHERE delete_flag = 0
					GROUP BY customer_code
					HAVING COUNT(customer_code) > 1
				) AND account_id IS NULL
			`

	GET_JOB_FROM_CUSTOMER = `
				SELECT
					top 1000 customer_code,
					job,
					update_date AS updated_at
				FROM
					customer_tbl
				WHERE
					job IS NOT NULL
					AND update_date > CONVERT(varchar, DATEADD(day, -1, GETDATE()), 23)
					AND update_date < CONVERT(varchar, DATEADD(day, 1, GETDATE()), 23)
					AND delete_flag = 0
	`

	GET_JOB_FROM_POS = `
				SELECT
					top 1000 code AS customer_code,
					job,
					modified_date AS updated_at
				FROM
					khachhang
				WHERE
					job <> '0'
					AND job <> ''
					AND modified_date > CONVERT(varchar, DATEADD(day, -1, GETDATE()), 23)
					AND modified_date < CONVERT(varchar, DATEADD(day, 1, GETDATE()), 23)
					AND trangthai = 4
	`

	GET_CUSTOMER_POS_BY_CODE = `
				SELECT
					code AS customer_code,
					job,
					modified_date AS updated_at
				FROM
					khachhang
				WHERE
					code = '%s'
	`

	GET_CUSTOMER_CRM_BY_CODE = `
				SELECT
					customer_code,
					CASE WHEN job IS NULL THEN
						'none'
					ELSE
						job
					END AS job,
					update_date AS updated_at
				FROM
					customer_tbl
				WHERE
					customer_code = '%s'
	`

	UPDATE_JOB_OF_CRM = `UPDATE customer_tbl SET job = '%s' WHERE customer_code = '%s'`

	UPDATE_JOB_OF_POS = `UPDATE khachhang SET job = '%s' WHERE code = '%s'`
)
