package voucher_gift

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/queries"
)

type VoucherGiftRepository struct{}

func (vg *VoucherGiftRepository) GetVoucherGiftExpire(ctx context.Context, db *sql.DB) ([]VoucherGift, error) {
	var vouchers []VoucherGift

	err := queries.Raw(GET_VOUCHER_GIFT_EXPIRE).Bind(ctx, db, &vouchers)

	if err != nil {
		return nil, err
	}

	return vouchers, nil
}

func (vg *VoucherGiftRepository) UpdateVoucherGiftExpire(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, UPDATE_VOUCHER_GIFT_EXPIRE)

	if err != nil {
		return err
	}

	return nil
}

func (vg *VoucherGiftRepository) GetVoucherBirthDuplicate(ctx context.Context, db *sql.DB) ([]VoucherGift, error) {
	var vouchers []VoucherGift

	err := queries.Raw(GET_VOUCHER_BIRTHDAY_DUPLICATE).Bind(ctx, db, &vouchers)

	if err != nil {
		return nil, err
	}

	return vouchers, nil
}
