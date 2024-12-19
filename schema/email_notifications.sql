CREATE TABLE IF NOT EXISTS email_notifications
(
    id              UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    notification_id uuid REFERENCES notifications (id) NOT NULL,
    sent_to         VARCHAR(200)                       NOT NULL,
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);