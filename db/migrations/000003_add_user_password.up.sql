alter table if exists users
    add column if not exists password varchar(255) not null;