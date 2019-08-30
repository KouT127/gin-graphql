create table items
(
    id          int unsigned auto_increment primary key,
    created_at  timestamp    null,
    updated_at  timestamp    null,
    deleted_at  timestamp    null,
    name        varchar(255) null,
    description varchar(255) null
);

create table carts
(
    id         int unsigned auto_increment primary key,
    created_at timestamp    null,
    updated_at timestamp    null,
    deleted_at timestamp    null,
    user_refer int unsigned null,
    item_refer int unsigned null
);

create index idx_items_deleted_at on items (deleted_at);
create index idx_carts_deleted_at on carts (deleted_at);

