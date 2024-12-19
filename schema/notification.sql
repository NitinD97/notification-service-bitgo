CREATE TYPE notification_status AS ENUM ('SENT', 'PENDING', 'FAILED');

-- name: notifications table
CREATE TABLE IF NOT EXISTS notifications
(
    id             UUID PRIMARY KEY             DEFAULT uuid_generate_v4(),
    current_price  DECIMAL(10, 2)      NOT NULL,
    percent_change DECIMAL(10, 2)      NOT NULL,
    volume         INTEGER             NOT NULL,
    user_id        UUID                NOT NULL REFERENCES users (id),
    status         notification_status NOT NULL DEFAULT 'PENDING',
    created_at     TIMESTAMP WITH TIME ZONE     DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP WITH TIME ZONE     DEFAULT CURRENT_TIMESTAMP,
    deleted_at     TIMESTAMP WITH TIME ZONE
);