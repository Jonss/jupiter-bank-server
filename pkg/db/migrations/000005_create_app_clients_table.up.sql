CREATE TABLE IF NOT EXISTS app_clients
(
    id
    bigserial
    PRIMARY
    KEY,
    name
    VARCHAR
    UNIQUE
    NOT
    NULL,
    api_key
    VARCHAR
    UNIQUE
    NOT
    NULL,
    secret
    VARCHAR
    UNIQUE
    NOT
    NULL,
    created_at
    timestamp
    NOT
    NULL
    DEFAULT (
    now
(
))
    );

CREATE INDEX app_clients_api_key_idx ON app_clients (api_key);
CREATE INDEX app_clients_secret_idx ON app_clients (secret);