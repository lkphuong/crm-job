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
)
