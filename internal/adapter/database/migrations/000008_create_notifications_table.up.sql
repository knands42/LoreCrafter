CREATE TYPE notification_status AS ENUM ('unread', 'read');
CREATE TYPE notification_type AS ENUM ('campaign_invite');

CREATE TABLE notifications
(
    id         UUID PRIMARY KEY,
    user_id    UUID                NOT NULL,
    type       notification_type   NOT NULL,
    payload    JSONB               NOT NULL DEFAULT '{}'::jsonb,
    status     notification_status NOT NULL DEFAULT 'unread',
    created_at TIMESTAMP                    DEFAULT now(),
    read_at    TIMESTAMP                    DEFAULT NULL
);

CREATE INDEX "idx_notifications_user_id" ON notifications (user_id);
CREATE INDEX "idx_notifications_type" ON notifications (type);
CREATE INDEX "idx_notifications_status" ON notifications (status);
CREATE INDEX "idx_notifications_user_id_type" ON notifications (user_id, type);

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
            'payload', NEW.payload,
            'type', NEW.type,
            'status', NEW.status,
            'created_at', NEW.created_at
               );

    -- Send a notification on a specific channel
    PERFORM pg_notify('new_notification', payload::text);

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER notify_new_notification
    AFTER INSERT
    ON notifications
    FOR EACH ROW
EXECUTE PROCEDURE notify_new_notification();

-- Update read_at whenever status change to read
CREATE OR REPLACE FUNCTION update_read_at()
    RETURNS trigger AS
$$
BEGIN
    IF NEW.status = 'read' THEN
        NEW.read_at := now();
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_read_at
    BEFORE UPDATE
    ON notifications
    FOR EACH ROW
EXECUTE PROCEDURE update_read_at();
