package customer

import (
	"context"
)

var (
	_repository Repository
)

func UpdateCustomerDuplicate(ctx context.Context) error {

	customers, err := _repository.GetCustomerDuplicate(ctx)

	if err != nil {
		return err
	}

	if len(customers) > 0 {
		return _repository.UpdateCustomerDuplicate(ctx)
	}

	return nil
}
