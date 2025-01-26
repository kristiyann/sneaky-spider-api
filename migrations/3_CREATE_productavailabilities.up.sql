CREATE TABLE if not exists product_availabilities(
    id UUID DEFAULT gen_random_uuid () PRIMARY KEY NOT NULL UNIQUE,
    vendor product_subscriptions_vendor NOT NULL,
    available bool NOT NULL DEFAULT false,
    product_external_id text NOT NULL,
    market text[] NOT NULL DEFAULT array[]::text[],
    available_sizes text[] NOT NULL DEFAULT array[]::text[],
    create_date timestamptz DEFAULT now() NOT NULL,
    last_change_date timestamptz DEFAULT now() NOT NULL
);

ALTER TABLE
    product_availabilities ENABLE ROW LEVEL SECURITY;