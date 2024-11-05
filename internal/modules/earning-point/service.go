package earning_point

import (
	"context"
	"fmt"
	"strconv"
)

var (
	repository Repository
)

type Service struct{}

func (s *Service) EaringPoint(ctx context.Context) error {
	saleReceiptInfos, err := repository.GetSaleReceiptInfo(ctx)
	if err != nil {
		return err
	}

	for _, saleReceiptInfo := range saleReceiptInfos {
		fmt.Println("Earning point: ", saleReceiptInfo)
		if err := repository.InsertEarningPointHistory(ctx, saleReceiptInfo.CustomerCode, saleReceiptInfo.ReceiptNumber); err != nil {
			return err
		}

		if saleReceiptInfo.MembershipLevelCode == "KC" {
			if err := repository.InsertEarningPointHistoryRankDiamond(ctx, saleReceiptInfo.CustomerCode, saleReceiptInfo.ReceiptNumber); err != nil {
				return err
			}
		}

		point := 0

		if saleReceiptInfo.MembershipLevelCode == "KC" {
			pointRankDaimond, err := repository.CalculatorPointRankDaimond(ctx, saleReceiptInfo.ReceiptNumber)

			if err != nil {
				return err
			}

			point = pointRankDaimond[0].Point
		} else {
			pointRankNormal, err := repository.CalculatorPoint(ctx, saleReceiptInfo.ReceiptNumber)

			if err != nil {
				return err
			}

			if len(pointRankNormal) > 0 {
				point = pointRankNormal[0].Point
			}
		}

		if point > 0 {
			if err := repository.SendNotification(ctx, saleReceiptInfo.CustomerCode, strconv.Itoa(point)); err != nil {
				return err
			}

			if err := repository.AddLoyaltyFirstBill(ctx, saleReceiptInfo.CustomerCode, strconv.Itoa(point)); err != nil {
				return err
			}

			if err := repository.AddReferralReward(ctx, saleReceiptInfo.CustomerCode, saleReceiptInfo.ReceiptNumber); err != nil {
				return err
			}

		}
	}

	return nil
}

func (s *Service) EaringPointHoangDieu2(ctx context.Context) error {
	billEarningPointHoangDieuStore, err := repository.GetBillHoangDieu2(ctx)

	if err != nil {
		return err
	}

	for _, bill := range billEarningPointHoangDieuStore {
		if repository.InsertEarningPointHistoryHoangDieu2(ctx, bill.ID, bill.CuahangID); err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) ExpiredPoint(ctx context.Context) error {
	expiredPointResponses, err := repository.GetExpiredPoint30Days(ctx)

	if err != nil {
		return err
	}

	fmt.Println("start")
	for _, expiredPointResponse := range expiredPointResponses {

		fmt.Println("TransactionNumber: ", expiredPointResponse.TransactionNumber)

		if err := repository.InsertAlmostExpiredPoints(ctx, expiredPointResponse.TransactionNumber, expiredPointResponse.CustomerCode, expiredPointResponse.AvalaibleValue); err != nil {
			return err
		}
	}

	return nil
}
