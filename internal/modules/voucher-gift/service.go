package voucher_gift

import (
	"context"
	"database/sql"
)

var (
	_repository Repository
)

func init() {
	_repository = Repository{}
}

func UpdateVoucherGiftExpire(ctx context.Context, db *sql.DB) error {

	vouchers, err := _repository.GetVoucherGiftExpire(ctx, db)

	if err != nil {
		return err
	}

	if len(vouchers) > 0 {
		return _repository.UpdateVoucherGiftExpire(ctx, db)
	}

	return nil
}

func InsertVoucherGift(ctx context.Context, db *sql.DB) error {
	salePuclicCode, err := _repository.GetSalePublicCode(ctx, db)

	if err != nil {
		return err
	}

	coupon, err := _repository.GetCoupon(ctx, db, salePuclicCode[0].SaleID)

	if err != nil {
		return err
	}

	voucherGift, _ := _repository.GetVoucherGift(ctx, db, coupon.Code)

	if voucherGift == nil {
		err := _repository.InsertVoucherGift(ctx, db, coupon.Name, coupon.Begin, coupon.End, coupon.SaleID, coupon.RefSku)

		if err != nil {
			return err
		}
	}

	return nil

}
