CREATE TYPE notification_status AS ENUM ('unread', 'read');

CREATE TABLE notifications
(
    id         UUID PRIMARY KEY,
    user_id    UUID        NOT NULL,
    type       VARCHAR(30) NOT NULL,
    payload    JSONB       NOT NULL DEFAULT '{}'::jsonb,
    status     TEXT        NOT NULL DEFAULT 'unread',
    created_at TIMESTAMP            DEFAULT now()
);

-- Emit notifications whenever a new input is created
CREATE OR REPLACE FUNCTION notify_new_notification()
    RETURNS trigger AS
$$
DECLARE
    payload JSON;
BEGIN
    -- You can customize the payload structure
    payload := json_build_object(
            'id', NEW.id,
            'user_id', NEW.user_id,
            'type', NEW.type,
            'status', NEW.status,
            'created_at', NEW.created_at
               );

    -- Send a notification on a specific channel
    PERFORM pg_notify('new_notification', payload::text);

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
