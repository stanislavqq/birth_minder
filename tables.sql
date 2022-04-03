CREATE TABLE birth_users
(
    id         integer
        constraint birth_users_pk
            primary key autoincrement,
    name       varchar not null,
    day        integer not null,
    month      integer not null,
    year       integer,
    comment    text,
    account_id integer not null
        constraint birth_users_accounts_id_fk
            references accounts (id)
);

CREATE TABLE accounts
(
    id         integer
        constraint accounts_pk
            primary key,
    created_at TIMESTAMP default CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX accounts_id_uindex
    on accounts (id)
