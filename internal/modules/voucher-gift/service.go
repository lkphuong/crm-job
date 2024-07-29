package voucher_gift

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var (
	repository Repository
)

func init() {
	repository = Repository{}
}

func SendNotificationVoucherDuplicate(ctx context.Context) error {
	vouchers, _ := repository.GetVoucherDuplicate(ctx)

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	chatId := os.Getenv("CHAT_ID")
	teleToken := os.Getenv("TELE_TOKEN")

	if len(vouchers) > 0 {
		// Create a payload for the API request
		payload := map[string]interface{}{
			"chat_id": chatId,
			"text": struct {
				Title  string             `json:"title"`
				Values []VoucherDuplicate `json:"values"`
			}{
				Title:  "Voucher Duplicate",
				Values: vouchers,
			},
		}

		// Convert the payload to JSON
		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			return err
		}

		// Create a new HTTP request
		req, err := http.NewRequest("POST", fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", teleToken), bytes.NewBuffer(payloadBytes))
		if err != nil {
			return err
		}

		// Set the request headers
		req.Header.Set("Content-Type", "application/json")

		// Send the request
		client := http.DefaultClient
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		return nil
	}

	return nil
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
