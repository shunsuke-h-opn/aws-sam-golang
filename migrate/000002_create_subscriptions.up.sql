BEGIN;
CREATE TABLE if not exists "subscriptions" (
    "id" BIGSERIAL PRIMARY KEY,
    "event_type" text,
    "endpoint" text,
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP
);
comment on table subscriptions is 'events';
comment on column subscriptions.id is 'ID';
comment on column subscriptions.event_type is 'event_type';
comment on column subscriptions.endpoint is 'endpoint';
comment on column subscriptions.created_at is '作成日時';
comment on column subscriptions.updated_at is '更新日時';
COMMIT;
