package earning_point

type SaleReceiptInfo struct {
	ReceiptNumber       string `boil:"receipt_number" json:"receipt_number"`
	CustomerCode        string `boil:"customer_code" json:"customer_code"`
	MembershipLevelCode string `boil:"membership_level_code" json:"membership_level_code"`
	CrmAddPoint         bool   `boil:"crm_add_point" json:"crm_add_point"`
}

type BillEarningPointResponse struct {
	ID        string `boil:"id" json:"id"`
	CuahangID string `boil:"cuahang_id" json:"cuahang_id"`
}

type PointResponse struct {
	Point int `boil:"point" json:"point"`
}

type ExpiredPointResponse struct {
	TransactionNumber string `boil:"transaction_number" json:"transaction_number"`
	CustomerCode      string `boil:"customer_code" json:"customer_code"`
	AvalaibleValue    string `boil:"avalaible_value" json:"avalaible_value"`
}

type CurrentPointResponse struct {
	CustomerCode string  `boil:"customer_code" json:"customer_code"`
	RemainPoints float64 `boil:"remain_points" json:"remain_points"`
	TotalPoints  float64 `boil:"total_points" json:"total_points"`
}

type EarningPointExpired struct {
	EarningPointHistoryId int64   `json:"earning_point_history_id" boil:"earning_point_history_id"`
	TransactionNumber     string  `json:"transaction_number" boil:"transaction_number"`
	CustomerCode          string  `json:"customer_code" boil:"customer_code"`
	StoreCode             string  `json:"store_code" boil:"store_code"`
	AvalaibleValue        float64 `json:"avalaible_value" boil:"avalaible_value"`
}
