package earning_point

const (
	GET_BILL_HOANG_DIEU2 = `
			SELECT
				HoaDon.id, hoadon.cuahang_id
			FROM
				HoaDon
				JOIN Ca ON Ca.id = HoaDon.ca_id
				LEFT JOIN HoaDon_Earning_Point ON HoaDon.id = HoaDon_Earning_Point.bill_id
				AND HoaDon_Earning_Point.cuahang_id = '31'
			WHERE
				HoaDon.cuahang_id = '31'
				AND HoaDon_Earning_Point.bill_id IS NULL
				AND Ca.[date] >= '2024-03-01'
				AND HoaDon.customer_code IS NOT NULL
				AND hoadon.customer_code <> '' 
				AND hoadon.finish IS NOT NULL
	`

	GET_SALE_RECEIPT_INFO = `
		 	select 
				s.receipt_number, 
				s.customer_code, 
				m.membership_level_code, 
				s.total, s.crm_add_point, 
				e.[value], 
				e.avalaible_value
			from 
				sale_receipt_tbl s 
				left join earning_point_history_tbl e on e.ref_doc_number = s.receipt_number 
				left join membership_history_tbl m on m.customer_code = s.customer_code
			where 
				s.created_at > '2024-01-01 00:00:00' and crm_add_point = 1 
				and (m.status = 1 or m.status is null) and [value] is null
	`

	INSERT_EARNING_POINT_HISTORY_HOANGDIEU2 = `
		INSERT INTO [HoaDon_Earning_Point]
          (
            bill_id,
            new_flag,
            cuahang_id
          )
        VALUES
          (
            '%s',
            1,
            '%s'
          )
	`

	INSERT_EARNING_POINT_HISTORY = `
				INSERT INTO [earning_point_history_tbl]
                    ( [transaction_number]
                    , [customer_code]
                    , [created_date]
                    , [activated_date]
                    , [expired_date]
                    , [store_code]
                    , [earning_type]
                    , [earning_status]
                    , [value]
                    , [avalaible_value]
                    , [ref_doc_number]
                    , [ref_transaction_number]
                    , [amount]
                    , [product_brand]
                    , [level_code]
                    , [update_date]
                    , [ref_store_name])
                SELECT
                    NEWID()
                    , h.customer_code
                    , CAST(h.created_at AS datetime)
                    , CAST(h.created_at AS datetime)
                    , dbo.GetExpiredDateCrm(h.created_at,360)
                    , h.store_code
                    , 0
                    , 2
                    , ROUND(h.total/m.point_rule_amount,0,1)
                    , ROUND(h.total/m.point_rule_amount,0,1)
                    , h.receipt_number
                    , ''
                    , h.total
                    , ''
                    , '%s'
                    , CURRENT_TIMESTAMP
                    , i.inventory_name
                FROM sale_receipt_tbl h 
                    JOIN customer_tbl c 
                        ON h.customer_code = c.customer_code AND 
                            c.loyalty_flag = 1 and c.delete_flag = 0
                    JOIN membership_tbl m 
                        ON m.membership_code = 'ICOOL'
                    LEFT JOIN inventory_tbl i
                        ON i.inventory_code = h.store_code
                WHERE h.receipt_number = '%s';
	`

	INSERT_EARNING_POINT_HISTORY_RANK_DIAMOND = `
				INSERT INTO [earning_point_history_tbl]
                    ( [transaction_number]
                    , [customer_code]
                    , [created_date]
                    , [activated_date]
                    , [expired_date]
                    , [store_code]
                    , [earning_type]
                    , [earning_status]
                    , [value]
                    , [avalaible_value]
                    , [ref_doc_number]
                    , [ref_transaction_number]
                    , [amount]
                    , [product_brand]
                    , [level_code]
                    , [update_date]
                    , [ref_store_name])
                SELECT
                    NEWID()
                    , h.customer_code
                    , CAST(h.created_at AS datetime)
                    , CAST(h.created_at AS datetime)
                    , dbo.GetExpiredDateCrm(h.created_at,360)
                    , h.store_code
                    , 0
                    , 2
                    , ROUND(h.total/m.point_rule_amount * 0.1,0,1)
                    , ROUND(h.total/m.point_rule_amount * 0.1,0,1)
                    , h.receipt_number
                    , ''
                    , h.total
                    , ''
                    , '%s'
                    , CURRENT_TIMESTAMP
                    , i.inventory_name
                FROM sale_receipt_tbl h 
                    JOIN customer_tbl c 
                        ON h.customer_code = c.customer_code AND 
                            c.loyalty_flag = 1 and c.delete_flag = 0
                    JOIN membership_tbl m 
                        ON m.membership_code = 'ICOOL'
                    LEFT JOIN inventory_tbl i
                        ON i.inventory_code = h.store_code
                WHERE h.receipt_number = '%s';
	`

	CALCULATOR_POINT_RANK_DAIMOND = `
			SELECT
                CAST(ROUND(h.total/m.point_rule_amount,0,1) + ROUND(h.total/m.point_rule_amount * 0.1,0,1) as INT) as point
            FROM sale_receipt_tbl h JOIN customer_tbl c on h.customer_code = c.customer_code AND c.loyalty_flag = 1 and c.delete_flag=0
            JOIN membership_tbl m ON m.membership_code ='ICOOL'
            WHERE h.receipt_number = '%s';
	`

	CALCULATOR_POINT = `
			SELECT
               CAST(ROUND(h.total/m.point_rule_amount,0,1) as INT) as point
            FROM sale_receipt_tbl h JOIN customer_tbl c on h.customer_code = c.customer_code AND c.loyalty_flag = 1 and c.delete_flag=0
            JOIN membership_tbl m ON m.membership_code ='ICOOL'
            WHERE h.receipt_number = '%s';
	`

	SEND_NOTIFICATION = `
			EXEC dbo.SendNotification '%s', 'ADD_POINT', 'NEW_BILL', '/histories?type=point', '%s', null;
	`

	ADD_LOYALTY_FIRST_BILL = `
			EXEC dbo.AddLoyaltyFirstBill '%s', '%s';
	`

	ADD_REFERRAL_REWARD = `
		EXEC dbo.AddReferralReward '%s', '%s';
	`

	SELECT_EXPIRED_30DAYS = `
        SELECT
            earning_point_history_tbl.transaction_number,
            earning_point_history_tbl.customer_code,
            earning_point_history_tbl.avalaible_value
        FROM
            earning_point_history_tbl
            JOIN customer_tbl ON earning_point_history_tbl.customer_code = customer_tbl.customer_code
        WHERE
            NOT EXISTS (
                SELECT
                    almost_expired_points_tbl.transaction_number
                FROM
                    almost_expired_points_tbl
                WHERE
                    almost_expired_points_tbl.transaction_number = earning_point_history_tbl.transaction_number
                    AND almost_expired_points_tbl.expired_category = 0)
                AND(CAST(earning_point_history_tbl.expired_date AS DATE) >= CONVERT(varchar, GETDATE (), 23) 
                    AND CAST(earning_point_history_tbl.expired_date AS DATE) = CONVERT(varchar, DATEADD(day, 29, GETDATE()), 23))
                AND earning_point_history_tbl.earning_status = 2
                AND earning_point_history_tbl.avalaible_value > 0
                AND earning_point_history_tbl.delete_flag = 0
    `

	INSERT_ALMOST_EXPIRED_POINTS = `
                INSERT INTO almost_expired_points_tbl (transaction_number, customer_code, points_value, expired_category, new_flag)
		        values('%s', '%s', %s, 0, 1)
    `
)
