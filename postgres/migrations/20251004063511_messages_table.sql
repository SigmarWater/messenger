-- +goose Up
-- +goose StatementBegin
create table messages(
    id_message serial primary key,
    id_chat int not null,
    from_user int not null,
    text_message text, 
    time_at timestamp default now()

    foreign key (from_user) references users(id) on delete cascade;
    foreign key (id_chat) references chat(id_chat) on delete cascade;
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists messages;
-- +goose StatementEnd
