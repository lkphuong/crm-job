package customer

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/queries"
)

type CustomerRepository struct{}

func (c *CustomerRepository) GetCustomerDuplicate(ctx context.Context, db *sql.DB) ([]CustomerResponse, error) {
	var customers []CustomerResponse

	err := queries.Raw(GET_CUSTOMER_DUPLICATE).Bind(ctx, db, &customers)

	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (c *CustomerRepository) UpdateCustomerDuplicate(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, UPDATE_CUSTOMER_DUPLICATE)

	if err != nil {
		return err
	}

	return nil
}
