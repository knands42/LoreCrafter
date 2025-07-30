DROP FUNCTION IF EXISTS notify_new_notification();

DROP TRIGGER IF EXISTS notify_new_notification ON notifications;
DROP FUNCTION IF EXISTS notify_new_notification();

DROP TRIGGER IF EXISTS update_read_at ON notifications;
DROP FUNCTION IF EXISTS update_read_at();

DROP INDEX IF EXISTS idx_notifications_user_id_type;
DROP INDEX IF EXISTS idx_notifications_status;
DROP INDEX IF EXISTS idx_notifications_type;
DROP INDEX IF EXISTS notifications_user_id_idx;

DROP TABLE IF EXISTS notifications;
DROP TYPE IF EXISTS notification_type;