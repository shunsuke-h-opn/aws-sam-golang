BEGIN;
ALTER TABLE events
   DROP COLUMN IF EXISTS "deleted_at";
ALTER TABLE subscriptions
   DROP COLUMN IF EXISTS "deleted_at";
COMMIT;