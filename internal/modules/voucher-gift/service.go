package voucher_gift

import (
	"context"
)

var (
	repository Repository
)

func init() {
	repository = Repository{}
}

func UpdateVoucherGiftExpire(ctx context.Context) error {

	repository.UpdateVoucherGiftExpire(ctx)

	return nil
}

func InsertVoucherGift(ctx context.Context) error {
	salePuclicCode, err := repository.GetSalePublicCode(ctx)

	if err != nil {
		return err
	}

	coupon, err := repository.GetCoupon(ctx, salePuclicCode[0].SaleID)

	if err != nil {
		return err
	}

	voucherGift, _ := repository.GetVoucherGift(ctx, coupon.Code)

	if voucherGift == nil {
		err := repository.InsertVoucherGift(ctx, coupon.Name, coupon.Begin, coupon.End, coupon.SaleID, coupon.RefSku)

		if err != nil {
			return err
		}
	}

	return nil

}

func UpdateVoucherUsed(ctx context.Context) error {

	repository.UpdateVoucherGiftUsed(ctx)

	return nil
}
