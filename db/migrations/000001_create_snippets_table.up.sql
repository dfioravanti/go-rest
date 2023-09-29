create function now_utc() returns timestamp as $$
  select now() at time zone 'utc';
$$ language sql;


CREATE TABLE snippets (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT now_utc(),
    expires TIMESTAMPTZ NOT NULL
);

-- Add an index on the created column.
CREATE INDEX idx_snippets_created ON snippets(created);