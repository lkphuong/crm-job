package voucher_gift

const (
	GET_VOUCHER_GIFT_EXPIRE = `
		SELECT
			voucher_gift_id
		FROM
			voucher_gift_tbl
		WHERE
			valid_to > getdate()
			AND status = 2
			AND used_customer_code IS NULL
	`

	UPDATE_VOUCHER_GIFT_EXPIRE = `
		UPDATE
			voucher_gift_tbl
		SET
			status = 1
		WHERE
			voucher_gift_id IN (
				SELECT voucher_gift_id
				FROM voucher_gift_tbl
				WHERE valid_to > getdate()
					AND status = 2
					AND used_customer_code IS NULL
			)
	`

	GET_VOUCHER_BIRTHDAY_DUPLICATE = `
		SELECT
			customer_code,
			COUNT(customer_code) AS count
		FROM
			 voucher_gift_tbl
		WHERE
			ref_sku IN ('CRM-03', 'CRM-14')
		GROUP BY
			customer_code
		HAVING
			COUNT(customer_code) > 1
	`
)
