-- +goose Up
-- +goose StatementBegin
create table chats(
    id_chat serial primary key,
    chat_name TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table note;
-- +goose StatementEnd