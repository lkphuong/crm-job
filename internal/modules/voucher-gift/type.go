package voucher_gift

type VoucherGift struct {
	VoucherGiftID int64 `json:"voucher_gift_id"`
}

type VoucherBirthday struct {
	CustomerCode string `json:"customer_code"`
	Count        int    `json:"count"`
}

type VoucherGiftCode struct {
	VoucherGiftCode string `boil:"voucher_gift_code" json:"voucher_gift_code"`
}

type SalePublicCode struct {
	Code   string `boil:"code" json:"code"`
	SaleID string `boil:"sale_id" json:"sale_id"`
}

type Coupon struct {
	Code   string `boil:"code" json:"code"`
	SaleID string `boil:"sale_id" json:"sale_id"`
	Name   string `boil:"name" json:"name"`
	Begin  string `boil:"begin" json:"begin"`
	End    string `boil:"end" json:"end"`
	RefSku string `boil:"ref_sku" json:"ref_sku"`
}
