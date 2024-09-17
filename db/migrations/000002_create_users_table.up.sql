
create table if not exists users (
	id serial primary key,
	username varchar(255) not null,
	first_name varchar(255),
	last_name varchar(255) not null,
    user_role varchar(255) not null,
	created_at timestamp with time zone not null default now(),
	updated_at timestamp with time zone not null default now()
);