package customer

type CustomerResponse struct {
	CustomerCode string `json:"customer_code"`
}

type JobCustomerResponse struct {
	CustomerCode string `boil:"customer_code" json:"customer_code"`
	Job          string `boil:"job" json:"job"`
	UpdatedAt    string `boil:"updated_at" json:"updated_at"`
}
