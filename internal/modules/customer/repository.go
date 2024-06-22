package customer

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/queries"
)

type Repository struct{}

func (c *Repository) GetCustomerDuplicate(ctx context.Context) ([]CustomerResponse, error) {
	var customers []CustomerResponse

	err := queries.Raw(GET_CUSTOMER_DUPLICATE).Bind(ctx, db, &customers)

	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (c *Repository) UpdateCustomerDuplicate(ctx context.Context) error {
	_, err := db.ExecContext(ctx, UPDATE_CUSTOMER_DUPLICATE)

	if err != nil {
		return err
	}

	return nil
}
