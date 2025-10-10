-- +goose Up
-- +goose StatementBegin
create table chats(
    id_chat serial primary key,
    chat_name TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists chats;
-- +goose StatementEnd