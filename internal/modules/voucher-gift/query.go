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

	UPDATE_VOUCHER_GIFT_USED = `
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

	GET_VOUCHER_GIFT = `
		SELECT
			TOP 1 voucher_gift_code
		FROM
			voucher_gift_tbl
		WHERE
			voucher_gift_code = '%s'
	`

	GET_SALE_CODE_PUBLIC = `
		SELECT
			TOP 1 code, sale_id
		FROM
			SaleCodePublic
		WHERE 
			trangthai = 4
		ORDER BY created_date DESC
	`

	GET_COUPON = `
		SELECT
			TOP 1 Coupon.code,
			Coupon.sale_id,
			Sale.name,
			Coupon.[begin],
			Coupon.[end],
			Sale.code as ref_sku
		FROM
			Coupon JOIN Sale ON Coupon.sale_id = Sale.id
		WHERE
			Sale.trangthai = 4 AND Coupon.sale_id = '%s'
	`

	INSERT_VOUCHER_GIFT = `
		INSERT INTO voucher_gift_tbl (
			voucher_gift_code, 
			voucher_gift_name, 
			customer_name, 
			customer_phone, 
			valid_from, 
			valid_to, 
			status, 
			ref_sku
		)
		SELECT
			code,
			N'%s',
			customer_name,
			tel,
			'%s',
			'%s',
			0,
			'%s'
		FROM Coupon
		WHERE sale_id = '%s'
	`

	UPDATE_EXPIRED_VOUCHER_GIFT = `
		UPDATE voucher_gift_tbl set
			valid_to = '%s'
		WHERE
			valid_to = '%s'
	`

	GET_VOUCHER_DUPLICATE = `
		SELECT
			customer_code,
			ref_sku,
			count(ref_sku) AS count
		FROM
			voucher_gift_tbl
		WHERE
			year(valid_to) = year(getdate ())
			AND ref_sku IN ('CRM03','CRM02')
			AND delete_flag = 0
			AND valid_to > getdate ()
			AND customer_code IS NOT NULL
		GROUP BY
			customer_code,
			ref_sku
		HAVING
			count(ref_sku) > 1
	`
)
