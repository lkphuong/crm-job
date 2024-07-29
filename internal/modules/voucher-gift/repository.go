package voucher_gift

import (
	"context"
	"fmt"

	"github.com/lkphuong/crm-job/internal/utils"
	"github.com/volatiletech/sqlboiler/queries"
)

type Repository struct{}

func (vg *Repository) GetSalePublicCode(ctx context.Context) ([]SalePublicCode, error) {

	var salePublicCode []SalePublicCode

	err := queries.Raw(GET_SALE_CODE_PUBLIC).Bind(ctx, db, &salePublicCode)

	if err != nil {
		return nil, err
	}

	return salePublicCode, nil

}

func (vg *Repository) GetCoupon(ctx context.Context, saleID string) (*Coupon, error) {

	var coupon Coupon

	err := queries.Raw(fmt.Sprintf(GET_COUPON, saleID)).Bind(ctx, db, &coupon)

	if err != nil {
		return nil, err
	}

	return &coupon, nil
}

func (vg *Repository) GetVoucherGift(ctx context.Context, code string) (*VoucherGiftCode, error) {
	var voucherGiftCode VoucherGiftCode

	err := queries.Raw(fmt.Sprintf(GET_VOUCHER_GIFT, code)).Bind(ctx, db, &voucherGiftCode)

	if err != nil {
		return nil, err
	}

	return &voucherGiftCode, nil
}

func (vg *Repository) GetVoucherGiftExpire(ctx context.Context) ([]VoucherGift, error) {
	var vouchers []VoucherGift

	err := queries.Raw(GET_VOUCHER_GIFT_EXPIRE).Bind(ctx, db, &vouchers)

	if err != nil {
		return nil, err
	}

	return vouchers, nil
}

func (vg *Repository) GetVoucherDuplicate(ctx context.Context) ([]VoucherDuplicate, error) {
	var vouchers []VoucherDuplicate

	err := queries.Raw(GET_VOUCHER_DUPLICATE).Bind(ctx, db, &vouchers)

	if err != nil {
		return nil, err
	}

	return vouchers, nil
}

func (vg *Repository) UpdateVoucherGiftUsed(ctx context.Context) error {
	_, err := db.ExecContext(ctx, UPDATE_VOUCHER_GIFT_USED)

	if err != nil {
		return err
	}

	return nil
}

func (vg *Repository) GetVoucherBirthDuplicate(ctx context.Context) ([]VoucherGift, error) {
	var vouchers []VoucherGift

	err := queries.Raw(GET_VOUCHER_BIRTHDAY_DUPLICATE).Bind(ctx, db, &vouchers)

	if err != nil {
		return nil, err
	}

	return vouchers, nil
}

func (vg *Repository) InsertVoucherGift(ctx context.Context, name string, start string, end string, saleId string, refSku string) error {

	_, err := db.ExecContext(ctx, fmt.Sprintf(INSERT_VOUCHER_GIFT, name, start, end, refSku, saleId))

	if err != nil {
		return err
	}

	return nil
}

func (vg *Repository) UpdateVoucherGiftExpire(ctx context.Context) error {

	startOfDay := utils.GetCurrentDateStartTime()
	endOfDay := utils.GetCurrentDateEndTime()

	_, err := db.ExecContext(ctx, fmt.Sprintf(UPDATE_EXPIRED_VOUCHER_GIFT, endOfDay, startOfDay))

	if err != nil {
		return err
	}

	return nil
}
