DO $$ 
BEGIN
  IF NOT EXISTS (
    SELECT 1 
    FROM pg_type 
    WHERE typname = 'product_subscriptions_status'
  ) THEN
    CREATE TYPE product_subscriptions_status AS ENUM ('Active', 'Completed', 'Cancelled', 'Expired');
  END IF;
END $$;

DO $$ 
BEGIN
  IF NOT EXISTS (
    SELECT 1 
    FROM pg_type 
    WHERE typname = 'product_subscriptions_vendor'
  ) THEN
    CREATE TYPE product_subscriptions_vendor as ENUM ('Nike', 'SNKRS', 'Shopify', 'Supreme');
  END IF;
END $$;

CREATE TABLE if not exists product_subscriptions(
  id UUID DEFAULT gen_random_uuid () PRIMARY KEY NOT NULL UNIQUE,
  user_id UUID NOT NULL,
  vendor product_subscriptions_vendor NOT NULL,
  product_external_id text NOT NULL,
  size text NOT NULL,
  notification_email text NOT NULL,
  market text NOT NULL,
  create_date timestamptz DEFAULT now() NOT NULL,
  status_last_change_date timestamptz DEFAULT now() NOT NULL,
  status product_subscriptions_status DEFAULT 'Active' NOT NULL
);

ALTER TABLE
  product_subscriptions DROP CONSTRAINT IF EXISTS fk_product_subscriptions_users;

ALTER TABLE
  product_subscriptions
ADD
  CONSTRAINT fk_product_subscriptions_users FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE
  product_subscriptions ENABLE ROW LEVEL SECURITY;