CREATE TABLE if not exists webhook_urls(
  id UUID DEFAULT gen_random_uuid () PRIMARY KEY NOT NULL UNIQUE,
  create_date timestamptz DEFAULT now() NOT NULL,
  last_change_date timestamptz DEFAULT now() NOT NULL,
  user_id UUID NOT NULL,
  url text NOT NULL
);

ALTER TABLE
  webhook_urls DROP CONSTRAINT IF EXISTS fk_webhook_urls_users;

ALTER TABLE
  webhook_urls
ADD
  CONSTRAINT fk_webhook_urls_users FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE
  webhook_urls ENABLE ROW LEVEL SECURITY;