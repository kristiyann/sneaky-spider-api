DO $$ 
BEGIN
  IF NOT EXISTS (
    SELECT 1 
    FROM pg_type 
    WHERE typname = 'logs_level'
  ) THEN
    CREATE TYPE logs_level AS ENUM ('None', 'Info', 'Warn', 'Debug', 'Error', 'Fatal');
  END IF;
END $$;

DO $$ 
BEGIN
  IF NOT EXISTS (
    SELECT 1 
    FROM pg_type 
    WHERE typname = 'logs_source'
  ) THEN
    CREATE TYPE logs_source as ENUM ('API', 'Alerts', 'Interface');
  END IF;
END $$;

CREATE TABLE if not exists logs(
  id UUID DEFAULT gen_random_uuid () PRIMARY KEY NOT NULL UNIQUE,
  create_date timestamptz DEFAULT now() NOT NULL,
  user_id UUID NULL,
  message text NOT NULL,
  context text NULL,
  request_body text NULL,
  stack_trace text NULL,
  status_code smallint NOT NULL,
  api_url text NOT NULL,
  interface_url text NULL,
  level logs_level NOT NULL DEFAULT 'None',
  source logs_source NOT NULL DEFAULT 'API'
);

ALTER TABLE
  logs
ADD
  CONSTRAINT fk_logs_users FOREIGN KEY (user_id) REFERENCES public.users(id);

ALTER TABLE
  logs ENABLE ROW LEVEL SECURITY;