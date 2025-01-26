ALTER TABLE
    product_availabilities
ADD
    COLUMN IF NOT EXISTS last_recorded_launch_date json;

ALTER TABLE
    product_availabilities DROP COLUMN IF EXISTS available;