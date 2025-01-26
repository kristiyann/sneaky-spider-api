ALTER TABLE
    product_availabilities
ADD
    COLUMN IF NOT EXISTS available boolean NOT NULL DEFAULT false;

ALTER TABLE
    product_availabilities DROP COLUMN IF EXISTS last_recorded_launch_date;