-- +goose Up
-- +goose StatementBegin
create table note(
    id serial primary key,
    title text not null,
    body text not null,
    create_at timestamp not null default now(),
    update_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table note;
-- +goose StatementEnd
