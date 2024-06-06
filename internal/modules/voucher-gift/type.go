package voucher_gift

type VoucherGift struct {
	VoucherGiftID int64 `json:"voucher_gift_id"`
}

type VoucherBirthday struct {
	CustomerCode string `json:"customer_code"`
	Count        int    `json:"count"`
}
