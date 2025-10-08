-- +goose Up
-- +goose StatementBegin
create type user_role as ENUM('user', 'admin');

create table users(
    id serial primary key,
    name text not null,
    email text not null unique,
    password_hash TEXT not null,
    role user_role not null default 'user' 
    create_at timestamp default now() 
    update_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
drop type if exists user_role;
-- +goose StatementEnd