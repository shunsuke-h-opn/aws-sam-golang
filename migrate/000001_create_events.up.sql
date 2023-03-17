BEGIN;
CREATE TABLE if not exists "events" (
    "id" BIGSERIAL PRIMARY KEY,
    "event_type" text,
    "payload" text,
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP
);
comment on table events is 'events';
comment on column events.id is 'ID';
comment on column events.event_type is 'event_type';
comment on column events.payload is 'payload';
comment on column events.created_at is '作成日時';
comment on column events.updated_at is '更新日時';
COMMIT;
