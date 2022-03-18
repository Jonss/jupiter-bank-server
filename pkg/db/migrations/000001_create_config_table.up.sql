CREATE TABLE IF NOT EXISTS configs
(
    id            bigserial PRIMARY KEY,
    key varchar   NOT NULL,
    value varchar NOT NULL,
    created_at    timestamp        NOT NULL DEFAULT (now()),
    updated_at    timestamp        NOT NULL DEFAULT (now())
)