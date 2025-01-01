package notification_schedule

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

var (
	repository Repository
)

type Service struct{}

func sendMultiNotification(client *messaging.Client, tokens []string, title string, body string) error {
	message := &messaging.MulticastMessage{
		Tokens: tokens,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
	}

	response, err := client.SendEachForMulticast(context.Background(), message)

	if err != nil {
		return err
	}

	fmt.Println("Successfully sent message to", response.SuccessCount, "devices")

	if response.FailureCount > 0 {
		fmt.Println("Failed to send message to", response.FailureCount, "devices")
	}

	return nil
}

func convertTokens(tokens []FirebaseToken) []string {
	stringTokens := make([]string, len(tokens))
	for i, token := range tokens {
		stringTokens[i] = token.FirebaseToken
	}
	return stringTokens
}

func sendNotificationToAllTokens(ctx context.Context, client *messaging.Client, title string, body string) error {
	firebaseTokens, err := repository.GetFirebaseTokens(ctx)

	if err != nil {
		fmt.Println("Error getting firebase tokens: ", err)
		return err
	}

	batchSize := os.Getenv("BATCH_SIZE")

	numBatchSize, err := strconv.Atoi(batchSize)

	if err != nil {
		numBatchSize = 500
	}

	totalTokens := len(firebaseTokens)

	// #region loop through all tokens
	for i := 0; i < totalTokens; i += numBatchSize {
		// create a batch of 500 tokens
		end := i + numBatchSize
		if end > totalTokens {
			end = totalTokens
		}

		batch := firebaseTokens[i:end]

		stringBatch := convertTokens(batch)

		err := sendMultiNotification(client, stringBatch, title, body)
		if err != nil {
			fmt.Println("Error sending notification: ", err)
		}

		// // send the batch in a goroutine
		// go func(batch []string) {
		// 	defer wg.Done()

		// 	err := sendMultiNotification(client, batch, title, body)
		// 	if err != nil {
		// 		fmt.Println("Error sending notification: ", err)
		// 	}
		// }(stringBatch)
	}

	fmt.Println("All batches sent")
	// #endregion
	return nil
}

func (s *Service) InsertNotificationSchedule(ctx context.Context, db *sql.DB, title string, body string) error {

	err := repository.InsertNotificationSchedule(ctx, db, title, body)

	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteNotificationDraft(ctx context.Context, db *sql.DB) error {

	err := repository.DeleteNotificationDraft(ctx, db)

	if err != nil {
		return err
	}

	return nil
}

func (s *Service) PushNotificationCampaign(ctx context.Context, db *sql.DB) error {

	notificationCampaign, err := repository.GetNotificationCampaign(ctx)

	if err != nil || len(notificationCampaign) == 0 {
		return err
	}

	fmt.Println("notificationCampaign: ", notificationCampaign)

	repository.UpdateNotificationCampaign(ctx, db, notificationCampaign[0].ID)

	opt := option.WithCredentialsFile("./keys/loyalty-service-account.json")

	// Init firebase app
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Println("Error initializing app: ", err)
	}

	// Get a messaging client
	client, err := app.Messaging(ctx)
	if err != nil {
		fmt.Println("Error getting messaging client: ", err)
	}

	title := notificationCampaign[0].Title
	body := notificationCampaign[0].Summary

	sendNotificationToAllTokens(ctx, client, title, body)
	return nil
}
