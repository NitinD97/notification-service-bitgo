-- name: CreateEmailNotification :exec
INSERT INTO email_notifications (notification_id, sent_to)
VALUES ($1, $2);