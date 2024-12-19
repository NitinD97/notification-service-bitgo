package main

import "github.com/google/uuid"

type CreateNotificationInput struct {
	CurrentPrice  float64   `json:"current_price"`
	PercentChange float64   `json:"percent_change"`
	Volume        int64     `json:"volume"`
	UserID        uuid.UUID `json:"user_id"`
}

type SendNotificationInput struct {
	Emails         []string  `json:"emails"`
	NotificationId uuid.UUID `json:"notification_id"`
}
