create table album(
    id bigserial primary key,
    title varchar(255) not null,
    artist varchar(255) not null,
    price numeric(10, 2) not null
);