CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE if not exists users(
  id UUID DEFAULT gen_random_uuid () PRIMARY KEY NOT NULL UNIQUE,
  username text NOT NULL,
  phone_num text NULL,
  market text NOT NULL,
  snkrs_global_enabled bool DEFAULT false NOT NULL,
  stripe_customer_id text NOT NULL UNIQUE,
  stripe_plan text DEFAULT 'FREE' NOT NULL,
  stripe_subscription_id text NULL UNIQUE,
  plan_ends timestamptz NULL,
  auth_user_id UUID NOT NULL UNIQUE
);

ALTER TABLE
  users DROP CONSTRAINT IF EXISTS fk_users_auth_users;

ALTER TABLE
  users
ADD
  CONSTRAINT fk_users_auth_users FOREIGN KEY (auth_user_id) REFERENCES auth.users(id);

ALTER TABLE
  users ENABLE ROW LEVEL SECURITY;