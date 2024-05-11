-- +migrate Up
create table users (
    user_id VARCHAR(255) primary key,
    full_name VARCHAR(255),
    email VARCHAR(255) unique,
    password VARCHAR(255),
    role VARCHAR(255),
    created_at datetime not null,
    updated_at datetime not null
);

create table repos (
    name VARCHAR(255) primary key,
    description VARCHAR(255),
    url VARCHAR(255),
    color VARCHAR(255),
    lang VARCHAR(255),
    fork VARCHAR(255),
    stars VARCHAR(255),
    stars_today VARCHAR(255),
    build_by VARCHAR(255),
    created_at datetime not null,
    updated_at datetime not null
);

create table bookmarks (
    bid VARCHAR(255) primary key,
    user_id VARCHAR(255),
    repo_name VARCHAR(255),
    created_at datetime not null,
    updated_at datetime not null
);

alter table bookmarks add foreign key (user_id) references users (user_id);
alter table bookmarks add foreign key (repo_name) references repos (name);

-- +migrate Down
DROP TABLE users;
DROP TABLE repos;
DROP TABLE bookmarks;