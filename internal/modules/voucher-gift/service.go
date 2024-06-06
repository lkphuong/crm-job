package voucher_gift

import (
	"context"
	"database/sql"
)

var (
	_repository VoucherGiftRepository
)

func init() {
	_repository = VoucherGiftRepository{}
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
