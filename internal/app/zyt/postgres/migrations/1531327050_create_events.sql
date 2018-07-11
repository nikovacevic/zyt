CREATE TABLE IF NOT EXISTS events (
  id uuid NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
  story_id uuid DEFAULT NULL,
  user_id uuid NOT NULL,
  duration interval SECOND DEFAULT NULL,
  end_time timestamp DEFAULT NULL,
  start_time timestamp NOT NULL
);
