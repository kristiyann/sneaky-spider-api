ALTER TABLE
    product_availabilities
ADD
    COLUMN IF NOT EXISTS availability_json json;

ALTER TABLE
    product_availabilities DROP COLUMN IF EXISTS market;