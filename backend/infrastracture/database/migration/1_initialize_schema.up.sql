create table tasks
(
    id          int unsigned auto_increment primary key,
    created_at  timestamp    null,
    updated_at  timestamp    null,
    deleted_at  timestamp    null,
    user_refer  int unsigned null,
    title       varchar(255) null,
    description varchar(255) null
);

create table users
(
    id         int unsigned auto_increment primary key,
    created_at timestamp    null,
    updated_at timestamp    null,
    deleted_at timestamp    null,
    name       varchar(255) null,
    birth_day  varchar(255) null,
    gender     varchar(255) null,
    photo_url  varchar(255) null,
    active     tinyint(1)   null
);

create index idx_tasks_deleted_at on tasks (deleted_at);
create index idx_users_deleted_at on users (deleted_at);

