package customer

import (
	"context"
	"fmt"

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

func (c *Repository) GetJobFromCustomer(ctx context.Context) ([]JobCustomerResponse, error) {
	var jobs []JobCustomerResponse

	err := queries.Raw(GET_JOB_FROM_CUSTOMER).Bind(ctx, db, &jobs)

	if err != nil {
		return nil, err
	}

	return jobs, nil
}

func (c *Repository) GetJobFromPOS(ctx context.Context) ([]JobCustomerResponse, error) {
	var jobs []JobCustomerResponse

	err := queries.Raw(GET_JOB_FROM_POS).Bind(ctx, db, &jobs)

	if err != nil {
		return nil, err
	}

	return jobs, nil
}

func (c *Repository) GetCustomerPOSByCode(ctx context.Context, customerCode string) (*JobCustomerResponse, error) {
	var jobs JobCustomerResponse

	err := queries.Raw(fmt.Sprintf(GET_CUSTOMER_POS_BY_CODE, customerCode)).Bind(ctx, db, &jobs)

	if err != nil {
		return nil, err
	}

	return &jobs, nil
}

func (c *Repository) GetCustomerCRMByCode(ctx context.Context, customerCode string) (*JobCustomerResponse, error) {
	var jobs JobCustomerResponse

	err := queries.Raw(fmt.Sprintf(GET_CUSTOMER_CRM_BY_CODE, customerCode)).Bind(ctx, db, &jobs)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	return &jobs, nil
}

func (c *Repository) UpdateJobOfCRM(ctx context.Context, customerCode string, job string) error {
	_, err := db.ExecContext(ctx, fmt.Sprintf(UPDATE_JOB_OF_CRM, job, customerCode))

	if err != nil {
		return err
	}

	return nil
}

func (c *Repository) UpdateJobOfPOS(ctx context.Context, customerCode string, job string) error {
	_, err := db.ExecContext(ctx, fmt.Sprintf(UPDATE_JOB_OF_POS, job, customerCode))

	if err != nil {
		return err
	}

	return nil
}
