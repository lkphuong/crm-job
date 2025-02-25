package main

import (
	"context"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/lkphuong/crm-job/internal/modules/customer"
	earning_point "github.com/lkphuong/crm-job/internal/modules/earning-point"
	notification_schedule "github.com/lkphuong/crm-job/internal/modules/notification-schedule"
	voucher_gift "github.com/lkphuong/crm-job/internal/modules/voucher-gift"
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

	notification_schedule.DeleteNotificationDraft(ctx)
	notification_schedule.PushNotificationCampaign(ctx)

	voucher_gift.VoucherGiftExpire(ctx)
	voucher_gift.InsertVoucherGiftSalePublicCode(ctx)
	voucher_gift.UpdateVoucherUsed(ctx)
	voucher_gift.GetVoucherDuplicate(ctx)

	customer.UpdateCustomer(ctx)
	customer.UpdateJob(ctx)

	earning_point.EarningPointBillHoangDieu2(ctx)
	earning_point.EaringPoint(ctx)
	earning_point.ExpiredPoint30Days(ctx)
	earning_point.UpdateNewPoint(ctx)
	earning_point.InsertEarningPointExpired(ctx)

	//r.Run(httpPort)

	select {}
}
