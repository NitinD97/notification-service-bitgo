-- name: GetNotificationById :one
SELECT *
from notifications
where id = $1;

-- name: GetNotificationsByUserId :many
SELECT *
from notifications
where user_id = $1;

-- name: CreateNotification :exec
INSERT INTO notifications (id, current_price, percent_change, volume, user_id, status)
VALUES ($1, $2, $3, $4, $5, 'PENDING');

-- name: UpdateNotificationStatusById :exec
UPDATE notifications
SET status = $1
where id = $2;