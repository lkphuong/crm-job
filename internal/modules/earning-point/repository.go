package earning_point

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/queries"
)

type Repository struct{}

func (r *Repository) GetSaleReceiptInfo(ctx context.Context) ([]SaleReceiptInfo, error) {
	var saleReceiptInfos []SaleReceiptInfo

	err := queries.Raw(GET_SALE_RECEIPT_INFO).Bind(ctx, db, &saleReceiptInfos)

	if err != nil {
		return nil, err
	}

	return saleReceiptInfos, nil
}

func (r *Repository) InsertEarningPointHistory(ctx context.Context, customerCode string, receiptNumber string) error {
	_, err := db.ExecContext(ctx, fmt.Sprintf(INSERT_EARNING_POINT_HISTORY, customerCode, receiptNumber))

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) InsertEarningPointHistoryRankDiamond(ctx context.Context, customerCode string, receiptNumber string) error {
	_, err := db.ExecContext(ctx, fmt.Sprintf(INSERT_EARNING_POINT_HISTORY_RANK_DIAMOND, customerCode, receiptNumber))

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetBillHoangDieu2(ctx context.Context) ([]BillEarningPointResponse, error) {
	var billEarningPointResponses []BillEarningPointResponse

	err := queries.Raw(GET_BILL_HOANG_DIEU2).Bind(ctx, db, &billEarningPointResponses)

	if err != nil {
		return nil, err
	}

	return billEarningPointResponses, nil
}

func (r *Repository) InsertEarningPointHistoryHoangDieu2(ctx context.Context, billID string, storeID string) error {
	_, err := db.ExecContext(ctx, fmt.Sprintf(INSERT_EARNING_POINT_HISTORY_HOANGDIEU2, billID, storeID))

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) CalculatorPointRankDaimond(ctx context.Context, receiptNumber string) ([]PointResponse, error) {
	var pointResponse []PointResponse

	err := queries.Raw(fmt.Sprintf(CALCULATOR_POINT_RANK_DAIMOND, receiptNumber)).Bind(ctx, db, &pointResponse)

	if err != nil {
		return nil, err
	}

	return pointResponse, nil
}

func (r *Repository) CalculatorPoint(ctx context.Context, receiptNumber string) ([]PointResponse, error) {
	var pointResponse []PointResponse

	err := queries.Raw(fmt.Sprintf(CALCULATOR_POINT, receiptNumber)).Bind(ctx, db, &pointResponse)

	if err != nil {
		return nil, err
	}

	return pointResponse, nil
}

func (r *Repository) SendNotification(ctx context.Context, customerCode string, point string) error {
	_, err := db.ExecContext(ctx, fmt.Sprintf(SEND_NOTIFICATION, customerCode, point))

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) AddLoyaltyFirstBill(ctx context.Context, customerCode string, point string) error {
	_, err := db.ExecContext(ctx, fmt.Sprintf(ADD_LOYALTY_FIRST_BILL, customerCode, point))

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) AddReferralReward(ctx context.Context, customerCode string, point string) error {
	_, err := db.ExecContext(ctx, fmt.Sprintf(ADD_REFERRAL_REWARD, customerCode, point))

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetExpiredPoint30Days(ctx context.Context) ([]ExpiredPointResponse, error) {
	var expiredPointResponses []ExpiredPointResponse

	err := queries.Raw(SELECT_EXPIRED_30DAYS).Bind(ctx, db, &expiredPointResponses)

	if err != nil {
		return nil, err
	}

	return expiredPointResponses, nil
}

func (r *Repository) InsertAlmostExpiredPoints(ctx context.Context, transactionNumber string, customerCode string, pointValue string) error {
	_, err := db.ExecContext(ctx, fmt.Sprintf(INSERT_ALMOST_EXPIRED_POINTS, transactionNumber, customerCode, pointValue))

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetCurrentPoint(ctx context.Context) ([]CurrentPointResponse, error) {
	var currentPointResponses []CurrentPointResponse

	err := queries.Raw(GET_CURRENT_POINT).Bind(ctx, db, &currentPointResponses)

	fmt.Println("err: ", err)

	if err != nil {
		return nil, err
	}

	return currentPointResponses, nil
}

func (r *Repository) UpdateNewPoint(ctx context.Context, customerCode string) error {
	_, err := db.ExecContext(ctx, fmt.Sprintf(UPDATE_NEW_POINT, customerCode))

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetEarningPointExpired(ctx context.Context) ([]EarningPointExpired, error) {
	var earningPointExpiredResponses []EarningPointExpired

	err := queries.Raw(GET_EARNING_POINT_EXPIRED).Bind(ctx, db, &earningPointExpiredResponses)

	fmt.Println("err: ", err)
	if err != nil {
		return nil, err
	}

	return earningPointExpiredResponses, nil
}

func (r *Repository) InsertEarningPointHistoryExpired(ctx context.Context, param EarningPointExpired) error {

	newUUID := uuid.New()

	_, err := db.ExecContext(
		ctx,
		INSERT_EARNING_POINT_HISTORY_EXPIRED,
		sql.Named("TransactionNumber", newUUID),
		sql.Named("CustomerCode", param.CustomerCode),
		sql.Named("StoreCode", param.StoreCode),
		sql.Named("Value", -param.AvalaibleValue), // Giá trị là số
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateEarningPointExpired(ctx context.Context, Id int64) error {
	_, err := db.ExecContext(
		ctx,
		UPDATE_EARNING_POINT_EXPIRED,
		sql.Named("Id", Id),
	)

	if err != nil {
		return err
	}

	return nil
}
