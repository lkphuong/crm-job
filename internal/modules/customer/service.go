package customer

import (
	"context"
	"fmt"
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

func UpdateJobFromCustomerToKhanhHang(ctx context.Context) error {

	jobPos, _ := _repository.GetJobFromPOS(ctx)
	jobCrm, _ := _repository.GetJobFromCustomer(ctx)

	fmt.Println(len(jobPos))
	fmt.Println(len(jobCrm))

	//  cập nhật thông tin nghề nghiệp từ bảng customer sang bảng khách hàng
	// for _, job := range jobCrm {
	// 	if job.CustomerCode != "" {
	// 		customer, _ := _repository.GetCustomerPOSByCode(ctx, job.CustomerCode)
	// 		if customer != nil {
	// 			if (customer.UpdatedAt <= job.UpdatedAt && customer.Job != job.Job) || customer.Job == "" {

	// 				err := _repository.UpdateJobOfPOS(ctx, job.CustomerCode, job.Job)

	// 				if err != nil {
	// 					fmt.Println("Cập nhật thông tin nghề nghiệp từ bảng customer sang bảng khách hàng thất bại")
	// 				}

	// 			}
	// 		}
	// 	}
	// }

	//cập nhật thông tin nghề nghiệp từ bảng POS sang bảng khách hàng
	for _, job := range jobPos {
		if job.CustomerCode != "" {
			fmt.Println("job: ", job)
			customer, _ := _repository.GetCustomerCRMByCode(ctx, job.CustomerCode)
			fmt.Println("customer: ", customer)
			if customer != nil {
				if (customer.UpdatedAt <= job.UpdatedAt && customer.Job != job.Job) || customer.Job == "none" {
					err := _repository.UpdateJobOfCRM(ctx, job.CustomerCode, job.Job)

					if err != nil {
						fmt.Println("Cập nhật thông tin nghề nghiệp từ bảng POS sang bảng khách hàng thất bại")
					}
				}
			}
		}
	}

	fmt.Println("Cập nhật thông tin nghề nghiệp thành công")

	return nil
}
