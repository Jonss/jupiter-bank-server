CREATE TABLE IF NOT EXISTS users
(
    id
               bigserial
        PRIMARY
            KEY,
    external_id
               uuid
        UNIQUE
                           NOT
                               NULL,
    fullname
               varchar
                           NOT
                               NULL,
    email
               varchar
        UNIQUE
                           NOT
                               NULL,
    password_hash
               varchar(60) NOT NULL,
    pin        varchar(4),
    tax_id     varchar(60) UNIQUE,
    is_active  bool                 default true,
    created_at timestamp   NOT NULL DEFAULT
                                        (
                                            now
                                                (
                                                )),
    updated_at timestamp   NOT NULL DEFAULT
                                        (
                                            now
                                                (
                                                ))
)
