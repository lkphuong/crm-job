package customer

import (
	"context"
	"database/sql"
)

var (
	_repository CustomerRepository
)

func init() {
	_repository = CustomerRepository{}
}

func UpdateCustomerDuplicate(ctx context.Context, db *sql.DB) error {

	customers, err := _repository.GetCustomerDuplicate(ctx, db)

	if err != nil {
		return err
	}

	if len(customers) > 0 {
		return _repository.UpdateCustomerDuplicate(ctx, db)
	}

	return nil
}
