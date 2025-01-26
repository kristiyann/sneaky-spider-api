ALTER TABLE
    product_availabilities
ADD
    COLUMN IF NOT EXISTS market text[];

ALTER TABLE
    product_availabilities DROP COLUMN IF EXISTS availability_json;