-- +goose Up
-- +goose StatementBegin
CREATE TABLE birth_list (
    id  serial primary key,
    firstname varchar(50) not null,
    lastname varchar(50) not null,
    day int not null,
    month int not null,
    year int null,
    comment varchar(255) null
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE birth_list;
-- +goose StatementEnd
