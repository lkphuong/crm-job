package main

import (
	"context"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/lkphuong/crm-job/internal/modules/customer"
)

func main() {

	ctx := context.Background()

	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	//httpPort := fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))

	fmt.Println("Server running on port")

	//r := gin.Default()

	// notification_schedule.DeleteNotificationDraft(ctx)
	// voucher_gift.VoucherGiftExpire(ctx)
	// voucher_gift.InsertVoucherGiftSalePublicCode(ctx)
	// voucher_gift.UpdateVoucherUsed(ctx)
	// voucher_gift.GetVoucherDuplicate(ctx)
	// customer.UpdateCustomer(ctx)
	// earning_point.EarningPointBillHoangDieu2(ctx)
	// earning_point.EaringPoint(ctx)
	// earning_point.ExpiredPoint30Days(ctx)
	customer.UpdateJob(ctx)

	//r.Run(httpPort)

	select {}
}
