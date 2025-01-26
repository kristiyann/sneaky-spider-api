DROP TABLE if exists logs;

DO $$ 
BEGIN IF EXISTS (
  SELECT
    1
  FROM
    pg_type
  WHERE
    typname = 'logs_level'
) THEN DROP TYPE logs_level;
END IF;
END $$;

DO $$ 
BEGIN IF EXISTS (
  SELECT
    1
  FROM
    pg_type
  WHERE
    typname = 'logs_source'
) THEN DROP TYPE logs_source;
END IF;
END $$;

DROP TYPE if exists logs_source;