CREATE TABLE users (
    id serial primary key,
    email varchar(50) not null,
    firs_name varchar(50) not null,
    last_name varchar(50) not null,
    password varchar(255) not null,
    profile_image varchar(255) not null
)