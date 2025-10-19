-- +goose Up
-- +goose StatementBegin
create table messages(
    id_message serial primary key,
    id_chat int not null references chats(id_chat) on delete cascade,
    from_user text not null,
    text_message text, 
    time_at timestamp default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists messages;
-- +goose StatementEnd
