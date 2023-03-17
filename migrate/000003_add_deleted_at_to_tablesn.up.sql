BEGIN;
ALTER TABLE events
   ADD COLUMN "deleted_at" timestamptz DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE subscriptions
   ADD COLUMN "deleted_at" timestamptz DEFAULT CURRENT_TIMESTAMP;
COMMIT;