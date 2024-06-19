package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	notification_schedule "github.com/lkphuong/crm-job/internal/modules/notification_schedule"
)

func main() {

	ctx := context.Background()

	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	httpPort := fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))

	fmt.Println("Server running on port", httpPort)

	r := gin.Default()

	notification_schedule.DeleteNotificationDraft(ctx)

	r.Run(httpPort)

	select {}
}
